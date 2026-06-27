//go:build unix && !darwin

package os

import (
	"errors"
	"fmt"
	"net/url"
	stdos "os"
	stdpath "path/filepath"
	"strconv"
	"syscall"
	"time"
)

// deletionDateFormat is the local, timezone-less timestamp the FreeDesktop trash
// specification records in a .trashinfo file.
const deletionDateFormat = "2006-01-02T15:04:05"

// trash moves path to a FreeDesktop.org trash directory: the home trash for a
// file on the same device as the home trash, or the matching top-directory trash
// ($top/.Trash/$uid or $top/.Trash-$uid) on the file's own device otherwise.
// The trash is always on the file's device, so the move is a plain rename that
// preserves every file type; a file with no usable same-device trash yields an
// error wrapping [errors.ErrUnsupported] rather than a cross-device copy.
func trash(path string) error {
	dev, err := device(path)
	if err != nil {
		return err
	}
	root, topDir, err := trashDir(path, dev)
	if err != nil {
		return err
	}
	return trashInto(path, root, topDir)
}

// trashDir resolves the trash root for a file on device dev, and the top
// directory its recorded path should be made relative to (empty for the home
// trash, which records an absolute path).
func trashDir(path string, dev uint64) (string, string, error) {
	home, err := homeTrash()
	if err != nil {
		return "", "", err
	}
	if hdev, derr := homeTrashDevice(home); derr == nil && hdev == dev {
		return home, "", nil
	}

	top := topDirOf(path, dev)
	root, err := topDirTrash(top)
	if err != nil {
		return "", "", err
	}
	return root, top, nil
}

// homeTrash returns $XDG_DATA_HOME/Trash, defaulting to ~/.local/share/Trash. A
// relative XDG_DATA_HOME is ignored, per the XDG base-directory specification.
func homeTrash() (string, error) {
	if dataHome := stdos.Getenv("XDG_DATA_HOME"); stdpath.IsAbs(dataHome) {
		return stdpath.Join(dataHome, "Trash"), nil
	}
	home, err := stdos.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to locate home directory: %w", err)
	}
	return stdpath.Join(home, ".local", "share", "Trash"), nil
}

// homeTrashDevice returns the device the home trash lives on, taken from the
// nearest existing ancestor of the trash root (the Trash dir itself may not exist
// yet). This reflects XDG_DATA_HOME, which may be on a different device than $HOME.
func homeTrashDevice(root string) (uint64, error) {
	for dir := root; ; {
		if dev, err := device(dir); err == nil {
			return dev, nil
		}
		parent := stdpath.Dir(dir)
		if parent == dir {
			return 0, fmt.Errorf("cannot determine device for %q", root)
		}
		dir = parent
	}
}

// device returns the ID of the device path resides on.
func device(path string) (uint64, error) {
	info, err := stdos.Lstat(path)
	if err != nil {
		return 0, err
	}
	st, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, fmt.Errorf("cannot determine device for %q", path)
	}
	return uint64(st.Dev), nil //nolint:unconvert // Dev is int32 on some platforms
}

// topDirOf returns the mount point path is under: the highest ancestor still on
// device dev, found by walking up until the device changes or root is reached.
//
// This locates a mount point by device boundary, which misses mounts that share
// a device ID with their parent (bind mounts, btrfs subvolumes); such a file is
// trashed to a higher directory's trash on the same device, still recoverable.
func topDirOf(path string, dev uint64) string {
	dir := stdpath.Dir(path)
	for {
		parent := stdpath.Dir(dir)
		if parent == dir {
			return dir
		}
		if pdev, err := device(parent); err != nil || pdev != dev {
			return dir
		}
		dir = parent
	}
}

// topDirTrash returns a validated trash root on the mount point top, preferring
// an admin-created $top/.Trash/$uid (under a sticky, non-symlink .Trash, per the
// spec) and otherwise creating $top/.Trash-$uid.
func topDirTrash(top string) (string, error) {
	uid := stdos.Getuid()
	name := strconv.Itoa(uid)

	shared := stdpath.Join(top, ".Trash")
	if info, err := stdos.Lstat(shared); err == nil &&
		info.IsDir() && info.Mode()&stdos.ModeSticky != 0 {
		root := stdpath.Join(shared, name)
		if err := stdos.MkdirAll(root, 0o700); err == nil {
			if err := validateTrashRoot(root, uid); err == nil {
				return root, nil
			}
		}
	}

	root := stdpath.Join(top, ".Trash-"+name)
	if err := stdos.MkdirAll(root, 0o700); err != nil {
		return "", fmt.Errorf(
			"failed to create trash directory: %w: %w",
			err,
			errors.ErrUnsupported,
		)
	}
	if err := validateTrashRoot(root, uid); err != nil {
		return "", err
	}
	return root, nil
}

// validateTrashRoot rejects an unsafe top-directory trash root: the spec is
// strict here because these live on shared or removable mounts. The root must be
// a real directory (not a planted symlink), owned by the current user, and not
// accessible to others. A rejection wraps [errors.ErrUnsupported].
func validateTrashRoot(root string, uid int) error {
	info, err := stdos.Lstat(root)
	if err != nil {
		return fmt.Errorf("failed to stat trash directory: %w", err)
	}
	switch {
	case info.Mode()&stdos.ModeSymlink != 0 || !info.IsDir():
		return fmt.Errorf(
			"trash directory %q is not a real directory: %w",
			root,
			errors.ErrUnsupported,
		)
	case info.Mode().Perm()&0o077 != 0:
		return fmt.Errorf(
			"trash directory %q is accessible to other users: %w",
			root,
			errors.ErrUnsupported,
		)
	}
	if st, ok := info.Sys().(*syscall.Stat_t); ok && int(st.Uid) != uid {
		return fmt.Errorf(
			"trash directory %q is not owned by the current user: %w",
			root,
			errors.ErrUnsupported,
		)
	}
	return nil
}

// trashInto writes path's .trashinfo record and moves it into root's files/
// directory under a name unique within that trash. topDir, when set, is the
// mount point the recorded path is made relative to.
func trashInto(path, root, topDir string) error {
	filesDir := stdpath.Join(root, "files")
	infoDir := stdpath.Join(root, "info")
	for _, dir := range []string{filesDir, infoDir} {
		if err := stdos.MkdirAll(dir, 0o700); err != nil {
			return fmt.Errorf("failed to create trash directory: %w", err)
		}
	}

	name, infoFile, err := claimName(infoDir, filesDir, stdpath.Base(path))
	if err != nil {
		return err
	}

	if err := stdos.WriteFile(infoFile, []byte(trashInfo(path, topDir)), 0o600); err != nil {
		return fmt.Errorf("failed to write trash info: %w", err)
	}
	if err := move(path, stdpath.Join(filesDir, name)); err != nil {
		_ = stdos.Remove(infoFile) // release the reserved name on failure
		return err
	}
	return nil
}

// trashInfo renders a .trashinfo record. The stored path is absolute for the
// home trash and relative to topDir for a top-directory trash, URL-encoded with
// path separators preserved.
func trashInfo(path, topDir string) string {
	stored := path
	if topDir != "" {
		if rel, err := stdpath.Rel(topDir, path); err == nil {
			stored = rel
		}
	}
	return fmt.Sprintf(
		"[Trash Info]\nPath=%s\nDeletionDate=%s\n",
		(&url.URL{Path: stored}).EscapedPath(),
		time.Now().Format(deletionDateFormat),
	)
}

// maxTrashNameAttempts bounds the search for a free name in a trash directory.
const maxTrashNameAttempts = 1 << 16

// claimName reserves a name unique within a trash, appending _N on collision. It
// skips a name whose files/ slot is already taken and atomically creates the
// info/<name>.trashinfo (O_EXCL) to reserve the rest, so a reader never sees a
// trashed file lacking its record and the move does not clobber another entry.
func claimName(infoDir, filesDir, base string) (string, string, error) {
	for i := range maxTrashNameAttempts {
		name := base
		if i > 0 {
			name = fmt.Sprintf("%s_%d", base, i)
		}
		if _, err := stdos.Lstat(stdpath.Join(filesDir, name)); err == nil {
			continue // files/<name> already occupied; try the next candidate
		}
		infoFile := stdpath.Join(infoDir, name+".trashinfo")
		f, err := stdos.OpenFile(infoFile, stdos.O_CREATE|stdos.O_WRONLY|stdos.O_EXCL, 0o600)
		if err == nil {
			_ = f.Close()
			return name, infoFile, nil
		}
		if !errors.Is(err, stdos.ErrExist) {
			return "", "", fmt.Errorf("failed to reserve trash info: %w", err)
		}
	}
	return "", "", fmt.Errorf("no free name for %q in trash", base)
}

// move renames src to dst. The trash root was chosen on src's own device, so a
// cross-device rename should not arise; if it does, it is reported as unsupported
// rather than silently copying across the boundary (which a caller may prefer to
// handle, e.g. by deleting instead).
func move(src, dst string) error {
	switch err := stdos.Rename(src, dst); {
	case err == nil:
		return nil
	case errors.Is(err, syscall.EXDEV):
		return fmt.Errorf("cannot trash %q across devices: %w", src, errors.ErrUnsupported)
	default:
		return fmt.Errorf("failed to move to trash: %w", err)
	}
}
