// Package os provides OS helpers: file probes, safe writes, copy, line I/O, and
// platform/architecture detection.
package os

import (
	"os"

	xfilepath "github.com/gechr/x/filepath"
)

// SameFile reports whether `a` and `b` identify the same file. Missing leaf
// paths are compared after resolving their parent directories, and existing
// files are compared with [os.SameFile] to detect hard links.
func SameFile(a, b string) (bool, error) {
	return sameFile(a, nil, b)
}

// RemoveIfExists removes `path`. It succeeds without error if `path` does not
// exist.
func RemoveIfExists(path string) error {
	err := os.Remove(path)
	if notExist(err) {
		return nil
	}
	return err
}

func sameFile(a string, aInfo os.FileInfo, b string) (bool, error) {
	aResolved, err := xfilepath.ResolveLenient(a)
	if err != nil {
		return false, err
	}
	bResolved, err := xfilepath.ResolveLenient(b)
	if err != nil {
		return false, err
	}
	if aResolved == bResolved {
		return true, nil
	}
	if aInfo == nil {
		var ok bool
		aInfo, ok = sameFileInfo(a)
		if !ok {
			return false, nil
		}
	}
	bInfo, ok := sameFileInfo(b)
	if !ok {
		return false, nil
	}
	return os.SameFile(aInfo, bInfo), nil
}

func sameFileInfo(path string) (os.FileInfo, bool) {
	info, err := os.Stat(path)
	return info, err == nil
}
