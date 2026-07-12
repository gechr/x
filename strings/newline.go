package strings

import "strings"

// EnsureTrailingNewline trims any trailing newlines from `s` and appends exactly
// one, so the result always ends in a single `\n`. An empty string becomes
// `\n`.
func EnsureTrailingNewline(s string) string {
	return strings.TrimRight(s, "\n") + "\n"
}
