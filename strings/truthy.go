package strings

import "strings"

// IsTruthy reports whether `s` is an affirmative boolean token: `1`, `true`,
// `yes`, or `on`, case-insensitively and ignoring surrounding whitespace.
// IsTruthy is not the complement of [IsFalsy]: a value like `banana` is
// neither truthy nor falsy.
func IsTruthy(s string) bool {
	switch normalizeToken(s) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

// IsFalsy reports whether `s` is a negative boolean token: `0`, `false`,
// `no`, or `off`, case-insensitively and ignoring surrounding whitespace.
// The empty string is not falsy: it signals absence, not refusal - use
// IsFalsy alongside a presence check when the distinction matters.
func IsFalsy(s string) bool {
	switch normalizeToken(s) {
	case "0", "false", "no", "off":
		return true
	default:
		return false
	}
}

// normalizeToken lowercases and trims a value for token comparison.
func normalizeToken(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
