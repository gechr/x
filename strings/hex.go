package strings

// IsHex reports whether `s` is non-empty and consists entirely of hexadecimal
// digits. An empty string is not hex.
func IsHex(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !IsHexChar(c) {
			return false
		}
	}
	return true
}

// IsHexChar reports whether `c` is a valid hexadecimal digit (0-9, a-f, A-F).
func IsHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// IsGitCommit reports whether `s` is 40 hexadecimal digits (a Git commit hash).
func IsGitCommit(s string) bool {
	return len(s) == 40 && IsHex(s)
}

// IsSHA256 reports whether `s` is 64 hexadecimal digits (a SHA-256 digest).
func IsSHA256(s string) bool {
	return len(s) == 64 && IsHex(s)
}
