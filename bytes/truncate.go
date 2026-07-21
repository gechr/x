package bytes

import (
	"bytes"
	"unicode/utf8"
)

// TruncateRight shortens `s` to at most `n` runes (including `marker`) by
// removing characters from the right, appending `marker` when truncation
// occurs. The head is kept. For display-width-aware truncation of ANSI text
// use [github.com/gechr/x/ansi.Truncate].
//
//	TruncateRight([]byte("hello world"), 8, []byte("…")) // "hello w…"
//	TruncateRight([]byte("hi"), 8, []byte("…"))          // "hi"
func TruncateRight(s []byte, n int, marker []byte) []byte {
	return truncate(s, n, marker, func(runes []rune, keep int) []byte {
		return append([]byte(string(runes[:keep])), marker...)
	})
}

// TruncateLeft shortens `s` to at most `n` runes (including `marker`) by
// removing characters from the left, prepending `marker` when truncation
// occurs. The tail is kept.
//
//	TruncateLeft([]byte("hello world"), 8, []byte("…")) // "…o world"
func TruncateLeft(s []byte, n int, marker []byte) []byte {
	return truncate(s, n, marker, func(runes []rune, keep int) []byte {
		out := append([]byte(nil), marker...)
		return append(out, string(runes[len(runes)-keep:])...)
	})
}

// TruncateMiddle shortens `s` to at most `n` runes (including `marker`) by
// removing characters from the middle, inserting `marker` between the kept head
// and tail so both ends stay visible. This suits hashes and paths, where the
// start and end are the recognisable parts.
//
//	TruncateMiddle([]byte("0123456789abcdef"), 7, []byte("…")) // "012…def"
func TruncateMiddle(s []byte, n int, marker []byte) []byte {
	return truncate(s, n, marker, func(runes []rune, keep int) []byte {
		tail := keep / 2 //nolint:mnd // bisect the kept runes between head and tail
		head := keep - tail
		out := append([]byte(string(runes[:head])), marker...)
		return append(out, string(runes[len(runes)-tail:])...)
	})
}

// Truncate is an alias for [TruncateRight], the most common form: it keeps the
// head and trims the tail.
func Truncate(s []byte, n int, marker []byte) []byte {
	return TruncateRight(s, n, marker)
}

// truncate holds the shared truncation preamble: an empty result for a
// non-positive `n`, `s` unchanged when it already fits, and a clamped `marker`
// when it alone would meet or exceed `n`. Otherwise it hands `place` the runes
// and the budget of runes to keep (`n` minus the `marker`), letting each
// variant decide which runes survive.
func truncate(s []byte, n int, marker []byte, place func(runes []rune, keep int) []byte) []byte {
	if n <= 0 {
		return nil
	}
	runes := bytes.Runes(s)
	if len(runes) <= n {
		return s
	}
	markerRunes := utf8.RuneCount(marker)
	if markerRunes >= n {
		return []byte(string(bytes.Runes(marker)[:n]))
	}
	return place(runes, n-markerRunes)
}
