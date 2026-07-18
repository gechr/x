package slices

import (
	"slices"

	"github.com/gechr/x/internal/natural"
)

// SortNatural sorts a string slice in place in natural order, so embedded
// numbers compare by value (`item2` before `item10`) rather than lexically. See
// [github.com/gechr/x/strings.CompareNatural].
func SortNatural[S ~[]E, E ~string](s S) {
	slices.SortFunc(s, func(a, b E) int {
		return natural.Compare(string(a), string(b))
	})
}
