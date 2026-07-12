package strings

import "strings"

// TrimPrefixes returns `s` with the first matching prefix in `prefixes`
// removed. At most one prefix is removed; if none match, `s` is returned
// unchanged.
func TrimPrefixes(s string, prefixes ...string) string {
	for _, p := range prefixes {
		if after, ok := strings.CutPrefix(s, p); ok {
			return after
		}
	}
	return s
}

// TrimSuffixes returns `s` with the first matching suffix in `suffixes`
// removed. At most one suffix is removed; if none match, `s` is returned
// unchanged.
func TrimSuffixes(s string, suffixes ...string) string {
	for _, suffix := range suffixes {
		if before, ok := strings.CutSuffix(s, suffix); ok {
			return before
		}
	}
	return s
}
