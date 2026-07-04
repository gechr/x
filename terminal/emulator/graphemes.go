package emulator

// graphemeEmulators lists emulators known to measure text in grapheme
// clusters (per the mode 2027 / terminal-unicode-core proposal) rather than
// per-codepoint wcwidth.
var graphemeEmulators = map[string]struct{}{
	Contour:         {},
	Foot:            {},
	Ghostty:         {},
	WezTerm:         {},
	WindowsTerminal: {},
}

// SupportsGraphemes reports whether the detected terminal emulator is known
// to measure text in grapheme clusters rather than per-codepoint wcwidth.
// It returns false when the emulator cannot be determined.
func SupportsGraphemes() bool {
	_, ok := graphemeEmulators[Detect()]
	return ok
}
