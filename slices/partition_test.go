package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestPartition(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }

	even, odd := xslices.Partition([]int{1, 2, 3, 4, 5, 6}, isEven)
	require.Equal(t, []int{2, 4, 6}, even)
	require.Equal(t, []int{1, 3, 5}, odd)

	// All elements on one side.
	even, odd = xslices.Partition([]int{2, 4}, isEven)
	require.Equal(t, []int{2, 4}, even)
	require.Empty(t, odd)

	// Empty and nil inputs yield empty (non-nil) slices.
	even, odd = xslices.Partition([]int(nil), isEven)
	require.Empty(t, even)
	require.Empty(t, odd)
	require.NotNil(t, even)
	require.NotNil(t, odd)

	// Named slice types are preserved.
	type tags []string
	long, short := xslices.Partition(
		tags{"go", "rust", "c"},
		func(s string) bool { return len(s) > 1 },
	)
	require.IsType(t, tags{}, long)
	require.Equal(t, tags{"go", "rust"}, long)
	require.Equal(t, tags{"c"}, short)
}
