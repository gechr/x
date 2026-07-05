package os

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// AtomicWrite writes `data` to `path` via a temp-file-and-rename in the same
// directory. The temp file is removed on any failure.
func AtomicWrite(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
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
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		cleanup()
		return fmt.Errorf("failed to sync temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		cleanup()
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		cleanup()
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	return nil
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
// before close. When `src` and `dst` are the same file (including via hard link)
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
	return nil
}
