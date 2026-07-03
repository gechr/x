package set_test

import (
	"slices"
	"testing"

	"github.com/gechr/x/set"
	"github.com/stretchr/testify/require"
)

func TestOrderedNewAddDelete(t *testing.T) {
	t.Parallel()

	s := set.NewOrdered("a", "b", "a")
	require.Equal(t, 2, s.Len())
	require.True(t, s.Contains("a"))
	require.False(t, s.Contains("c"))

	s.Add("c", "d")
	require.Equal(t, 4, s.Len())

	s.Delete("a", "c", "missing")
	require.Equal(t, 2, s.Len())
	require.False(t, s.Contains("a"))

	// Empty set.
	require.Equal(t, 0, set.NewOrdered[string]().Len())
}

func TestCollectOrdered(t *testing.T) {
	t.Parallel()

	s := set.CollectOrdered(slices.Values([]int{1, 2, 2, 3}))
	require.Equal(t, 3, s.Len())
	require.True(t, s.Contains(2))
}

func TestOrderedEqualSubsetOf(t *testing.T) {
	t.Parallel()

	require.True(t, set.NewOrdered(1, 2).Equal(set.NewOrdered(2, 1)))
	require.False(t, set.NewOrdered(1, 2).Equal(set.NewOrdered(1, 3)))

	require.True(t, set.NewOrdered(1, 2).SubsetOf(set.NewOrdered(1, 2, 3)))
	require.False(t, set.NewOrdered(1, 4).SubsetOf(set.NewOrdered(1, 2, 3)))
}

func TestOrderedSetOperations(t *testing.T) {
	t.Parallel()

	a := set.NewOrdered(1, 2, 3)
	b := set.NewOrdered(3, 4)
	c := set.NewOrdered(5)

	require.True(t, a.Union(b).Equal(set.NewOrdered(1, 2, 3, 4)))
	require.True(t, a.Union(b, c).Equal(set.NewOrdered(1, 2, 3, 4, 5)))
	require.True(t, a.Union().Equal(a))
	require.True(t, a.Intersect(b).Equal(set.NewOrdered(3)))
	require.True(t, a.Intersect(b, c).Equal(set.NewOrdered[int]()))
	require.True(t, a.Difference(b).Equal(set.NewOrdered(1, 2)))
	require.True(t, a.Difference().Equal(a))

	// Operands are not modified.
	require.Equal(t, 3, a.Len())
	require.Equal(t, 2, b.Len())
}

func TestOrderedCloneSliceAllSort(t *testing.T) {
	t.Parallel()

	s := set.NewOrdered("b", "a")

	clone := s.Clone()
	clone.Add("c")
	require.Equal(t, 2, s.Len())
	require.Equal(t, 3, clone.Len())

	require.ElementsMatch(t, []string{"a", "b"}, s.Slice())

	var items []string
	for item := range s.All() {
		items = append(items, item)
	}
	require.ElementsMatch(t, []string{"a", "b"}, items)

	require.Equal(t, []string{"a", "b"}, s.Sort())
}

func TestOrderedSortChain(t *testing.T) {
	t.Parallel()

	a := []string{"charlie", "alpha"}
	b := []string{"beta", "alpha"}

	got := set.NewOrdered(a...).Union(set.NewOrdered(b...)).Sort()
	require.Equal(t, []string{"alpha", "beta", "charlie"}, got)
}
