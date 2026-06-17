package strings

import "unicode/utf8"

// TruncateRight shortens s to at most n runes (including marker) by removing
// characters from the right, appending marker when truncation occurs. The head
// is kept. For display-width-aware truncation of ANSI text use ansi.Truncate.
//
//	TruncateRight("hello world", 8, "…") // "hello w…"
//	TruncateRight("hi", 8, "…")          // "hi"
func TruncateRight(s string, n int, marker string) string {
	return truncate(s, n, marker, func(runes []rune, keep int) string {
		return string(runes[:keep]) + marker
	})
}

// TruncateLeft shortens s to at most n runes (including marker) by removing
// characters from the left, prepending marker when truncation occurs. The tail
// is kept.
//
//	TruncateLeft("hello world", 8, "…") // "…o world"
func TruncateLeft(s string, n int, marker string) string {
	return truncate(s, n, marker, func(runes []rune, keep int) string {
		return marker + string(runes[len(runes)-keep:])
	})
}

// TruncateMiddle shortens s to at most n runes (including marker) by removing
// characters from the middle, inserting marker between the kept head and tail so
// both ends stay visible. This suits hashes and paths, where the start and end
// are the recognisable parts.
//
//	TruncateMiddle("0123456789abcdef", 7, "…") // "012…def"
func TruncateMiddle(s string, n int, marker string) string {
	return truncate(s, n, marker, func(runes []rune, keep int) string {
		tail := keep / 2 //nolint:mnd // bisect the kept runes between head and tail
		head := keep - tail
		return string(runes[:head]) + marker + string(runes[len(runes)-tail:])
	})
}

// Truncate is an alias for [TruncateRight], the most common form: it keeps the
// head and trims the tail.
func Truncate(s string, n int, marker string) string {
	return TruncateRight(s, n, marker)
}

// truncate holds the shared truncation preamble: an empty result for a
// non-positive n, s unchanged when it already fits, and a clamped marker when it
// alone would meet or exceed n. Otherwise it hands place the runes and the
// budget of runes to keep (n minus the marker), letting each variant decide
// which runes survive.
func truncate(s string, n int, marker string, place func(runes []rune, keep int) string) string {
	if n <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	markerRunes := utf8.RuneCountInString(marker)
	if markerRunes >= n {
		return string([]rune(marker)[:n])
	}
	return place(runes, n-markerRunes)
}
