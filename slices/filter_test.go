package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }
	require.Equal(t, []int{2, 4}, xslices.Filter([]int{1, 2, 3, 4}, isEven))
	require.Empty(t, xslices.Filter([]int{1, 3}, isEven))

	// The input is not modified.
	items := []int{1, 2, 3}
	filtered := xslices.Filter(items, func(n int) bool { return n > 1 })
	require.Equal(t, []int{1, 2, 3}, items)
	require.Equal(t, []int{2, 3}, filtered)

	// Empty and nil inputs yield empty non-nil slices.
	filtered = xslices.Filter([]int(nil), isEven)
	require.Empty(t, filtered)
	require.NotNil(t, filtered)

	// Named slice types are preserved.
	type numbers []int
	named := xslices.Filter(numbers{1, 2, 3}, func(n int) bool { return n != 2 })
	require.IsType(t, numbers{}, named)
	require.Equal(t, numbers{1, 3}, named)
}
