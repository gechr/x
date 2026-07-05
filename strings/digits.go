package strings

// IsDigits reports whether `s` is non-empty and consists entirely of ASCII
// digits (0-9). An empty string is not digits.
func IsDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
