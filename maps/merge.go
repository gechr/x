package maps

import "maps"

// Merge returns a new map holding the entries of every map in `ms`. If a key
// appears in more than one map, the value from the last map wins. The inputs
// are left unmodified.
func Merge[M ~map[K]V, K comparable, V any](ms ...M) M {
	merged := make(M, mergeSize(ms))
	for _, m := range ms {
		maps.Copy(merged, m)
	}
	return merged
}

// MergeFunc returns a new map holding the entries of every map in `ms`. If a
// key appears in more than one map, `resolve` receives the key, the value
// merged so far and the incoming value, and returns the value to keep. The
// inputs are left unmodified.
func MergeFunc[M ~map[K]V, K comparable, V any](
	resolve func(k K, existing, incoming V) V,
	ms ...M,
) M {
	merged := make(M, mergeSize(ms))
	for _, m := range ms {
		for k, v := range m {
			if existing, ok := merged[k]; ok {
				v = resolve(k, existing, v)
			}
			merged[k] = v
		}
	}
	return merged
}

// MergeDeep returns a new map holding the entries of every map in `ms`,
// merging nested `map[K]any` values recursively. If a key appears in more than
// one map and either value is not a `map[K]any`, the value from the last map
// wins. Slices are replaced, not concatenated.
//
// The inputs are left unmodified: every merged level is a new map. Nested maps
// that appear in only one input are shared with that input, not copied.
func MergeDeep[M ~map[K]any, K comparable](ms ...M) M {
	merged := make(M, mergeSize(ms))
	for _, m := range ms {
		for k, v := range m {
			if dst, ok := merged[k].(map[K]any); ok {
				if src, ok := v.(map[K]any); ok {
					v = MergeDeep(dst, src)
				}
			}
			merged[k] = v
		}
	}
	return merged
}

// mergeSize returns the length of the largest map in `ms`, a lower bound on
// the size of the merged map.
func mergeSize[M ~map[K]V, K comparable, V any](ms []M) int {
	size := 0
	for _, m := range ms {
		size = max(size, len(m))
	}
	return size
}
