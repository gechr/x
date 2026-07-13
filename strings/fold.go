package strings

import (
	"cmp"
	"strings"
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

// ContainsFold reports whether `s` contains `substr`, case-insensitively using
// the same simple case-folding as [strings.EqualFold].
func ContainsFold(s, substr string) bool {
	if substr == "" {
		return true
	}
	for s != "" {
		if HasPrefixFold(s, substr) {
			return true
		}
		_, size := utf8.DecodeRuneInString(s)
		s = s[size:]
	}
	return false
}

// HasPrefixFold reports whether `s` begins with `prefix`, case-insensitively
// using the same simple case-folding as [strings.EqualFold].
func HasPrefixFold(s, prefix string) bool {
	end, ok := prefixRunes(s, utf8.RuneCountInString(prefix))
	return ok && strings.EqualFold(s[:end], prefix)
}

// HasSuffixFold reports whether `s` ends with `suffix`, case-insensitively
// using the same simple case-folding as [strings.EqualFold].
func HasSuffixFold(s, suffix string) bool {
	start, ok := suffixRunes(s, utf8.RuneCountInString(suffix))
	return ok && strings.EqualFold(s[start:], suffix)
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

// prefixRunes returns the byte offset immediately after the first `n` runes of
// `s`, or false when `s` contains fewer than `n` runes.
func prefixRunes(s string, n int) (int, bool) {
	end := 0
	for range n {
		if end == len(s) {
			return 0, false
		}
		_, size := utf8.DecodeRuneInString(s[end:])
		end += size
	}
	return end, true
}

// suffixRunes returns the byte offset at the start of the last `n` runes of
// `s`, or false when `s` contains fewer than `n` runes.
func suffixRunes(s string, n int) (int, bool) {
	start := len(s)
	for range n {
		if start == 0 {
			return 0, false
		}
		_, size := utf8.DecodeLastRuneInString(s[:start])
		start -= size
	}
	return start, true
}
