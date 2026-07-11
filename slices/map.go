package slices

// Map returns a new slice containing the result of applying `fn` to each
// element of `items`, preserving order.
func Map[S ~[]E, E, R any](items S, fn func(E) R) []R {
	out := make([]R, len(items))
	for i, item := range items {
		out[i] = fn(item)
	}
	return out
}
