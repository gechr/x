package set

import (
	"cmp"
	"slices"

	xstrings "github.com/gechr/x/strings"
)

// Sorted returns the items of `s` as a slice in ascending order.
//
// Sorted is a function rather than a [Set] method because it requires T to
// be ordered, not just comparable.
func Sorted[T cmp.Ordered](s Set[T]) []T {
	items := s.Slice()
	slices.Sort(items)
	return items
}

// SortedNatural returns the items of `s` as a slice in natural order, so
// embedded numbers compare by value ("item2" before "item10") rather than
// lexically. See [xstrings.CompareNatural].
//
// SortedNatural is a function rather than a [Set] method because it
// requires T to be string-like, not just comparable.
func SortedNatural[T ~string](s Set[T]) []T {
	items := s.Slice()
	slices.SortFunc(items, func(a, b T) int {
		return xstrings.CompareNatural(string(a), string(b))
	})
	return items
}
