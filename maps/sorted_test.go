package maps_test

import (
	"cmp"
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestSorted(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	var keys []string
	var values []int
	for k, v := range xmaps.Sorted(m) {
		keys = append(keys, k)
		values = append(values, v)
	}
	require.Equal(t, []string{"alpha", "beta", "charlie"}, keys)
	require.Equal(t, []int{1, 2, 3}, values)

	// Early break stops iteration.
	var first string
	for k := range xmaps.Sorted(m) {
		first = k
		break
	}
	require.Equal(t, "alpha", first)

	// Empty and nil maps yield nothing.
	for range xmaps.Sorted(map[string]int{}) {
		t.Fatal("unexpected iteration over empty map")
	}
	for range xmaps.Sorted(map[string]int(nil)) {
		t.Fatal("unexpected iteration over nil map")
	}
}

func TestSortedFunc(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	// Descending key order.
	var keys []string
	for k := range xmaps.SortedFunc(m, func(x, y string) int { return cmp.Compare(y, x) }) {
		keys = append(keys, k)
	}
	require.Equal(t, []string{"charlie", "beta", "alpha"}, keys)

	// Named map types satisfy the constraint.
	type registry map[string]int
	var got []string
	for k := range xmaps.SortedFunc(registry{"b": 2, "a": 1}, cmp.Compare) {
		got = append(got, k)
	}
	require.Equal(t, []string{"a", "b"}, got)
}
