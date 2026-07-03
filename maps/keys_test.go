package maps_test

import (
	"slices"
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestKeysSlice(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	keys := xmaps.KeysSlice(m)
	slices.Sort(keys)
	require.Equal(t, []string{"alpha", "beta", "charlie"}, keys)

	// Empty and nil maps yield empty (non-nil) slices.
	require.Empty(t, xmaps.KeysSlice(map[string]int{}))
	require.Empty(t, xmaps.KeysSlice(map[string]int(nil)))

	// Named map types satisfy the constraint.
	type registry map[string]int
	require.ElementsMatch(t, []string{"a", "b"}, xmaps.KeysSlice(registry{"a": 1, "b": 2}))
}

func TestValuesSlice(t *testing.T) {
	t.Parallel()

	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}

	values := xmaps.ValuesSlice(m)
	slices.Sort(values)
	require.Equal(t, []int{1, 2, 3}, values)

	require.Empty(t, xmaps.ValuesSlice(map[string]int{}))
	require.Empty(t, xmaps.ValuesSlice(map[string]int(nil)))

	// Duplicate values are preserved.
	require.ElementsMatch(t, []int{1, 1}, xmaps.ValuesSlice(map[string]int{"a": 1, "b": 1}))
}
