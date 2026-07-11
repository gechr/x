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

func TestKeysNatural(t *testing.T) {
	t.Parallel()

	m := map[string]int{"item10": 10, "item2": 2, "item1": 1}
	require.Equal(t, []string{"item1", "item2", "item10"}, xmaps.KeysNatural(m))

	// Empty and nil maps yield empty (non-nil) slices.
	require.Empty(t, xmaps.KeysNatural(map[string]int{}))
	require.Empty(t, xmaps.KeysNatural(map[string]int(nil)))

	// Named string key types satisfy the constraint.
	type name string
	require.Equal(t, []name{"a2", "a10"}, xmaps.KeysNatural(map[name]int{"a10": 10, "a2": 2}))
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

func TestValuesNatural(t *testing.T) {
	t.Parallel()

	m := map[int]string{10: "item10", 2: "item2", 1: "item1"}
	require.Equal(t, []string{"item1", "item2", "item10"}, xmaps.ValuesNatural(m))

	require.Empty(t, xmaps.ValuesNatural(map[int]string{}))
	require.Empty(t, xmaps.ValuesNatural(map[int]string(nil)))

	// Duplicate values are preserved.
	require.Equal(t, []string{"a", "a"}, xmaps.ValuesNatural(map[int]string{1: "a", 2: "a"}))
}
