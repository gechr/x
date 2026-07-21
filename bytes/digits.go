package bytes

// IsDigits reports whether `s` is non-empty and consists entirely of ASCII
// digits (0-9). An empty slice is not digits.
func IsDigits(s []byte) bool {
	if len(s) == 0 {
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
func IsDigitChar(c byte) bool {
	return c >= '0' && c <= '9'
}
