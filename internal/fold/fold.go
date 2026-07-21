// Package fold implements the simple Unicode case-folding shared by the public
// [github.com/gechr/x/strings] and [github.com/gechr/x/bytes] packages. It sits
// below them so both fold identically - the same orbit canonicaliser and the
// same rune-scanning traversal - without duplicating that logic and without an
// import cycle.
package fold

import (
	"bytes"
	"cmp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Text is the set of concrete types the fold helpers operate on. It lists
// `string` and `[]byte` exactly, rather than `~string | ~[]byte`, so the
// `any(s).(type)` switches that dispatch to the string- or byte-specific
// `utf8`/`EqualFold` primitives stay exhaustive.
type Text interface {
	string | []byte
}

// Rune maps `r` to the canonical (smallest) member of its case-fold orbit, so
// two runes have equal keys iff they are equal under simple case-folding.
// [unicode.ToLower] alone misses orbit members with distinct lowercase forms,
// for example Greek final sigma 'ς' vs 'σ'.
func Rune(r rune) rune {
	key := r
	for f := unicode.SimpleFold(r); f != r; f = unicode.SimpleFold(f) {
		key = min(key, f)
	}
	return key
}

// Compare compares `a` and `b` case-insensitively under simple case-folding and
// returns -1, 0, or 1 following the [cmp.Compare] convention. `Compare(a, b)`
// is 0 iff `a` and `b` are equal under [strings.EqualFold]/[bytes.EqualFold].
func Compare[T Text](a, b T) int {
	for len(a) > 0 && len(b) > 0 {
		ra, na := decodeRune(a)
		rb, nb := decodeRune(b)
		a, b = a[na:], b[nb:]
		if c := cmp.Compare(Rune(ra), Rune(rb)); c != 0 {
			return c
		}
	}
	return cmp.Compare(len(a), len(b))
}

// Contains reports whether `s` contains `substr`, case-insensitively under
// simple case-folding.
func Contains[T Text](s, substr T) bool {
	if len(substr) == 0 {
		return true
	}
	for len(s) > 0 {
		if HasPrefix(s, substr) {
			return true
		}
		_, size := decodeRune(s)
		s = s[size:]
	}
	return false
}

// HasPrefix reports whether `s` begins with `prefix`, case-insensitively under
// simple case-folding.
func HasPrefix[T Text](s, prefix T) bool {
	end, ok := prefixRunes(s, runeCount(prefix))
	return ok && equalFold(s[:end], prefix)
}

// HasSuffix reports whether `s` ends with `suffix`, case-insensitively under
// simple case-folding.
func HasSuffix[T Text](s, suffix T) bool {
	start, ok := suffixRunes(s, runeCount(suffix))
	return ok && equalFold(s[start:], suffix)
}

// prefixRunes returns the byte offset immediately after the first `n` runes of
// `s`, or false when `s` contains fewer than `n` runes.
func prefixRunes[T Text](s T, n int) (int, bool) {
	end := 0
	for range n {
		if end == len(s) {
			return 0, false
		}
		_, size := decodeRune(s[end:])
		end += size
	}
	return end, true
}

// suffixRunes returns the byte offset at the start of the last `n` runes of
// `s`, or false when `s` contains fewer than `n` runes.
func suffixRunes[T Text](s T, n int) (int, bool) {
	start := len(s)
	for range n {
		if start == 0 {
			return 0, false
		}
		_, size := decodeLastRune(s[:start])
		start -= size
	}
	return start, true
}

// decodeRune reads the first rune of `s` and its byte width, dispatching to the
// string or byte primitive so the traversal above stays type-agnostic.
func decodeRune[T Text](s T) (rune, int) {
	switch v := any(s).(type) {
	case string:
		return utf8.DecodeRuneInString(v)
	case []byte:
		return utf8.DecodeRune(v)
	default:
		return utf8.RuneError, 0
	}
}

// decodeLastRune reads the last rune of `s` and its byte width, dispatching to
// the string or byte primitive.
func decodeLastRune[T Text](s T) (rune, int) {
	switch v := any(s).(type) {
	case string:
		return utf8.DecodeLastRuneInString(v)
	case []byte:
		return utf8.DecodeLastRune(v)
	default:
		return utf8.RuneError, 0
	}
}

// runeCount returns the number of runes in `s`, dispatching to the string or
// byte primitive.
func runeCount[T Text](s T) int {
	switch v := any(s).(type) {
	case string:
		return utf8.RuneCountInString(v)
	case []byte:
		return utf8.RuneCount(v)
	default:
		return 0
	}
}

// equalFold reports whether `a` and `b` are equal under simple case-folding,
// dispatching to [strings.EqualFold] or [bytes.EqualFold]. Both arguments share
// the concrete type `T`, so the assertion on `b` always succeeds.
func equalFold[T Text](a, b T) bool {
	switch v := any(a).(type) {
	case string:
		w, _ := any(b).(string)
		return strings.EqualFold(v, w)
	case []byte:
		w, _ := any(b).([]byte)
		return bytes.EqualFold(v, w)
	default:
		return false
	}
}
