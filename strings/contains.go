package strings

import "strings"

// ContainsAll reports whether s contains all of the given substrings.
func ContainsAll(s string, substrings ...string) bool {
	for _, ss := range substrings {
		if !strings.Contains(s, ss) {
			return false
		}
	}
	return true
}

// ContainsAny reports whether s contains any of the given substrings.
func ContainsAny(s string, substrings ...string) bool {
	for _, ss := range substrings {
		if strings.Contains(s, ss) {
			return true
		}
	}
	return false
}
