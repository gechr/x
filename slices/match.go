package slices

import "slices"

// AllFunc reports whether every element of `items` satisfies `match`.
// It returns true when `items` is empty.
func AllFunc[S ~[]E, E any](items S, match func(E) bool) bool {
	return !slices.ContainsFunc(items, func(item E) bool {
		return !match(item)
	})
}

// AnyFunc reports whether any element of `items` satisfies `match`.
// It returns false when `items` is empty.
func AnyFunc[S ~[]E, E any](items S, match func(E) bool) bool {
	return slices.ContainsFunc(items, match)
}
