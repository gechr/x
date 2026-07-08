package strings

import (
	"slices"
	"strings"
)

// IsBlank reports whether `s` is empty or consists only of whitespace.
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// AnyEmpty reports whether any of the given strings is empty.
func AnyEmpty(values ...string) bool {
	return slices.Contains(values, "")
}

// AnyNonEmpty reports whether any of the given strings is non-empty.
func AnyNonEmpty(values ...string) bool {
	for _, value := range values {
		if value != "" {
			return true
		}
	}
	return false
}

// AllEmpty reports whether every given string is empty.
func AllEmpty(values ...string) bool {
	for _, value := range values {
		if value != "" {
			return false
		}
	}
	return true
}

// AllNonEmpty reports whether every given string is non-empty.
func AllNonEmpty(values ...string) bool {
	return !slices.Contains(values, "")
}
