package strings

import "strings"

// IsBlank reports whether `s` is empty or consists only of whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
