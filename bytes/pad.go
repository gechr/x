package bytes

import (
	"bytes"
	"unicode/utf8"
)

// PadLeft pads `s` with spaces on the left to `width` runes, right-aligning it.
// Slices already `width` runes or longer are returned unchanged. Width is
// counted in runes; for display-width-aware handling of ANSI text use the
// [github.com/gechr/x/ansi] package. The returned slice never aliases `s`.
//
//	PadLeft([]byte("hi"), 5) // "   hi"
func PadLeft(s []byte, width int) []byte {
	pad := padding(s, width)
	out := make([]byte, 0, pad+len(s))
	out = append(out, bytes.Repeat([]byte(" "), pad)...)
	return append(out, s...)
}

// PadRight pads `s` with spaces on the right to `width` runes, left-aligning it.
// Slices already `width` runes or longer are returned unchanged. The returned
// slice never aliases `s`.
//
//	PadRight([]byte("hi"), 5) // "hi   "
func PadRight(s []byte, width int) []byte {
	pad := padding(s, width)
	out := make([]byte, 0, len(s)+pad)
	out = append(out, s...)
	return append(out, bytes.Repeat([]byte(" "), pad)...)
}

// PadCenter pads `s` with spaces on both sides to `width` runes, centring it.
// An odd rune of padding goes on the right. Slices already `width` runes or
// longer are returned unchanged. The returned slice never aliases `s`.
//
//	PadCenter([]byte("hi"), 5) // " hi  "
func PadCenter(s []byte, width int) []byte {
	pad := padding(s, width)
	left := pad / 2 //nolint:mnd // half the padding goes left
	out := make([]byte, 0, len(s)+pad)
	out = append(out, bytes.Repeat([]byte(" "), left)...)
	out = append(out, s...)
	return append(out, bytes.Repeat([]byte(" "), pad-left)...)
}

// padding returns the number of pad runes needed to bring `s` to `width`, never
// negative.
func padding(s []byte, width int) int {
	return max(0, width-utf8.RuneCount(s))
}
