package strings

import "github.com/gechr/x/internal/fold"

// CompareFold compares `a` and `b` case-insensitively, using the same simple
// case-folding as [strings.EqualFold], and returns -1, 0, or 1 following the
// [cmp.Compare] convention. `CompareFold(a, b) == 0` iff
// `strings.EqualFold(a, b)`.
func CompareFold(a, b string) int {
	return fold.Compare(a, b)
}

// ContainsFold reports whether `s` contains `substr`, case-insensitively using
// the same simple case-folding as [strings.EqualFold].
func ContainsFold(s, substr string) bool {
	return fold.Contains(s, substr)
}

// HasPrefixFold reports whether `s` begins with `prefix`, case-insensitively
// using the same simple case-folding as [strings.EqualFold].
func HasPrefixFold(s, prefix string) bool {
	return fold.HasPrefix(s, prefix)
}

// HasSuffixFold reports whether `s` ends with `suffix`, case-insensitively
// using the same simple case-folding as [strings.EqualFold].
func HasSuffixFold(s, suffix string) bool {
	return fold.HasSuffix(s, suffix)
}
