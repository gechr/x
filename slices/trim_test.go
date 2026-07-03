package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestTrim(t *testing.T) {
	t.Parallel()

	require.Equal(t, []int{1, 2}, xslices.Trim([]int{0, 0, 1, 2, 0}, []int{0}))
	require.Equal(t, []int{1, 0, 2}, xslices.Trim([]int{1, 0, 2}, []int{0}))
	require.Empty(t, xslices.Trim([]int{0, 0}, []int{0}))
	require.Empty(t, xslices.Trim([]int(nil), []int{0}))

	// Multi-element cutset.
	require.Equal(t, []string{"c"},
		xslices.Trim([]string{"a", "b", "c", "b", "a"}, []string{"a", "b"}))

	// Empty cutset trims nothing.
	require.Equal(t, []int{0, 1}, xslices.Trim([]int{0, 1}, nil))

	// Result shares the backing array (zero-copy).
	items := []int{0, 1, 2, 0}
	trimmed := xslices.Trim(items, []int{0})
	require.Equal(t, &items[1], &trimmed[0])
}

func TestTrimLeft(t *testing.T) {
	t.Parallel()

	require.Equal(t, []int{1, 2, 0}, xslices.TrimLeft([]int{0, 0, 1, 2, 0}, []int{0}))
	require.Empty(t, xslices.TrimLeft([]int{0}, []int{0}))
	require.Empty(t, xslices.TrimLeft([]int(nil), []int{0}))

	// Named slice types are preserved.
	type tags []string
	require.Equal(t, tags{"x"}, xslices.TrimLeft(tags{"", "x"}, tags{""}))
}

func TestTrimRight(t *testing.T) {
	t.Parallel()

	require.Equal(t, []int{0, 1, 2}, xslices.TrimRight([]int{0, 1, 2, 0, 0}, []int{0}))
	require.Empty(t, xslices.TrimRight([]int{0}, []int{0}))
	require.Empty(t, xslices.TrimRight([]int(nil), []int{0}))
}
