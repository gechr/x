package strings

import "strings"

// SplitLines splits s into non-empty trimmed lines.
func SplitLines(s string) []string {
	return SplitBy(s, "\n")
}

// SplitLinesRaw splits s into lines losslessly, normalizing CRLF to LF: every
// line is kept verbatim - empty lines and the trailing empty element included -
// so the result joins back with "\n" without losing content or line numbers.
func SplitLinesRaw(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
