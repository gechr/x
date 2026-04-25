package fs

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// stat returns the FileInfo for path, or (nil, nil) if it does not exist.
func stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil //nolint:nilnil // fine for an internal helper
	}
	return info, err
}

// Exists reports whether path exists.
func Exists(path string) (bool, error) {
	info, err := stat(path)
	return info != nil, err
}

// IsFile reports whether path is a regular file.
func IsFile(path string) (bool, error) {
	info, err := stat(path)
	return info != nil && info.Mode().IsRegular(), err
}

// IsDir reports whether path is a directory.
func IsDir(path string) (bool, error) {
	info, err := stat(path)
	return info != nil && info.IsDir(), err
}

// IsSymlink reports whether path is a symbolic link.
func IsSymlink(path string) (bool, error) {
	info, err := os.Lstat(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		return false, err
	default:
		return info.Mode()&os.ModeSymlink != 0, nil
	}
}

// Resolve recursively follows every symlink along path and returns the fully
// resolved absolute path. On any error (missing component, cycle, permission)
// the input path is returned alongside the error so callers can choose whether
// to handle it or fall back.
func Resolve(path string) (string, error) {
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path, err
	}
	abs, err := filepath.Abs(resolved)
	if err != nil {
		return resolved, err
	}
	return abs, nil
}

// IsWritableDir reports whether dir exists and the current process can create
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

// IsWithin reports whether all target paths are equal to or contained within
// base. Returns false when no targets are provided.
//
// Example:
//
//	IsWithin("src", "src/foo.go")             // true
//	IsWithin(".", "src/foo.go", "lib/bar.go") // true
//	IsWithin("src", "lib/foo.go")             // false
func IsWithin(base string, targets ...string) bool {
	if len(targets) == 0 {
		return false
	}
	absBase, err := filepath.Abs(base)
	if err != nil {
		return false
	}
	prefix := absBase
	if !strings.HasSuffix(prefix, string(filepath.Separator)) {
		prefix += string(filepath.Separator)
	}
	for _, target := range targets {
		absTarget, err := filepath.Abs(target)
		if err != nil {
			return false
		}
		if absTarget != absBase && !strings.HasPrefix(absTarget, prefix) {
			return false
		}
	}
	return true
}
