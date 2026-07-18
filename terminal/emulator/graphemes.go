package emulator

import "github.com/gechr/x/set"

// graphemeEmulators lists emulators known to measure text in grapheme
// clusters by default rather than per-codepoint wcwidth - whether via the
// mode 2027 / terminal-unicode-core proposal or their own segmentation, as
// with kitty, which never used wcwidth and rejected mode 2027 in favour of
// its text-sizing protocol.
var graphemeEmulators = set.New(
	Contour,
	Foot,
	Ghostty,
	Kitty,
	Rio,
	WezTerm,
	WindowsTerminal,
)

// SupportsGraphemes reports whether the detected terminal emulator is known
// to measure text in grapheme clusters rather than per-codepoint wcwidth.
// It returns false when the emulator cannot be determined.
func SupportsGraphemes() bool {
	return graphemeEmulators.Contains(Detect())
}
