package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestAllFunc(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }

	require.True(t, xslices.AllFunc([]int{2, 4, 6}, isEven))
	require.False(t, xslices.AllFunc([]int{2, 3, 4}, isEven))
	require.False(t, xslices.AllFunc([]int{1, 3, 5}, isEven))

	// An empty slice is vacuously true.
	require.True(t, xslices.AllFunc([]int{}, isEven))
	require.True(t, xslices.AllFunc([]int(nil), isEven))

	// Named slice types satisfy the constraint.
	type tags []string
	require.True(t, xslices.AllFunc(tags{"a", "b"}, func(s string) bool { return s != "" }))
}

func TestAnyFunc(t *testing.T) {
	t.Parallel()

	isEven := func(n int) bool { return n%2 == 0 }

	require.True(t, xslices.AnyFunc([]int{1, 2, 3}, isEven))
	require.True(t, xslices.AnyFunc([]int{2, 4, 6}, isEven))
	require.False(t, xslices.AnyFunc([]int{1, 3, 5}, isEven))

	// An empty slice reports false.
	require.False(t, xslices.AnyFunc([]int{}, isEven))
	require.False(t, xslices.AnyFunc([]int(nil), isEven))

	// Named slice types satisfy the constraint.
	type tags []string
	require.True(t, xslices.AnyFunc(tags{"", "b"}, func(s string) bool { return s != "" }))
}

func TestAllFuncShortCircuits(t *testing.T) {
	t.Parallel()

	var calls int
	require.False(t, xslices.AllFunc([]int{1, 2, 3}, func(n int) bool {
		calls++
		return n != 2
	}))
	require.Equal(t, 2, calls)
}

func TestAnyFuncShortCircuits(t *testing.T) {
	t.Parallel()

	var calls int
	require.True(t, xslices.AnyFunc([]int{1, 2, 3}, func(n int) bool {
		calls++
		return n == 2
	}))
	require.Equal(t, 2, calls)
}
