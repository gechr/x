package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestSortNatural(t *testing.T) {
	t.Parallel()

	got := []string{"item10", "item2", "item1", "item20", "item3"}
	xslices.SortNatural(got)
	require.Equal(t, []string{"item1", "item2", "item3", "item10", "item20"}, got)
}

func TestSortNaturalNamedType(t *testing.T) {
	t.Parallel()

	type name string
	got := []name{"v10", "v9", "v1"}
	xslices.SortNatural(got)
	require.Equal(t, []name{"v1", "v9", "v10"}, got)
}
