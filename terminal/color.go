package terminal

import (
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/charmbracelet/colorprofile"
)

// preferenceVars are the environment variables that state whether the user
// wants color rather than whether the terminal can render it. They are removed
// before detection so [SupportsTrueColor] keeps reporting capability alone.
var preferenceVars = []string{"NO_COLOR", "CLICOLOR", "CLICOLOR_FORCE"}

// trueColor caches detection for the process, which reads the terminfo
// database and, inside tmux, runs `tmux info`.
var trueColor = sync.OnceValue(func() bool {
	return detectTrueColor(os.Environ())
})

// SupportsTrueColor reports whether the terminal supports 24-bit "true color"
// output.
//
// Detection reads COLORTERM, TERM, the terminfo Tc and RGB capabilities and,
// inside tmux, `tmux info` - tmux does not forward COLORTERM, so only its own
// capabilities settle the question there. TERM alone is not trusted for
// capability: the ubiquitous TERM=xterm-256color advertises only 256 colors
// even on terminals that render 24-bit, which is precisely why COLORTERM
// exists.
//
// This reports capability, not preference: NO_COLOR and CLICOLOR are ignored
// and a redirected stream is still measured against the terminal, so honoring
// either is left to the caller. The first call detects, and the result is
// cached for the process.
func SupportsTrueColor() bool {
	return trueColor()
}

// detectTrueColor reports true-color support for `env`, the uncached core of
// [SupportsTrueColor].
func detectTrueColor(env []string) bool {
	// The output stream is deliberately not passed: TTY_FORCE already makes
	// this a capability check, so there is nothing to test it for.
	if colorprofile.Detect(nil, capabilityEnv(env)) == colorprofile.TrueColor {
		return true
	}
	// colorprofile matches TERM by name and suffix, so a TERM that advertises
	// true color in its own words - with no terminfo entry to confirm it -
	// still needs the substring this package has always accepted.
	return strings.Contains(strings.ToLower(lookupEnv(env, "TERM")), "truecolor")
}

// capabilityEnv returns `env` with the color-preference variables removed and
// TTY_FORCE set, leaving colorprofile to report what the terminal can render
// whether or not output is a terminal and whether or not color was asked for.
func capabilityEnv(env []string) []string {
	capability := make([]string, 0, len(env)+1)
	for _, entry := range env {
		if name, _, ok := strings.Cut(entry, "="); ok && slices.Contains(preferenceVars, name) {
			continue
		}
		capability = append(capability, entry)
	}
	return append(capability, "TTY_FORCE=1")
}

// lookupEnv returns the value `env` gives `name`, or "" when it is unset. The
// last entry wins, matching how colorprofile reads the same slice.
func lookupEnv(env []string, name string) string {
	var value string
	for _, entry := range env {
		if n, v, ok := strings.Cut(entry, "="); ok && n == name {
			value = v
		}
	}
	return value
}
