package bytes

import "bytes"

// IsBlank reports whether `s` is empty or consists only of whitespace.
func IsBlank(s []byte) bool {
	return len(bytes.TrimSpace(s)) == 0
}

// AnyEmpty reports whether any of the given slices is empty.
func AnyEmpty(values ...[]byte) bool {
	for _, value := range values {
		if len(value) == 0 {
			return true
		}
	}
	return false
}

// AnyNonEmpty reports whether any of the given slices is non-empty.
func AnyNonEmpty(values ...[]byte) bool {
	for _, value := range values {
		if len(value) != 0 {
			return true
		}
	}
	return false
}

// AllEmpty reports whether every given slice is empty.
func AllEmpty(values ...[]byte) bool {
	for _, value := range values {
		if len(value) != 0 {
			return false
		}
	}
	return true
}

// AllNonEmpty reports whether every given slice is non-empty.
func AllNonEmpty(values ...[]byte) bool {
	for _, value := range values {
		if len(value) == 0 {
			return false
		}
	}
	return true
}
