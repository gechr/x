package strings

// IsHex reports whether s consists entirely of hexadecimal digits.
// An empty string returns true.
func IsHex(s string) bool {
	for _, c := range s {
		if !IsHexChar(c) {
			return false
		}
	}
	return true
}

// IsHexChar reports whether c is a valid hexadecimal digit (0-9, a-f, A-F).
func IsHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}
