package terminal

import (
	"os"
	"strings"
)

// SupportsTrueColor reports whether the terminal supports 24-bit "true color"
// output, based on the COLORTERM and TERM environment variables.
//
// COLORTERM=truecolor (or 24bit) is the de-facto signal modern terminals set.
// TERM alone is deliberately not trusted for capability - the ubiquitous
// TERM=xterm-256color advertises only 256 colors even on terminals that render
// 24-bit, which is precisely why COLORTERM exists. This reports capability, not
// preference: honoring NO_COLOR or a non-terminal stream is left to the caller.
func SupportsTrueColor() bool {
	switch strings.ToLower(os.Getenv("COLORTERM")) {
	case "truecolor", "24bit":
		return true
	}
	return strings.Contains(strings.ToLower(os.Getenv("TERM")), "truecolor")
}
