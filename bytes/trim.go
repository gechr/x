package bytes

import "bytes"

// TrimPrefixes returns `s` with the first matching prefix in `prefixes`
// removed. At most one prefix is removed; if none match, `s` is returned
// unchanged.
func TrimPrefixes(s []byte, prefixes ...[]byte) []byte {
	for _, p := range prefixes {
		if after, ok := bytes.CutPrefix(s, p); ok {
			return after
		}
	}
	return s
}

// TrimSuffixes returns `s` with the first matching suffix in `suffixes`
// removed. At most one suffix is removed; if none match, `s` is returned
// unchanged.
func TrimSuffixes(s []byte, suffixes ...[]byte) []byte {
	for _, suffix := range suffixes {
		if before, ok := bytes.CutSuffix(s, suffix); ok {
			return before
		}
	}
	return s
}
