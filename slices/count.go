package slices

// Count returns the number of elements in items equal to target.
func Count[S ~[]E, E comparable](items S, target E) int {
	var n int
	for _, item := range items {
		if item == target {
			n++
		}
	}
	return n
}

// CountFunc returns the number of elements in items satisfying match.
func CountFunc[S ~[]E, E any](items S, match func(E) bool) int {
	var n int
	for _, item := range items {
		if match(item) {
			n++
		}
	}
	return n
}
