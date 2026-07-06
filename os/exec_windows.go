//go:build windows

package os

import (
	"os"
	"path/filepath"
	"strings"
)

// executable reports whether Windows would run `path` as a program. Windows has
// no execute permission bit, so - as os/exec.LookPath does - executability is
// decided by the file's extension appearing in `%PATHEXT%` (falling back to the
// documented default set when the variable is unset).
func executable(path string) bool {
	exts := os.Getenv("PATHEXT")
	if exts == "" {
		exts = ".COM;.EXE;.BAT;.CMD"
	}
	ext := filepath.Ext(path)
	if ext == "" {
		return false
	}
	for _, e := range strings.Split(exts, ";") {
		if strings.EqualFold(strings.TrimSpace(e), ext) {
			return true
		}
	}
	return false
}
