package emulator

import "slices"

// Recognized terminal emulator names, as returned by Detect.
const (
	Alacritty       = "alacritty"
	AppleTerminal   = "apple-terminal"
	ConEmu          = "conemu"
	Contour         = "contour"
	Foot            = "foot"
	Ghostty         = "ghostty"
	GNOMETerminal   = "gnome-terminal"
	Hyper           = "hyper"
	ITerm2          = "iterm2"
	JetBrains       = "jetbrains"
	Kitty           = "kitty"
	Konsole         = "konsole"
	Mintty          = "mintty"
	Rio             = "rio"
	Screen          = "screen"
	ST              = "st"
	Tabby           = "tabby"
	Terminator      = "terminator"
	Termux          = "termux"
	Tilix           = "tilix"
	Tmux            = "tmux"
	URxvt           = "urxvt"
	VSCode          = "vscode"
	Warp            = "warp"
	WezTerm         = "wezterm"
	WindowsTerminal = "windows-terminal"
	Zed             = "zed"
)

var knownEmulators = []string{
	Alacritty,
	AppleTerminal,
	ConEmu,
	Contour,
	Foot,
	Ghostty,
	GNOMETerminal,
	Hyper,
	ITerm2,
	JetBrains,
	Kitty,
	Konsole,
	Mintty,
	Rio,
	Screen,
	ST,
	Tabby,
	Terminator,
	Termux,
	Tilix,
	Tmux,
	URxvt,
	VSCode,
	Warp,
	WezTerm,
	WindowsTerminal,
	Zed,
}

var knownEmulatorSet = map[string]struct{}{
	Alacritty:       {},
	AppleTerminal:   {},
	ConEmu:          {},
	Contour:         {},
	Foot:            {},
	Ghostty:         {},
	GNOMETerminal:   {},
	Hyper:           {},
	ITerm2:          {},
	JetBrains:       {},
	Kitty:           {},
	Konsole:         {},
	Mintty:          {},
	Rio:             {},
	Screen:          {},
	ST:              {},
	Tabby:           {},
	Terminator:      {},
	Termux:          {},
	Tilix:           {},
	Tmux:            {},
	URxvt:           {},
	VSCode:          {},
	Warp:            {},
	WezTerm:         {},
	WindowsTerminal: {},
	Zed:             {},
}

// Known returns the set of recognized terminal emulator names.
func Known() []string {
	return slices.Clone(knownEmulators)
}

// IsKnown reports whether name matches a known terminal emulator.
func IsKnown(name string) bool {
	_, ok := knownEmulatorSet[name]
	return ok
}
