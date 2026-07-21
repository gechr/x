package bytes

import "bytes"

// ContainsAll reports whether `s` contains all of the given `subslices`.
func ContainsAll(s []byte, subslices ...[]byte) bool {
	for _, sub := range subslices {
		if !bytes.Contains(s, sub) {
			return false
		}
	}
	return true
}

// ContainsAny reports whether `s` contains any of the given `subslices`.
func ContainsAny(s []byte, subslices ...[]byte) bool {
	for _, sub := range subslices {
		if bytes.Contains(s, sub) {
			return true
		}
	}
	return false
}
