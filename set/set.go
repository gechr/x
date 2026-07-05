// Package set provides a generic set backed by a map.
package set

import (
	"iter"
	"slices"
)

// Set is a set of comparable items backed by a map. Pointer types and
// structs containing pointer fields are compared using shallow equality.
// The zero value is nil: read operations work, but [Set.Add] panics - use [New]
// or [Collect] to create a usable Set.
type Set[T comparable] map[T]struct{}

// New returns a Set containing `items`.
func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	s.Add(items...)
	return s
}

// Collect returns a Set containing the values of `seq`.
func Collect[T comparable](seq iter.Seq[T]) Set[T] {
	s := make(Set[T])
	for item := range seq {
		s[item] = struct{}{}
	}
	return s
}

// Add adds `items` to `s`.
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

// Delete removes `items` from `s`.
func (s Set[T]) Delete(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

// Contains returns whether `item` is present in `s`.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Len returns the number of items in `s`.
func (s Set[T]) Len() int {
	return len(s)
}

// Equal returns whether `s` and `other` contain the same items.
func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// SubsetOf returns whether every item in `s` is present in `other`.
func (s Set[T]) SubsetOf(other Set[T]) bool {
	if len(s) > len(other) {
		return false
	}
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Union returns a new Set containing the items of `s` and all `others`.
func (s Set[T]) Union(others ...Set[T]) Set[T] {
	union := make(Set[T], len(s))
	for item := range s {
		union[item] = struct{}{}
	}
	for _, other := range others {
		for item := range other {
			union[item] = struct{}{}
		}
	}
	return union
}

// Intersect returns a new Set containing the items of `s` present in every one
// of `others`.
func (s Set[T]) Intersect(others ...Set[T]) Set[T] {
	intersection := make(Set[T])
	for item := range s {
		if !slices.ContainsFunc(others, func(other Set[T]) bool { return !other.Contains(item) }) {
			intersection[item] = struct{}{}
		}
	}
	return intersection
}

// Difference returns a new Set containing the items of `s` not present in any
// of `others`.
func (s Set[T]) Difference(others ...Set[T]) Set[T] {
	diff := make(Set[T])
	for item := range s {
		if !slices.ContainsFunc(others, func(other Set[T]) bool { return other.Contains(item) }) {
			diff[item] = struct{}{}
		}
	}
	return diff
}

// Clone returns a copy of `s`.
func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T], len(s))
	for item := range s {
		clone[item] = struct{}{}
	}
	return clone
}

// Slice returns the items of `s` as a slice, in indeterminate order.
func (s Set[T]) Slice() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

// All returns an iterator over the items of `s`, in indeterminate order.
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s {
			if !yield(item) {
				return
			}
		}
	}
}
