package os

import (
	stdos "os"

	xfilepath "github.com/gechr/x/filepath"
)

// SameFile reports whether a and b identify the same file. Missing leaf paths
// are compared after resolving their parent directories, and existing files are
// compared with os.SameFile to detect hard links.
func SameFile(a, b string) (bool, error) {
	return sameFile(a, nil, b)
}

func sameFile(a string, aInfo stdos.FileInfo, b string) (bool, error) {
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
	return stdos.SameFile(aInfo, bInfo), nil
}

func sameFileInfo(path string) (stdos.FileInfo, bool) {
	info, err := stdos.Stat(path)
	return info, err == nil
}
