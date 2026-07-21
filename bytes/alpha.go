package bytes

import "unicode/utf8"

// IsASCII reports whether `s` is non-empty and consists entirely of ASCII
// characters (code points 0-127). An empty slice is not ASCII.
func IsASCII(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// IsAlpha reports whether `s` is non-empty and consists entirely of ASCII
// letters (a-z, A-Z). An empty slice is not alpha.
func IsAlpha(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsAlphaChar(c) {
			return false
		}
	}
	return true
}

// IsAlphaChar reports whether `c` is an ASCII letter (a-z, A-Z).
func IsAlphaChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// IsAlphanumeric reports whether `s` is non-empty and consists entirely of
// ASCII letters (a-z, A-Z) or digits (0-9). An empty slice is not
// alphanumeric.
func IsAlphanumeric(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsAlphanumericChar(c) {
			return false
		}
	}
	return true
}

// IsAlphanumericChar reports whether `c` is an ASCII letter (a-z, A-Z) or digit
// (0-9).
func IsAlphanumericChar(c byte) bool {
	return IsAlphaChar(c) || IsDigitChar(c)
}
