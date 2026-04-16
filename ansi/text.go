package ansi

import xansi "github.com/charmbracelet/x/ansi"

// Method controls how display width is calculated.
type Method = xansi.Method

// Display width methods.
const (
	WcWidth       = xansi.WcWidth
	GraphemeWidth = xansi.GraphemeWidth
)

// Strip removes ANSI escape codes from a string.
func Strip(s string) string {
	return xansi.Strip(s)
}

// StringWidth returns the display width of a string in cells,
// ignoring ANSI escape codes and accounting for wide characters.
// Uses grapheme clustering.
func StringWidth(s string) int {
	return xansi.StringWidth(s)
}

// Truncate truncates a string to a given cell width, appending tail
// if the string was truncated. ANSI escape codes are preserved.
func Truncate(s string, length int, tail string) string {
	return xansi.Truncate(s, length, tail)
}
