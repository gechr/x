package set

import (
	"cmp"
	"slices"
)

// Sorted returns the items of s as a slice in ascending order.
//
// Sorted is a function rather than a [Set] method because it requires T to
// be ordered, not just comparable.
func Sorted[T cmp.Ordered](s Set[T]) []T {
	items := s.Slice()
	slices.Sort(items)
	return items
}
