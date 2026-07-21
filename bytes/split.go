package bytes

import "bytes"

// SplitBy splits `s` by `sep`, trims whitespace from each part, and drops empty
// values.
func SplitBy(s, sep []byte) [][]byte {
	raw := bytes.Split(bytes.TrimSpace(s), sep)
	parts := make([][]byte, 0, len(raw))
	for _, part := range raw {
		part = bytes.TrimSpace(part)
		if len(part) != 0 {
			parts = append(parts, part)
		}
	}
	return parts
}
