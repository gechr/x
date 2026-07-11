package maps_test

import (
	"slices"
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	keys := xmaps.Keys(m)
	slices.Sort(keys)
	require.Equal(t, []string{"alpha", "beta", "charlie"}, keys)

	// Empty and nil maps yield empty (non-nil) slices.
	require.Empty(t, xmaps.Keys(map[string]int{}))
	require.Empty(t, xmaps.Keys(map[string]int(nil)))

	// Named map types satisfy the constraint.
	type registry map[string]int
	require.ElementsMatch(t, []string{"a", "b"}, xmaps.Keys(registry{"a": 1, "b": 2}))
}

func TestValues(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	values := xmaps.Values(m)
	slices.Sort(values)
	require.Equal(t, []int{1, 2, 3}, values)

	require.Empty(t, xmaps.Values(map[string]int{}))
	require.Empty(t, xmaps.Values(map[string]int(nil)))

	// Duplicate values are preserved.
	require.ElementsMatch(t, []int{1, 1}, xmaps.Values(map[string]int{"a": 1, "b": 1}))
}
