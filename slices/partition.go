package slices

// Partition splits `items` into two slices: elements satisfying `match`, and
// elements that do not, preserving the original relative order in both.
func Partition[S ~[]E, E any](items S, match func(E) bool) (S, S) {
	matched := make(S, 0, len(items))
	unmatched := make(S, 0, len(items))
	for _, item := range items {
		if match(item) {
			matched = append(matched, item)
		} else {
			unmatched = append(unmatched, item)
		}
	}
	return matched, unmatched
}
