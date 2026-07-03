package slices

// LastIndex returns the index of the last occurrence of target in items, or
// -1 if not present.
func LastIndex[S ~[]E, E comparable](items S, target E) int {
	for i := len(items) - 1; i >= 0; i-- {
		if items[i] == target {
			return i
		}
	}
	return -1
}

// LastIndexFunc returns the index of the last element of items satisfying
// match, or -1 if none do.
func LastIndexFunc[S ~[]E, E any](items S, match func(E) bool) int {
	for i := len(items) - 1; i >= 0; i-- {
		if match(items[i]) {
			return i
		}
	}
	return -1
}
