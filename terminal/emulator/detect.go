// Package emulator identifies the terminal emulator hosting the process.
package emulator

import (
	"os"
	"strings"
)

// Environment variables consulted by [Detect].
const (
	EnvTerm             = "TERM"
	EnvTermProgram      = "TERM_PROGRAM"
	EnvTerminalEmulator = "TERMINAL_EMULATOR"
)

// termPrograms maps lowercased `TERM_PROGRAM` values to emulator names.
var termPrograms = map[string]string{
	"alacritty":      Alacritty,
	"apple_terminal": AppleTerminal,
	"contour":        Contour,
	"ghostty":        Ghostty,
	"hyper":          Hyper,
	"iterm.app":      ITerm2,
	"mintty":         Mintty,
	"rio":            Rio,
	"tabby":          Tabby,
	"terminator":     Terminator,
	"vscode":         VSCode,
	"warpterminal":   Warp,
	"wezterm":        WezTerm,
	"zed":            Zed,
}

// muxVars lists multiplexer-specific environment variables. Multiplexers own
// the screen model of everything running inside them, so they are checked
// before anything else - including `TERM`, because tmux configs commonly set
// `TERM` to `screen-256color`.
//
// Neither tmux nor screen scrubs the other's variable, so when multiplexers
// are nested both markers are present and the environment cannot identify
// the innermost one; tmux is reported as the more common case. `TERM` cannot
// break the tie either: both muxes commonly use `screen-256color`.
var muxVars = []struct {
	env  string
	name string
}{
	{"TMUX", Tmux},
	{"STY", Screen},
}

// markerVars lists emulator-specific environment variables, checked in order.
var markerVars = []struct {
	env  string
	name string
}{
	{"ALACRITTY_WINDOW_ID", Alacritty},
	{"ALACRITTY_SOCKET", Alacritty},
	{"ALACRITTY_LOG", Alacritty},
	{"ConEmuPID", ConEmu},
	{"GHOSTTY_RESOURCES_DIR", Ghostty},
	{"GNOME_TERMINAL_SCREEN", GNOMETerminal},
	{"GNOME_TERMINAL_SERVICE", GNOMETerminal},
	{"ITERM_SESSION_ID", ITerm2},
	{"KITTY_WINDOW_ID", Kitty},
	{"KITTY_PID", Kitty},
	{"KONSOLE_VERSION", Konsole},
	{"KONSOLE_DBUS_SERVICE", Konsole},
	{"KONSOLE_DBUS_SESSION", Konsole},
	{"TERMINATOR_UUID", Terminator},
	{"TERMUX_VERSION", Termux},
	{"TILIX_ID", Tilix},
	{"WEZTERM_EXECUTABLE", WezTerm},
	{"WEZTERM_PANE", WezTerm},
	{"WT_SESSION", WindowsTerminal},
}

// termValues maps normalized `TERM` values to emulator names.
var termValues = map[string]string{
	"alacritty":      Alacritty,
	"contour":        Contour,
	"contour-latest": Contour,
	"foot":           Foot,
	"rio":            Rio,
	"rxvt-unicode":   URxvt,
	"screen":         Screen,
	"st":             ST,
	"tmux":           Tmux,
	"wezterm":        WezTerm,
	"xterm-ghostty":  Ghostty,
	"xterm-kitty":    Kitty,
}

// Detect returns the terminal emulator hosting the process, or empty if it
// cannot be determined. Detection is best-effort, based on environment
// variables inherited from the emulator.
// Priority: multiplexer variables, `TERM`, `TERM_PROGRAM`,
// `TERMINAL_EMULATOR`, emulator-specific variables. Multiplexers win because
// they own the screen model of everything inside them; `TERM` beats
// `TERM_PROGRAM` because the innermost emulator always sets it fresh for its
// own session, whereas `TERM_PROGRAM` and marker variables leak through from
// an outer terminal when one emulator is launched from another that does not
// scrub them (e.g. kitty launched from iTerm2 inherits both
// `TERM_PROGRAM=iTerm.app` and `ITERM_SESSION_ID`).
func Detect() string {
	for _, m := range muxVars {
		if os.Getenv(m.env) != "" {
			return m.name
		}
	}
	if name, ok := termValues[normalizeTerm(os.Getenv(EnvTerm))]; ok {
		return name
	}
	program := strings.ToLower(strings.TrimSpace(os.Getenv(EnvTermProgram)))
	if name, ok := termPrograms[program]; ok {
		return name
	}
	// JetBrains IDEs identify via `TERMINAL_EMULATOR`, e.g. `JetBrains-JediTerm`.
	if strings.HasPrefix(os.Getenv(EnvTerminalEmulator), "JetBrains") {
		return JetBrains
	}
	for _, m := range markerVars {
		if os.Getenv(m.env) != "" {
			return m.name
		}
	}
	return ""
}

// normalizeTerm strips color/variant suffixes so values like `st-256color`,
// `foot-extra`, and `alacritty-direct` match their base [termValues] entry.
func normalizeTerm(term string) string {
	for {
		base := term
		for _, suffix := range []string{"-256color", "-direct", "-extra"} {
			term = strings.TrimSuffix(term, suffix)
		}
		if term == base {
			return term
		}
	}
}
