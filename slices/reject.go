package slices

// Reject returns the elements of `items` not satisfying `drop`, preserving
// their original order.
func Reject[S ~[]E, E any](items S, drop func(E) bool) S {
	return Filter(items, func(item E) bool { return !drop(item) })
}
