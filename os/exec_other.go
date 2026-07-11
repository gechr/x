//go:build !unix && !windows

package os

import "os"

// executable falls back to the file's permission bits on platforms outside Unix
// and Windows (e.g. Plan 9, WASM), which offer neither `access(2)` nor a
// `%PATHEXT%` analogue.
func executable(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode()&0o111 != 0
}
