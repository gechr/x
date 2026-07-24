package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestReject(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }
	require.Equal(t, []int{1, 3}, xslices.Reject([]int{1, 2, 3, 4}, isEven))
	require.Empty(t, xslices.Reject([]int{2, 4}, isEven))

	// The input is not modified.
	items := []int{1, 2, 3}
	rejected := xslices.Reject(items, func(n int) bool { return n > 1 })
	require.Equal(t, []int{1, 2, 3}, items)
	require.Equal(t, []int{1}, rejected)

	// Empty and nil inputs yield empty non-nil slices.
	rejected = xslices.Reject([]int(nil), isEven)
	require.Empty(t, rejected)
	require.NotNil(t, rejected)

	// Named slice types are preserved.
	type numbers []int
	named := xslices.Reject(numbers{1, 2, 3}, func(n int) bool { return n != 2 })
	require.IsType(t, numbers{}, named)
	require.Equal(t, numbers{2}, named)
}
