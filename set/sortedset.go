package set

import (
	"cmp"
	"iter"
	"slices"
)

// SortedSet is a set of ordered items, kept in ascending sorted order at all
// times: [SortedSet.Add] inserts in sorted position, and combining sets
// ([SortedSet.Union]/[SortedSet.Intersect]/[SortedSet.Difference]) always
// yields a sorted result. Unlike [Set], [SortedSet.Slice] and [SortedSet.All]
// iterate in deterministic ascending order rather than indeterminate map
// order.
//
// The zero value is an empty, usable [SortedSet].
type SortedSet[T cmp.Ordered] struct {
	items []T
}

// NewSorted returns a [SortedSet] containing `items`, sorted ascending with
// duplicates removed.
func NewSorted[T cmp.Ordered](items ...T) SortedSet[T] {
	var s SortedSet[T]
	s.Add(items...)
	return s
}

// CollectSorted returns a [SortedSet] containing the values of `seq`.
func CollectSorted[T cmp.Ordered](seq iter.Seq[T]) SortedSet[T] {
	var s SortedSet[T]
	for item := range seq {
		s.Add(item)
	}
	return s
}

// Add adds `items` to `s`, inserting each in sorted position and ignoring
// duplicates.
func (s *SortedSet[T]) Add(items ...T) {
	for _, item := range items {
		i, found := slices.BinarySearch(s.items, item)
		if !found {
			s.items = slices.Insert(s.items, i, item)
		}
	}
}

// Delete removes `items` from `s`.
func (s *SortedSet[T]) Delete(items ...T) {
	for _, item := range items {
		if i, found := slices.BinarySearch(s.items, item); found {
			s.items = slices.Delete(s.items, i, i+1)
		}
	}
}

// Contains returns whether `item` is present in `s`.
func (s SortedSet[T]) Contains(item T) bool {
	_, found := slices.BinarySearch(s.items, item)
	return found
}

// Len returns the number of items in `s`.
func (s SortedSet[T]) Len() int {
	return len(s.items)
}

// Equal returns whether `s` and `other` contain the same items.
func (s SortedSet[T]) Equal(other SortedSet[T]) bool {
	return slices.Equal(s.items, other.items)
}

// SubsetOf returns whether every item in `s` is present in `other`.
func (s SortedSet[T]) SubsetOf(other SortedSet[T]) bool {
	return toSet(s).SubsetOf(toSet(other))
}

// Union returns a new [SortedSet] containing the items of `s` and all
// `others`.
func (s SortedSet[T]) Union(others ...SortedSet[T]) SortedSet[T] {
	return combineSorted(Set[T].Union, s, others)
}

// Intersect returns a new [SortedSet] containing the items of `s` present in
// every one of `others`.
func (s SortedSet[T]) Intersect(others ...SortedSet[T]) SortedSet[T] {
	return combineSorted(Set[T].Intersect, s, others)
}

// Difference returns a new [SortedSet] containing the items of `s` not
// present in any of `others`.
func (s SortedSet[T]) Difference(others ...SortedSet[T]) SortedSet[T] {
	return combineSorted(Set[T].Difference, s, others)
}

// toSet converts a [SortedSet] to an (unordered) [Set].
func toSet[T cmp.Ordered](s SortedSet[T]) Set[T] {
	return New(s.items...)
}

// combineSorted applies a [Set]-combining method
// ([Set.Union]/[Set.Intersect]/[Set.Difference]) to `s` and `others`,
// converting to [Set] and back to [SortedSet] around the call.
func combineSorted[T cmp.Ordered](
	method func(Set[T], ...Set[T]) Set[T],
	s SortedSet[T],
	others []SortedSet[T],
) SortedSet[T] {
	otherSets := make([]Set[T], len(others))
	for i, other := range others {
		otherSets[i] = toSet(other)
	}
	return NewSorted(method(toSet(s), otherSets...).Slice()...)
}

// Clone returns a copy of `s`.
func (s SortedSet[T]) Clone() SortedSet[T] {
	return SortedSet[T]{items: slices.Clone(s.items)}
}

// Slice returns the items of `s` as a slice, in ascending order.
func (s SortedSet[T]) Slice() []T {
	return slices.Clone(s.items)
}

// All returns an iterator over the items of `s`, in ascending order.
func (s SortedSet[T]) All() iter.Seq[T] {
	return slices.Values(s.items)
}
