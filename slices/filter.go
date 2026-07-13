package slices

// Filter returns the elements of `items` satisfying `keep`, preserving their
// original order.
func Filter[S ~[]E, E any](items S, keep func(E) bool) S {
	filtered := make(S, 0, len(items))
	for _, item := range items {
		if keep(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
