package maps

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Sorted returns an iterator over the entries of m in ascending key order.
func Sorted[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range slices.Sorted(maps.Keys(m)) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

// SortedFunc returns an iterator over the entries of m in the key order
// determined by compare, which follows the [cmp.Compare] convention.
func SortedFunc[M ~map[K]V, K comparable, V any](m M, compare func(x, y K) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range slices.SortedFunc(maps.Keys(m), compare) {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
