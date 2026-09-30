package maps_test

import (
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	a := map[string]int{"a": 1, "b": 2}
	b := map[string]int{"b": 20, "c": 30}
	require.Equal(t, map[string]int{"a": 1, "b": 20, "c": 30}, xmaps.Merge(a, b))

	// Inputs are left unmodified.
	require.Equal(t, map[string]int{"a": 1, "b": 2}, a)
	require.Equal(t, map[string]int{"b": 20, "c": 30}, b)

	// The last map wins across more than two inputs.
	require.Equal(t, map[string]int{"k": 3},
		xmaps.Merge(map[string]int{"k": 1}, map[string]int{"k": 2}, map[string]int{"k": 3}))

	// No inputs, nil inputs and empty inputs yield empty (non-nil) maps.
	require.NotNil(t, xmaps.Merge[map[string]int]())
	require.Empty(t, xmaps.Merge[map[string]int]())
	require.Empty(t, xmaps.Merge(map[string]int(nil), map[string]int{}))

	// A single input is copied, not aliased.
	src := map[string]int{"a": 1}
	merged := xmaps.Merge(src)
	merged["a"] = 2
	require.Equal(t, 1, src["a"])

	// Named map types are preserved.
	type env map[string]string
	require.Equal(t, env{"HOME": "/root", "SHELL": "fish"},
		xmaps.Merge(env{"HOME": "/root"}, env{"SHELL": "fish"}))

	// Zero values are real values and overwrite earlier ones.
	require.Equal(t, map[string]int{"k": 0},
		xmaps.Merge(map[string]int{"k": 1}, map[string]int{"k": 0}))
}

func TestMergeFunc(t *testing.T) {
	t.Parallel()

	sum := func(_ string, existing, incoming int) int { return existing + incoming }
	require.Equal(t, map[string]int{"a": 1, "b": 22, "c": 300},
		xmaps.MergeFunc(sum,
			map[string]int{"a": 1, "b": 2},
			map[string]int{"b": 20, "c": 300}))

	// Keeping the existing value makes the first map win.
	first := func(_ string, existing, _ int) int { return existing }
	require.Equal(t, map[string]int{"k": 1},
		xmaps.MergeFunc(first, map[string]int{"k": 1}, map[string]int{"k": 2}))

	// resolve runs only for duplicate keys and receives the key.
	var calls []string
	record := func(k string, _, incoming int) int {
		calls = append(calls, k)
		return incoming
	}
	xmaps.MergeFunc(record, map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
	require.Equal(t, []string{"b"}, calls)

	// No inputs yield an empty (non-nil) map.
	require.NotNil(t, xmaps.MergeFunc[map[string]int](sum))
}

func TestMergeDeep(t *testing.T) {
	t.Parallel()

	// The last map wins, nested maps merge and slices are replaced.
	dst := map[string]any{
		"a": "one",
		"c": 3,
		"d": map[string]any{"f": 5},
		"g": []int{8, 9},
		"i": "eye",
		"k": map[string]any{"l": true},
	}
	src1 := map[string]any{
		"a": 1,
		"b": 2,
		"d": map[string]any{"e": "four"},
		"g": []int{6, 7},
		"i": "aye",
		"j": "jay",
		"k": map[string]any{"l": false},
	}
	src2 := map[string]any{"h": 10, "i": "i", "j": "j"}
	require.Equal(t, map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": map[string]any{"e": "four", "f": 5},
		"g": []int{6, 7}, // slices are replaced
		"h": 10,
		"i": "i",
		"j": "j",
		"k": map[string]any{"l": false},
	}, xmaps.MergeDeep(dst, src1, src2))

	// Inputs, including nested maps, are left unmodified.
	require.Equal(t, map[string]any{"f": 5}, dst["d"])
	require.Equal(t, map[string]any{"e": "four"}, src1["d"])
	require.Equal(t, map[string]any{"l": true}, dst["k"])

	// Nested maps merge at every depth.
	require.Equal(t,
		map[string]any{"x": map[string]any{"y": map[string]any{"a": 1, "b": 2}}},
		xmaps.MergeDeep(
			map[string]any{"x": map[string]any{"y": map[string]any{"a": 1}}},
			map[string]any{"x": map[string]any{"y": map[string]any{"b": 2}}},
		))

	// A nested map taken from one input is not mutated by a later merge.
	shared := map[string]any{"a": 1}
	merged := xmaps.MergeDeep(
		map[string]any{"n": shared},
		map[string]any{"n": map[string]any{"b": 2}},
		map[string]any{"n": map[string]any{"c": 3}},
	)
	require.Equal(t, map[string]any{"a": 1, "b": 2, "c": 3}, merged["n"])
	require.Equal(t, map[string]any{"a": 1}, shared)

	// A map and a non-map replace each other, in either order.
	require.Equal(t, map[string]any{"k": "scalar"},
		xmaps.MergeDeep(map[string]any{"k": map[string]any{"a": 1}}, map[string]any{"k": "scalar"}))
	require.Equal(t, map[string]any{"k": map[string]any{"a": 1}},
		xmaps.MergeDeep(map[string]any{"k": "scalar"}, map[string]any{"k": map[string]any{"a": 1}}))

	// Nil and zero values are real values and overwrite earlier ones.
	require.Equal(t, map[string]any{"k": nil, "b": false},
		xmaps.MergeDeep(
			map[string]any{"k": map[string]any{"a": 1}, "b": true},
			map[string]any{"k": nil, "b": false},
		))

	// Nested maps of other types are replaced, not merged.
	require.Equal(t, map[string]any{"k": map[string]int{"b": 2}},
		xmaps.MergeDeep(
			map[string]any{"k": map[string]int{"a": 1}},
			map[string]any{"k": map[string]int{"b": 2}},
		))

	// No inputs, nil inputs and empty inputs yield empty (non-nil) maps.
	require.NotNil(t, xmaps.MergeDeep[map[string]any]())
	require.Empty(t, xmaps.MergeDeep(map[string]any(nil), map[string]any{}))

	// Named map types are preserved at the top level.
	type values map[string]any
	require.Equal(t, values{"a": 1, "b": 2}, xmaps.MergeDeep(values{"a": 1}, values{"b": 2}))
}
