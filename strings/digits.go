package strings

// IsDigits reports whether `s` is non-empty and consists entirely of ASCII
// digits (0-9). An empty string is not digits.
func IsDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !IsDigitChar(c) {
			return false
		}
	}
	return true
}

// IsDigitChar reports whether `c` is an ASCII digit (0-9).
func IsDigitChar(c rune) bool {
	return c >= '0' && c <= '9'
}
