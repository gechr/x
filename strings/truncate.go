package strings

import "unicode/utf8"

// Truncate shortens s to at most n runes (including suffix), appending suffix
// when truncation occurs. For display-width-aware truncation of ANSI text use
// ansi.Truncate.
//
//	Truncate("hello world", 8, "…") // "hello w…"
//	Truncate("hi", 8, "…")          // "hi"
func Truncate(s string, n int, suffix string) string {
	if n <= 0 {
		return ""
	}
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	suffixRunes := utf8.RuneCountInString(suffix)
	if suffixRunes >= n {
		return string([]rune(suffix)[:n])
	}
	keep := n - suffixRunes
	runes := []rune(s)
	return string(runes[:keep]) + suffix
}
