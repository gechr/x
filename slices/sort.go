package slices

import (
	"slices"

	xstrings "github.com/gechr/x/strings"
)

// SortNatural sorts a string slice in place in natural order, so embedded
// numbers compare by value (`item2` before `item10`) rather than lexically. See
// [github.com/gechr/x/strings.CompareNatural].
func SortNatural[S ~[]E, E ~string](s S) {
	slices.SortFunc(s, func(a, b E) int {
		return xstrings.CompareNatural(string(a), string(b))
	})
}
