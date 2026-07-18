//go:build windows

package os

import (
	"os"
	"path/filepath"

	xslices "github.com/gechr/x/slices"
	xstrings "github.com/gechr/x/strings"
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
	return xslices.ContainsFold(xstrings.SplitBy(exts, ";"), ext)
}
