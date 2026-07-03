package maps_test

import (
	"slices"
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestGroup(t *testing.T) {
	t.Parallel()

	pairs := [][2]string{
		{"fruit", "apple"},
		{"veg", "carrot"},
		{"fruit", "banana"},
	}
	seq := func(yield func(string, string) bool) {
		for _, p := range pairs {
			if !yield(p[0], p[1]) {
				return
			}
		}
	}

	require.Equal(t, map[string][]string{
		"fruit": {"apple", "banana"},
		"veg":   {"carrot"},
	}, xmaps.Group(seq))

	// Empty sequence yields an empty (non-nil) map.
	empty := xmaps.Group(func(func(string, string) bool) {})
	require.NotNil(t, empty)
	require.Empty(t, empty)
}

func TestGroupFunc(t *testing.T) {
	t.Parallel()

	words := []string{"apple", "avocado", "banana", "blueberry", "cherry"}

	byInitial := xmaps.GroupFunc(slices.Values(words), func(s string) byte { return s[0] })
	require.Equal(t, map[byte][]string{
		'a': {"apple", "avocado"},
		'b': {"banana", "blueberry"},
		'c': {"cherry"},
	}, byInitial)

	// Grouping preserves encounter order within each group.
	byLen := xmaps.GroupFunc(slices.Values(words), func(s string) int { return len(s) })
	require.Equal(t, []string{"banana", "cherry"}, byLen[6])
}
