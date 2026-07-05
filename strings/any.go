// Package strings provides string helpers: split, contains, indent/dedent, truncate, and blank checks.
package strings

import (
	"strings"
	"unicode/utf8"
)

// SplitAny splits `s` around each occurrence of any Unicode code point in
// `chars`, following the cutset convention of [strings.IndexAny]. Empty
// segments between adjacent separators are preserved, matching
// [strings.Split] semantics. If `chars` is empty, SplitAny returns a
// single-element slice containing `s`.
func SplitAny(s, chars string) []string {
	if chars == "" {
		return []string{s}
	}
	var parts []string
	for {
		i := strings.IndexAny(s, chars)
		if i < 0 {
			return append(parts, s)
		}
		parts = append(parts, s[:i])
		_, width := utf8.DecodeRuneInString(s[i:])
		s = s[i+width:]
	}
}

// CountAny returns the number of Unicode code points in `s` that are contained
// in `chars`, following the cutset convention of [strings.IndexAny].
func CountAny(s, chars string) int {
	var n int
	for _, r := range s {
		if strings.ContainsRune(chars, r) {
			n++
		}
	}
	return n
}
