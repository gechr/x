package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	t.Parallel()

	require.Equal(t, 3, xslices.Count([]string{"a", "b", "a", "a"}, "a"))
	require.Equal(t, 1, xslices.Count([]string{"a", "b"}, "b"))
	require.Equal(t, 0, xslices.Count([]string{"a", "b"}, "c"))
	require.Equal(t, 0, xslices.Count([]string(nil), "a"))

	// Named slice types satisfy the constraint.
	type tags []string
	require.Equal(t, 2, xslices.Count(tags{"x", "x"}, "x"))
}

func TestCountFunc(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }

	require.Equal(t, 3, xslices.CountFunc([]int{2, 1, 4, 3, 6}, isEven))
	require.Equal(t, 0, xslices.CountFunc([]int{1, 3, 5}, isEven))
	require.Equal(t, 0, xslices.CountFunc([]int(nil), isEven))
}
