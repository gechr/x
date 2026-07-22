package slices

// Surround returns a new slice with `prefix` and `suffix` concatenated onto
// each element of `items`.
func Surround[S ~[]E, E ~string](items S, prefix, suffix E) []E {
	out := make([]E, len(items))
	for i, item := range items {
		out[i] = prefix + item + suffix
	}
	return out
}
