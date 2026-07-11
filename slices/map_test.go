package slices_test

import (
	"strconv"
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Parallel()

	// Type-changing transform; function values pass directly.
	require.Equal(t, []string{"1", "2", "3"}, xslices.Map([]int{1, 2, 3}, strconv.Itoa))

	// Same-type transform.
	double := func(n int) int { return n * 2 }
	require.Equal(t, []int{2, 4, 6}, xslices.Map([]int{1, 2, 3}, double))

	// Empty and nil inputs yield empty (non-nil) slices.
	out := xslices.Map([]int(nil), strconv.Itoa)
	require.Empty(t, out)
	require.NotNil(t, out)

	// Named slice types are accepted as input.
	type ids []int
	require.Equal(t, []string{"7", "8"}, xslices.Map(ids{7, 8}, strconv.Itoa))
}
