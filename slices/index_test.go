package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestLastIndex(t *testing.T) {
	t.Parallel()

	require.Equal(t, 3, xslices.LastIndex([]string{"a", "b", "a", "a"}, "a"))
	require.Equal(t, 1, xslices.LastIndex([]string{"a", "b"}, "b"))
	require.Equal(t, -1, xslices.LastIndex([]string{"a", "b"}, "c"))
	require.Equal(t, -1, xslices.LastIndex([]string(nil), "a"))
	require.Equal(t, -1, xslices.LastIndex([]string{}, ""))

	// Named slice types satisfy the constraint.
	type tags []string
	require.Equal(t, 0, xslices.LastIndex(tags{"x"}, "x"))
}

func TestLastIndexFunc(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }

	require.Equal(t, 4, xslices.LastIndexFunc([]int{2, 1, 4, 3, 6}, isEven))
	require.Equal(t, 0, xslices.LastIndexFunc([]int{2, 1, 3}, isEven))
	require.Equal(t, -1, xslices.LastIndexFunc([]int{1, 3, 5}, isEven))
	require.Equal(t, -1, xslices.LastIndexFunc([]int(nil), isEven))
}
