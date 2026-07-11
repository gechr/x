package maps

import (
	xslices "github.com/gechr/x/slices"
)

// Keys returns the keys of `m` as a slice, in indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// KeysNatural returns the string keys of `m` as a slice, sorted in natural
// order ("item2" before "item10"). See
// [github.com/gechr/x/strings.CompareNatural].
func KeysNatural[M ~map[K]V, K ~string, V any](m M) []K {
	keys := Keys(m)
	xslices.SortNatural(keys)
	return keys
}

// Values returns the values of `m` as a slice, in indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ValuesNatural returns the string values of `m` as a slice, sorted in natural
// order ("item2" before "item10"). See
// [github.com/gechr/x/strings.CompareNatural].
func ValuesNatural[M ~map[K]V, K comparable, V ~string](m M) []V {
	values := Values(m)
	xslices.SortNatural(values)
	return values
}
