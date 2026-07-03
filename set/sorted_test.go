package set_test

import (
	"testing"

	"github.com/gechr/x/set"
	"github.com/stretchr/testify/require"
)

func TestSorted(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"alpha", "beta", "charlie"},
		set.Sorted(set.New("charlie", "alpha", "beta")))
	require.Equal(t, []int{1, 2, 3}, set.Sorted(set.New(3, 1, 2)))

	// Empty and nil sets yield empty (non-nil) slices.
	require.Empty(t, set.Sorted(set.New[int]()))
	require.Empty(t, set.Sorted(set.Set[int](nil)))

	// The source set is not modified.
	s := set.New(2, 1)
	_ = set.Sorted(s)
	require.Equal(t, 2, s.Len())
}
