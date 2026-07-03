package set

import (
	"cmp"
	"iter"
)

// OrderedSet is a Set whose elements support ordering, enabling [OrderedSet.Sort].
// It has the same underlying representation and semantics as [Set] - use
// [NewOrdered] or [CollectOrdered] to create a usable OrderedSet.
type OrderedSet[T cmp.Ordered] map[T]struct{}

// NewOrdered returns an OrderedSet containing items.
func NewOrdered[T cmp.Ordered](items ...T) OrderedSet[T] {
	return OrderedSet[T](New(items...))
}

// CollectOrdered returns an OrderedSet containing the values of seq.
func CollectOrdered[T cmp.Ordered](seq iter.Seq[T]) OrderedSet[T] {
	return OrderedSet[T](Collect(seq))
}

// Add adds items to s.
func (s OrderedSet[T]) Add(items ...T) {
	Set[T](s).Add(items...)
}

// Delete removes items from s.
func (s OrderedSet[T]) Delete(items ...T) {
	Set[T](s).Delete(items...)
}

// Contains returns whether item is present in s.
func (s OrderedSet[T]) Contains(item T) bool {
	return Set[T](s).Contains(item)
}

// Len returns the number of items in s.
func (s OrderedSet[T]) Len() int {
	return Set[T](s).Len()
}

// Equal returns whether s and other contain the same items.
func (s OrderedSet[T]) Equal(other OrderedSet[T]) bool {
	return Set[T](s).Equal(Set[T](other))
}

// SubsetOf returns whether every item in s is present in other.
func (s OrderedSet[T]) SubsetOf(other OrderedSet[T]) bool {
	return Set[T](s).SubsetOf(Set[T](other))
}

// Union returns a new OrderedSet containing the items of s and all others.
func (s OrderedSet[T]) Union(others ...OrderedSet[T]) OrderedSet[T] {
	return combine(Set[T].Union, s, others)
}

// Intersect returns a new OrderedSet containing the items of s present in
// every one of others.
func (s OrderedSet[T]) Intersect(others ...OrderedSet[T]) OrderedSet[T] {
	return combine(Set[T].Intersect, s, others)
}

// Difference returns a new OrderedSet containing the items of s not present
// in any of others.
func (s OrderedSet[T]) Difference(others ...OrderedSet[T]) OrderedSet[T] {
	return combine(Set[T].Difference, s, others)
}

// combine applies a Set-combining method (Union/Intersect/Difference) to s
// and others, converting to and from OrderedSet around the call.
func combine[T cmp.Ordered](
	method func(Set[T], ...Set[T]) Set[T],
	s OrderedSet[T],
	others []OrderedSet[T],
) OrderedSet[T] {
	converted := make([]Set[T], len(others))
	for i, other := range others {
		converted[i] = Set[T](other)
	}
	return OrderedSet[T](method(Set[T](s), converted...))
}

// Clone returns a copy of s.
func (s OrderedSet[T]) Clone() OrderedSet[T] {
	return OrderedSet[T](Set[T](s).Clone())
}

// Slice returns the items of s as a slice, in indeterminate order.
// Use [OrderedSet.Sort] for a deterministic, ascending-order slice.
func (s OrderedSet[T]) Slice() []T {
	return Set[T](s).Slice()
}

// All returns an iterator over the items of s, in indeterminate order.
func (s OrderedSet[T]) All() iter.Seq[T] {
	return Set[T](s).All()
}

// Sort returns the items of s as a slice in ascending order.
func (s OrderedSet[T]) Sort() []T {
	return Sorted(Set[T](s))
}
