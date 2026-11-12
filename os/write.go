package os

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// maxLinks matches [filepath.EvalSymlinks].
const maxLinks = 255

var errTooManyLinks = errors.New("too many levels of symbolic links")

// AtomicWrite writes `data` to `path` via a temp file renamed into place, and
// removes the temp file on failure.
//
// A symlink at `path` is followed: the file it points to is replaced, or created,
// and the link is kept. Pass [WithNoFollow] to replace the link itself.
//
// The temp file and its directory are synced so the write survives a crash; pass
// [WithNoSync] to skip that. Directory syncing is Unix-only. If it fails after
// the rename, the error is returned and the new file is left in place.
func AtomicWrite(path string, data []byte, perm os.FileMode, opts ...AtomicWriteOption) error {
	var cfg atomicWriteConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	if !cfg.noFollow {
		resolved, err := resolveLinks(path)
		if err != nil {
			return fmt.Errorf("failed to resolve symlink: %w", err)
		}
		path = resolved
	}
	dir := rawDir(path)
	tmp, err := os.CreateTemp(dir, "."+filepath.Base(path)+".*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpName := tmp.Name()
	cleanup := func() { _ = os.Remove(tmpName) }

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		cleanup()
		return fmt.Errorf("failed to write temp file: %w", err)
	}
	if err := tmp.Chmod(perm); err != nil {
		_ = tmp.Close()
		cleanup()
		return fmt.Errorf("failed to chmod temp file: %w", err)
	}
	if !cfg.noSync {
		if err := tmp.Sync(); err != nil {
			_ = tmp.Close()
			cleanup()
			return fmt.Errorf("failed to sync temp file: %w", err)
		}
	}
	if err := tmp.Close(); err != nil {
		cleanup()
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		cleanup()
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	if cfg.noSync {
		return nil
	}
	// Past the rename `cleanup()` must never run: `tmpName` is gone, so it would
	// remove `path` instead and turn a durability gap into data loss.
	if err := syncDir(dir); err != nil {
		return fmt.Errorf("failed to sync directory after rename: %w", err)
	}
	return nil
}

// resolveLinks follows the symlinks at `path` to the first path that is not a
// link, which may not exist. Relative targets are joined without cleaning, so the
// kernel resolves ".." through any symlinked directory.
func resolveLinks(path string) (string, error) {
	for range maxLinks {
		info, err := os.Lstat(path)
		if errors.Is(err, fs.ErrNotExist) {
			return path, nil
		}
		if err != nil {
			return "", err
		}
		if info.Mode()&fs.ModeSymlink == 0 {
			return path, nil
		}
		target, err := os.Readlink(path)
		if err != nil {
			return "", err
		}
		if !filepath.IsAbs(target) {
			target = rawDir(path) + string(filepath.Separator) + target
		}
		path = target
	}
	return "", &fs.PathError{Op: "resolve", Path: path, Err: errTooManyLinks}
}

// rawDir is [filepath.Dir] without cleaning, so ".." keeps its physical meaning.
func rawDir(path string) string {
	vol := len(filepath.VolumeName(path))
	i := len(path) - 1
	// UTF-8 continuation bytes never match a separator.
	for i >= vol && !os.IsPathSeparator(path[i]) {
		i--
	}
	switch {
	case i < vol:
		return path[:vol] + "."
	case i == vol:
		return path[:vol+1]
	default:
		return path[:i]
	}
}

// EnsureDir creates `dir` and any missing parents, and guarantees `dir` itself
// has mode `perm` even if it already existed with a different mode or the
// umask interfered at creation time. Pre-existing parents are left untouched.
func EnsureDir(dir string, perm os.FileMode) error {
	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	info, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("failed to stat directory: %w", err)
	}
	if info.Mode().Perm() != perm {
		if err := os.Chmod(dir, perm); err != nil {
			return fmt.Errorf("failed to chmod directory: %w", err)
		}
	}
	return nil
}

// EnsureFile creates `path` as an empty file with mode `perm` if it does not
// exist, creating any missing parent directories. An existing file's contents,
// mode, and timestamps are left untouched.
func EnsureFile(path string, perm os.FileMode) error {
	// Not [EnsureDir]: parents are incidental here, so a pre-existing parent's
	// mode must be left alone (e.g. a 0o700 ~/.ssh must not become 0o755).
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("failed to create parent directories: %w", err)
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}
	return nil
}

// CopyFile copies `src` to `dst`, preserving `src`'s mode bits. `dst` is fsynced
// before close and its directory afterwards, so a newly created `dst` is durable
// and not just written - the same reasoning as [AtomicWrite], and likewise
// Unix-only. When `src` and `dst` are the same file (including via hard link)
// [CopyFile] is a no-op.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("source is not a regular file: %s", src)
	}

	same, err := sameFile(src, info, dst)
	if err != nil {
		return fmt.Errorf("failed to compare source and destination files: %w", err)
	}
	if same {
		return nil
	}

	perm := info.Mode().Perm()
	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	if err := out.Chmod(perm); err != nil {
		_ = out.Close()
		return fmt.Errorf("failed to chmod destination file: %w", err)
	}
	if _, err := io.Copy(out, in); err != nil {
		_ = out.Close()
		return fmt.Errorf("failed to copy file contents: %w", err)
	}
	if err := out.Sync(); err != nil {
		_ = out.Close()
		return fmt.Errorf("failed to sync destination file: %w", err)
	}
	if err := out.Close(); err != nil {
		return fmt.Errorf("failed to close destination file: %w", err)
	}
	// Unconditional rather than gated on `dst` having been created: the contents
	// are already synced, so one more sync costs far less than deciding whether
	// the entry is new - which [os.Lstat] cannot answer for a dangling symlink.
	if err := syncDir(filepath.Dir(dst)); err != nil {
		return fmt.Errorf("failed to sync destination directory: %w", err)
	}
	return nil
}
