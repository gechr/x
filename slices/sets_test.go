package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestDifference(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "c"},
		xslices.Difference([]string{"a", "b", "c"}, []string{"b"}))

	// Duplicates in the first slice are preserved.
	require.Equal(t, []int{1, 1, 3},
		xslices.Difference([]int{1, 1, 2, 3}, []int{2}))

	// Empty and nil inputs.
	require.Empty(t, xslices.Difference([]int(nil), []int{1}))
	require.Equal(t, []int{1}, xslices.Difference([]int{1}, nil))
}

func TestIntersect(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"b", "c"},
		xslices.Intersect([]string{"a", "b", "c"}, []string{"c", "b", "d"}))

	// Order and duplicates follow the first slice.
	require.Equal(t, []int{2, 2},
		xslices.Intersect([]int{2, 1, 2}, []int{2, 3}))

	require.Empty(t, xslices.Intersect([]int{1, 2}, []int{3}))
	require.Empty(t, xslices.Intersect([]int(nil), []int{1}))
}

func TestUnion(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "b", "c"},
		xslices.Union([]string{"a", "b"}, []string{"b", "c"}))

	// Variadic: any number of slices, first-seen order.
	require.Equal(t, []int{1, 2, 3, 4},
		xslices.Union([]int{1, 2}, []int{2, 3}, []int{3, 4, 1}))

	// Single slice just dedupes.
	require.Equal(t, []int{1, 2}, xslices.Union([]int{1, 2, 1}))

	require.Empty(t, xslices.Union([]int(nil), []int(nil)))
}
