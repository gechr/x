package bytes

import "github.com/gechr/x/internal/fold"

// CompareFold compares `a` and `b` case-insensitively, using the same simple
// case-folding as [bytes.EqualFold], and returns -1, 0, or 1 following the
// [cmp.Compare] convention. `CompareFold(a, b) == 0` iff
// `bytes.EqualFold(a, b)`.
func CompareFold(a, b []byte) int {
	return fold.Compare(a, b)
}

// ContainsFold reports whether `s` contains `subslice`, case-insensitively
// using the same simple case-folding as [bytes.EqualFold].
func ContainsFold(s, subslice []byte) bool {
	return fold.Contains(s, subslice)
}

// HasPrefixFold reports whether `s` begins with `prefix`, case-insensitively
// using the same simple case-folding as [bytes.EqualFold].
func HasPrefixFold(s, prefix []byte) bool {
	return fold.HasPrefix(s, prefix)
}

// HasSuffixFold reports whether `s` ends with `suffix`, case-insensitively
// using the same simple case-folding as [bytes.EqualFold].
func HasSuffixFold(s, suffix []byte) bool {
	return fold.HasSuffix(s, suffix)
}
