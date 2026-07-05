package strings

import (
	"strings"
	"unicode/utf8"
)

// PadLeft pads `s` with spaces on the left to `width` runes, right-aligning it.
// Strings already `width` runes or longer are returned unchanged. Width is
// counted in runes; for display-width-aware handling of ANSI text use the
// [github.com/gechr/x/ansi] package.
//
//	PadLeft("hi", 5) // "   hi"
func PadLeft(s string, width int) string {
	return strings.Repeat(" ", padding(s, width)) + s
}

// PadRight pads `s` with spaces on the right to `width` runes, left-aligning it.
// Strings already `width` runes or longer are returned unchanged.
//
//	PadRight("hi", 5) // "hi   "
func PadRight(s string, width int) string {
	return s + strings.Repeat(" ", padding(s, width))
}

// PadCenter pads `s` with spaces on both sides to `width` runes, centring it.
// An odd rune of padding goes on the right. Strings already `width` runes or
// longer are returned unchanged.
//
//	PadCenter("hi", 5) // " hi  "
func PadCenter(s string, width int) string {
	pad := padding(s, width)
	left := pad / 2 //nolint:mnd // half the padding goes left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", pad-left)
}

// padding returns the number of pad runes needed to bring `s` to `width`, never
// negative.
func padding(s string, width int) int {
	return max(0, width-utf8.RuneCountInString(s))
}
