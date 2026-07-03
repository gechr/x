package maps

// Invert returns a new map with the keys and values of m swapped. If multiple
// keys map to the same value, exactly one of them survives as the value in
// the result, chosen arbitrarily due to map iteration order.
func Invert[M ~map[K]V, K, V comparable](m M) map[V]K {
	inverted := make(map[V]K, len(m))
	for k, v := range m {
		inverted[v] = k
	}
	return inverted
}
