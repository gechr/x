package encoding_test

import (
	"math"
	"testing"

	xencoding "github.com/gechr/x/encoding"
	"github.com/stretchr/testify/require"
)

func TestLookup(t *testing.T) {
	t.Parallel()

	doc := map[string]any{
		"items": []any{
			map[string]any{"foo": map[string]any{"bar": []any{1, 2}}},
		},
	}

	v, ok := xencoding.NewPath("items").Index(0).Child("foo", "bar").Index(1).Lookup(doc)
	require.True(t, ok)
	require.Equal(t, 2, v)

	// Key behaves exactly like Child.
	v, ok = xencoding.NewPath("items").Index(0).Key("foo").Lookup(doc)
	require.True(t, ok)
	require.Equal(t, map[string]any{"bar": []any{1, 2}}, v)

	// A single-segment path resolves a top-level key.
	v, ok = xencoding.NewPath("items").Lookup(doc)
	require.True(t, ok)
	require.Equal(t, doc["items"], v)

	// map[any]any documents (legacy YAML decoders) resolve too.
	legacy := map[any]any{"a": map[any]any{"b": 1}}
	v, ok = xencoding.NewPath("a", "b").Lookup(legacy)
	require.True(t, ok)
	require.Equal(t, 1, v)

	// Non-string map[any]any keys match by their string form.
	v, ok = xencoding.NewPath("2").Lookup(map[any]any{2: "two"})
	require.True(t, ok)
	require.Equal(t, "two", v)

	// An exact string key wins over a stringified match.
	v, ok = xencoding.NewPath("2").Lookup(map[any]any{2: "int", "2": "str"})
	require.True(t, ok)
	require.Equal(t, "str", v)

	// Missing key, index out of range, negative index, non-container.
	_, ok = xencoding.NewPath("missing").Lookup(doc)
	require.False(t, ok)
	_, ok = xencoding.NewPath("items").Index(1).Lookup(doc)
	require.False(t, ok)
	_, ok = xencoding.NewPath("items").Index(-1).Lookup(doc)
	require.False(t, ok)
	_, ok = xencoding.NewPath("items").Index(0).Child("foo", "bar").Index(0).Child("x").Lookup(doc)
	require.False(t, ok)

	// Wildcard paths cannot resolve to a single value.
	_, ok = xencoding.NewPath("items").Wildcard().Lookup(doc)
	require.False(t, ok)
}

func TestLookupAll(t *testing.T) {
	t.Parallel()

	doc := map[string]any{
		"items": []any{
			map[string]any{"name": "a"},
			map[string]any{"name": "b"},
		},
	}

	// Wildcards fan out over array elements in order.
	require.Equal(
		t,
		[]any{"a", "b"},
		xencoding.NewPath("items").Wildcard().Child("name").LookupAll(doc),
	)

	// A concrete path yields a single value.
	require.Equal(t, []any{"a"}, xencoding.NewPath("items").Index(0).Child("name").LookupAll(doc))

	// Map wildcards fan out in natural key order.
	cfg := map[string]any{"cfg": map[string]any{"item10": 10, "item2": 2}}
	require.Equal(t, []any{2, 10}, xencoding.NewPath("cfg").Wildcard().LookupAll(cfg))

	// map[any]any keys sort by their string form.
	legacy := map[string]any{"m": map[any]any{10: "ten", 2: "two", "a": "letter"}}
	require.Equal(
		t,
		[]any{"two", "ten", "letter"},
		xencoding.NewPath("m").Wildcard().LookupAll(legacy),
	)

	// Elements missing a later segment are skipped.
	mixed := map[string]any{"items": []any{
		map[string]any{"name": "a"},
		map[string]any{"other": 1},
	}}
	require.Equal(
		t,
		[]any{"a"},
		xencoding.NewPath("items").Wildcard().Child("name").LookupAll(mixed),
	)

	// NaN keys still yield their values (NaN is not equal to itself, so a
	// post-sort re-fetch by key would silently drop them).
	nan := map[string]any{"m": map[any]any{math.NaN(): "nan", "a": "x"}}
	require.ElementsMatch(t, []any{"nan", "x"}, xencoding.NewPath("m").Wildcard().LookupAll(nan))

	// Naturally-equal distinct keys ("a0" vs "a00") order lexically.
	ties := map[string]any{"m": map[string]any{"a00": 2, "a0": 1}}
	require.Equal(t, []any{1, 2}, xencoding.NewPath("m").Wildcard().LookupAll(ties))

	// Keys with identical string forms (2 vs "2") order by key type.
	typed := map[string]any{"m": map[any]any{2: "int", "2": "string"}}
	require.Equal(t, []any{"int", "string"}, xencoding.NewPath("m").Wildcard().LookupAll(typed))

	// Nothing matches: nil.
	require.Nil(t, xencoding.NewPath("missing").Wildcard().LookupAll(doc))

	// Wildcard on a scalar: nil.
	require.Nil(t, xencoding.NewPath("items").Index(0).Child("name").Wildcard().LookupAll(doc))
}
