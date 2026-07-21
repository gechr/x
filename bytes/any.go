// Package bytes provides byte-slice helpers mirroring
// [github.com/gechr/x/strings]: split, contains, indent/dedent, truncate, and
// blank checks.
package bytes

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// SplitAny splits `s` around each occurrence of any Unicode code point in
// `chars`, following the cutset convention of [bytes.IndexAny]. Empty
// segments between adjacent separators are preserved, matching [bytes.Split]
// semantics. If `chars` is empty, [SplitAny] returns a single-element slice
// containing `s`.
func SplitAny(s []byte, chars string) [][]byte {
	if chars == "" {
		return [][]byte{s}
	}
	var parts [][]byte
	for {
		i := bytes.IndexAny(s, chars)
		if i < 0 {
			return append(parts, s)
		}
		parts = append(parts, s[:i])
		_, width := utf8.DecodeRune(s[i:])
		s = s[i+width:]
	}
}

// CountAny returns the number of Unicode code points in `s` that are contained
// in `chars`, following the cutset convention of [bytes.IndexAny].
func CountAny(s []byte, chars string) int {
	var n int
	for _, r := range string(s) {
		if strings.ContainsRune(chars, r) {
			n++
		}
	}
	return n
}
