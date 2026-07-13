package set

import (
	"cmp"
	"iter"
	"slices"
)

// SortedSet is a set of ordered items, kept in ascending sorted order at all
// times: [SortedSet.Add] preserves sorted order, and combining sets
// ([SortedSet.Union]/[SortedSet.Intersect]/[SortedSet.Difference]) always
// yields a sorted result. Unlike [Set], [SortedSet.Slice] and [SortedSet.All]
// iterate in deterministic ascending order rather than indeterminate map
// order.
//
// The zero value is an empty, usable [SortedSet].
type SortedSet[T cmp.Ordered] struct {
	items []T
}

// CollectSorted returns a [SortedSet] containing the values of `seq`.
func CollectSorted[T cmp.Ordered](seq iter.Seq[T]) SortedSet[T] {
	return newSorted(slices.Collect(seq))
}

// NewSorted returns a [SortedSet] containing `items`, sorted ascending with
// duplicates removed.
func NewSorted[T cmp.Ordered](items ...T) SortedSet[T] {
	return newSorted(slices.Clone(items))
}

// Add adds `items` to `s`, preserving sorted order and ignoring duplicates.
func (s *SortedSet[T]) Add(items ...T) {
	if len(items) == 0 {
		return
	}
	s.items = unionSorted(s.items, NewSorted(items...).items)
}

// All returns an iterator over the items of `s`, in ascending order.
func (s SortedSet[T]) All() iter.Seq[T] {
	return slices.Values(s.items)
}

// Clone returns a copy of `s`.
func (s SortedSet[T]) Clone() SortedSet[T] {
	return SortedSet[T]{items: slices.Clone(s.items)}
}

// Contains returns whether `item` is present in `s`.
func (s SortedSet[T]) Contains(item T) bool {
	_, found := slices.BinarySearch(s.items, item)
	return found
}

// Delete removes `items` from `s`.
func (s *SortedSet[T]) Delete(items ...T) {
	if len(items) == 0 {
		return
	}
	s.items = differenceSorted(s.items, NewSorted(items...).items)
}

// Difference returns a new [SortedSet] containing the items of `s` not present
// in any of `others`.
func (s SortedSet[T]) Difference(others ...SortedSet[T]) SortedSet[T] {
	result := s.Clone()
	for _, other := range others {
		result.items = differenceSorted(result.items, other.items)
	}
	return result
}

// Equal returns whether `s` and `other` contain the same items.
func (s SortedSet[T]) Equal(other SortedSet[T]) bool {
	return slices.EqualFunc(s.items, other.items, func(a, b T) bool {
		return cmp.Compare(a, b) == 0
	})
}

// Intersect returns a new [SortedSet] containing the items of `s` present in
// every one of `others`.
func (s SortedSet[T]) Intersect(others ...SortedSet[T]) SortedSet[T] {
	result := s.Clone()
	for _, other := range others {
		result.items = intersectSorted(result.items, other.items)
	}
	return result
}

// Len returns the number of items in `s`.
func (s SortedSet[T]) Len() int {
	return len(s.items)
}

// Slice returns the items of `s` as a slice, in ascending order.
func (s SortedSet[T]) Slice() []T {
	return slices.Clone(s.items)
}

// SubsetOf returns whether every item in `s` is present in `other`.
func (s SortedSet[T]) SubsetOf(other SortedSet[T]) bool {
	i, j := 0, 0
	for i < len(s.items) && j < len(other.items) {
		switch cmp.Compare(s.items[i], other.items[j]) {
		case -1:
			return false
		case 0:
			i++
			j++
		case 1:
			j++
		}
	}
	return i == len(s.items)
}

// Union returns a new [SortedSet] containing the items of `s` and all
// `others`.
func (s SortedSet[T]) Union(others ...SortedSet[T]) SortedSet[T] {
	result := s.Clone()
	for _, other := range others {
		result.items = unionSorted(result.items, other.items)
	}
	return result
}

// differenceSorted returns the items in `a` that are absent from `b`.
func differenceSorted[T cmp.Ordered](a, b []T) []T {
	result := make([]T, 0, len(a))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch cmp.Compare(a[i], b[j]) {
		case -1:
			result = append(result, a[i])
			i++
		case 0:
			i++
			j++
		case 1:
			j++
		}
	}
	return append(result, a[i:]...)
}

// intersectSorted returns the items present in both `a` and `b`.
func intersectSorted[T cmp.Ordered](a, b []T) []T {
	result := make([]T, 0, min(len(a), len(b)))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch cmp.Compare(a[i], b[j]) {
		case -1:
			i++
		case 0:
			result = append(result, a[i])
			i++
			j++
		case 1:
			j++
		}
	}
	return result
}

// newSorted takes ownership of `items`, sorts it, and removes duplicates.
func newSorted[T cmp.Ordered](items []T) SortedSet[T] {
	slices.Sort(items)
	items = slices.CompactFunc(items, func(a, b T) bool {
		return cmp.Compare(a, b) == 0
	})
	return SortedSet[T]{items: items}
}

// unionSorted returns the items present in either `a` or `b`.
func unionSorted[T cmp.Ordered](a, b []T) []T {
	result := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch cmp.Compare(a[i], b[j]) {
		case -1:
			result = append(result, a[i])
			i++
		case 0:
			result = append(result, a[i])
			i++
			j++
		case 1:
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	return append(result, b[j:]...)
}
