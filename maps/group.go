package maps

import "iter"

// Group collects the pairs of seq into a map of slices, grouping values by
// key in encounter order.
func Group[K comparable, V any](seq iter.Seq2[K, V]) map[K][]V {
	groups := make(map[K][]V)
	for k, v := range seq {
		groups[k] = append(groups[k], v)
	}
	return groups
}

// GroupFunc collects the values of seq into a map of slices, grouping values
// in encounter order by the key returned by key.
func GroupFunc[K comparable, V any](seq iter.Seq[V], key func(V) K) map[K][]V {
	groups := make(map[K][]V)
	for v := range seq {
		k := key(v)
		groups[k] = append(groups[k], v)
	}
	return groups
}
