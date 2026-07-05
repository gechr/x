package strings

import "strings"

// Unwrap returns `s` with the leading `prefix` and trailing `suffix` removed and
// reports whether both were present. Unlike a [strings.TrimPrefix] +
// [strings.TrimSuffix] chain, nothing is removed unless `s` starts with `prefix`
// AND ends with `suffix`, so a one-sided match is returned unchanged.
func Unwrap(s, prefix, suffix string) (string, bool) {
	if len(s) < len(prefix)+len(suffix) ||
		!strings.HasPrefix(s, prefix) || !strings.HasSuffix(s, suffix) {
		return s, false
	}
	return s[len(prefix) : len(s)-len(suffix)], true
}
