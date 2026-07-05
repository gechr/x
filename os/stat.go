package os

import (
	"errors"
	"os"
	"syscall"
)

// notExist reports whether `err` means the path does not exist, including
// ENOTDIR (a non-directory component partway through the path).
func notExist(err error) bool {
	return errors.Is(err, os.ErrNotExist) || errors.Is(err, syscall.ENOTDIR)
}

// stat returns the FileInfo for `path`, or (nil, nil) if it does not exist.
func stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if notExist(err) {
		return nil, nil //nolint:nilnil // fine for an internal helper
	}
	return info, err
}

// Exists reports whether `path` exists.
func Exists(path string) (bool, error) {
	info, err := stat(path)
	return info != nil, err
}

// IsFile reports whether `path` is a regular file.
func IsFile(path string) (bool, error) {
	info, err := stat(path)
	return info != nil && info.Mode().IsRegular(), err
}

// IsDir reports whether `path` is a directory.
func IsDir(path string) (bool, error) {
	info, err := stat(path)
	return info != nil && info.IsDir(), err
}

// IsSymlink reports whether `path` is a symbolic link.
func IsSymlink(path string) (bool, error) {
	info, err := os.Lstat(path)
	switch {
	case notExist(err):
		return false, nil
	case err != nil:
		return false, err
	default:
		return info.Mode()&os.ModeSymlink != 0, nil
	}
}

// IsWritableDir reports whether `dir` exists and the current process can create
// files in it. Uses a probe file rather than permission-bit inspection so that
// ACLs and immutable mounts are handled correctly.
func IsWritableDir(dir string) bool {
	ok, err := IsDir(dir)
	if err != nil || !ok {
		return false
	}
	tmp, err := os.CreateTemp(dir, ".x-writable-check-*")
	if err != nil {
		return false
	}
	name := tmp.Name()
	_ = tmp.Close()
	_ = os.Remove(name)
	return true
}
