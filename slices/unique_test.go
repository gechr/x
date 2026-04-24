package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		[]string{"a", "b", "A", "c"},
		xslices.Unique([]string{"a", "b", "a", "A", "c", "b"}),
	)
	require.Equal(t, []int{1, 2, 3}, xslices.Unique([]int{1, 2, 1, 3, 2}))
	require.Equal(t, []string{}, xslices.Unique([]string{}))
}

func TestUniqueFold(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "B", "c"}, xslices.UniqueFold([]string{"a", "B", "A", "b", "c"}))
	require.Equal(
		t,
		[]namedString{"one", "two"},
		xslices.UniqueFold([]namedString{"one", "ONE", "two"}),
	)
}

type namedString string
