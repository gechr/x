package set_test

import (
	"slices"
	"testing"

	"github.com/gechr/x/set"
	"github.com/stretchr/testify/require"
)

func TestNewAddDelete(t *testing.T) {
	t.Parallel()

	s := set.New("a", "b", "a")
	require.Equal(t, 2, s.Len())
	require.True(t, s.Contains("a"))
	require.False(t, s.Contains("c"))

	s.Add("c", "d")
	require.Equal(t, 4, s.Len())

	s.Delete("a", "c", "missing")
	require.Equal(t, 2, s.Len())
	require.False(t, s.Contains("a"))

	// Empty set.
	require.Equal(t, 0, set.New[string]().Len())

	// Nil set supports reads.
	var nilSet set.Set[string]
	require.Equal(t, 0, nilSet.Len())
	require.False(t, nilSet.Contains("a"))
}

func TestCollect(t *testing.T) {
	t.Parallel()

	s := set.Collect(slices.Values([]int{1, 2, 2, 3}))
	require.Equal(t, 3, s.Len())
	require.True(t, s.Contains(2))
}

func TestEqualSubsetOf(t *testing.T) {
	t.Parallel()

	require.True(t, set.New(1, 2).Equal(set.New(2, 1)))
	require.False(t, set.New(1, 2).Equal(set.New(1, 3)))
	require.False(t, set.New(1).Equal(set.New(1, 2)))
	require.True(t, set.New[int]().Equal(nil))

	require.True(t, set.New(1, 2).SubsetOf(set.New(1, 2, 3)))
	require.True(t, set.New(1, 2).SubsetOf(set.New(1, 2)))
	require.True(t, set.New[int]().SubsetOf(set.New(1)))
	require.False(t, set.New(1, 4).SubsetOf(set.New(1, 2, 3)))
	require.False(t, set.New(1, 2, 3).SubsetOf(set.New(1, 2)))
}

func TestSetOperations(t *testing.T) {
	t.Parallel()

	a := set.New(1, 2, 3)
	b := set.New(3, 4)
	c := set.New(5)

	require.True(t, a.Union(b).Equal(set.New(1, 2, 3, 4)))
	require.True(t, a.Union(b, c).Equal(set.New(1, 2, 3, 4, 5)))
	require.True(t, a.Union().Equal(a))
	require.True(t, a.Intersect(b).Equal(set.New(3)))
	require.True(t, b.Intersect(a).Equal(set.New(3)))
	require.True(t, a.Intersect(b, c).Equal(set.New[int]()))
	require.True(t, a.Intersect().Equal(a))
	require.True(t, a.Difference(b).Equal(set.New(1, 2)))
	require.True(t, b.Difference(a).Equal(set.New(4)))
	require.True(t, a.Difference(b, c).Equal(set.New(1, 2)))
	require.True(t, a.Difference().Equal(a))

	// Operands are not modified.
	require.Equal(t, 3, a.Len())
	require.Equal(t, 2, b.Len())
}

func TestCloneSliceAll(t *testing.T) {
	t.Parallel()

	s := set.New("a", "b")

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

	// Early break stops iteration.
	var count int
	for range s.All() {
		count++
		break
	}
	require.Equal(t, 1, count)
}
