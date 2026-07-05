package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		[]string{"a", "b", "A", "c"},
		xslices.Unique([]string{"a", "b", "a", "A", "c", "b"}),
	)
	require.Equal(t, []int{1, 2, 3}, xslices.Unique([]int{1, 2, 1, 3, 2}))
	require.Equal(t, []string{}, xslices.Unique([]string{}))
}

func TestUniqueFunc(t *testing.T) {
	t.Parallel()

	type pair struct {
		key string
		val int
	}
	// First-seen wins for each key.
	require.Equal(
		t,
		[]pair{{"a", 1}, {"b", 2}, {"c", 4}},
		xslices.UniqueFunc(
			[]pair{{"a", 1}, {"b", 2}, {"a", 3}, {"c", 4}, {"b", 5}},
			func(p pair) string { return p.key },
		),
	)
	require.Equal(
		t,
		[]int{1, 2, 10},
		xslices.UniqueFunc([]int{1, 2, 11, 10, 22}, func(n int) int { return n % 10 }),
	)
	require.Equal(
		t,
		[]string{},
		xslices.UniqueFunc([]string{}, func(s string) string { return s }),
	)
}

func TestUniqueFold(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "B", "c"}, xslices.UniqueFold([]string{"a", "B", "A", "b", "c"}))
	require.Equal(
		t,
		[]namedString{"one", "two"},
		xslices.UniqueFold([]namedString{"one", "ONE", "two"}),
	)
	// EqualFold-equal pairs with distinct lowercase forms: Greek sigma
	// (Σ/σ/ς) and Kelvin sign (K/K/k).
	require.Equal(t, []string{"Σ"}, xslices.UniqueFold([]string{"Σ", "ς", "σ"}))
	require.Equal(t, []string{"K"}, xslices.UniqueFold([]string{"K", "K", "k"}))
}

type namedString string
