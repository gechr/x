package strings

import (
	"cmp"
	"unicode"
	"unicode/utf8"
)

// CompareFold compares `a` and `b` case-insensitively, using the same simple
// case-folding as [strings.EqualFold], and returns -1, 0, or 1 following the
// [cmp.Compare] convention. `CompareFold(a, b) == 0` iff
// `strings.EqualFold(a, b)`.
func CompareFold(a, b string) int {
	for a != "" && b != "" {
		ra, na := utf8.DecodeRuneInString(a)
		rb, nb := utf8.DecodeRuneInString(b)
		a, b = a[na:], b[nb:]
		if c := cmp.Compare(foldRune(ra), foldRune(rb)); c != 0 {
			return c
		}
	}
	return cmp.Compare(len(a), len(b))
}

// foldRune maps `r` to the canonical (smallest) member of its case-fold orbit,
// so two runes have equal keys iff they are equal under simple case-folding.
// [unicode.ToLower] alone misses orbit members with distinct lowercase forms,
// e.g. Greek final sigma 'ς' vs 'σ'.
func foldRune(r rune) rune {
	key := r
	for f := unicode.SimpleFold(r); f != r; f = unicode.SimpleFold(f) {
		key = min(key, f)
	}
	return key
}
