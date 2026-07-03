package set_test

import (
	"slices"
	"testing"

	"github.com/gechr/x/set"
	"github.com/stretchr/testify/require"
)

func TestSortedSetNewAddDelete(t *testing.T) {
	t.Parallel()

	s := set.NewSorted("charlie", "alpha", "beta", "alpha")
	require.Equal(t, 3, s.Len())
	require.Equal(t, []string{"alpha", "beta", "charlie"}, s.Slice())
	require.True(t, s.Contains("alpha"))
	require.False(t, s.Contains("delta"))

	s.Add("delta")
	require.Equal(t, []string{"alpha", "beta", "charlie", "delta"}, s.Slice())

	s.Delete("beta", "missing")
	require.Equal(t, []string{"alpha", "charlie", "delta"}, s.Slice())

	// Empty set.
	require.Equal(t, 0, set.NewSorted[string]().Len())

	// Zero value is usable.
	var zero set.SortedSet[string]
	require.Equal(t, 0, zero.Len())
	zero.Add("a")
	require.Equal(t, []string{"a"}, zero.Slice())
}

func TestCollectSorted(t *testing.T) {
	t.Parallel()

	s := set.CollectSorted(slices.Values([]int{3, 1, 2, 2}))
	require.Equal(t, 3, s.Len())
	require.Equal(t, []int{1, 2, 3}, s.Slice())
}

func TestSortedSetEqualSubsetOf(t *testing.T) {
	t.Parallel()

	require.True(t, set.NewSorted(1, 2).Equal(set.NewSorted(2, 1)))
	require.False(t, set.NewSorted(1, 2).Equal(set.NewSorted(1, 3)))

	require.True(t, set.NewSorted(1, 2).SubsetOf(set.NewSorted(1, 2, 3)))
	require.False(t, set.NewSorted(1, 4).SubsetOf(set.NewSorted(1, 2, 3)))
}

func TestSortedSetOperations(t *testing.T) {
	t.Parallel()

	a := set.NewSorted(3, 1, 2)
	b := set.NewSorted(4, 3)
	c := set.NewSorted(5)

	require.Equal(t, []int{1, 2, 3, 4}, a.Union(b).Slice())
	require.Equal(t, []int{1, 2, 3, 4, 5}, a.Union(b, c).Slice())
	require.Equal(t, []int{1, 2, 3}, a.Union().Slice())
	require.Equal(t, []int{3}, a.Intersect(b).Slice())
	require.Empty(t, a.Intersect(b, c).Slice())
	require.Equal(t, []int{1, 2}, a.Difference(b).Slice())
	require.Equal(t, []int{1, 2, 3}, a.Difference().Slice())

	// Operands are not modified.
	require.Equal(t, 3, a.Len())
	require.Equal(t, 2, b.Len())
}

func TestSortedSetCloneSliceAll(t *testing.T) {
	t.Parallel()

	s := set.NewSorted("b", "a")

	clone := s.Clone()
	clone.Add("c")
	require.Equal(t, 2, s.Len())
	require.Equal(t, 3, clone.Len())
	require.Equal(t, []string{"a", "b", "c"}, clone.Slice())

	var items []string
	for item := range s.All() {
		items = append(items, item)
	}
	require.Equal(t, []string{"a", "b"}, items)
}

func TestSortedSetMergeAlreadySorted(t *testing.T) {
	t.Parallel()

	a := []string{"charlie", "alpha"}
	b := []string{"beta", "alpha"}

	got := set.NewSorted(a...).Union(set.NewSorted(b...)).Slice()
	require.Equal(t, []string{"alpha", "beta", "charlie"}, got)
}
