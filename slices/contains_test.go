package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestContainsFold(t *testing.T) {
	t.Parallel()

	require.True(t, xslices.ContainsFold([]string{"alpha", "beta"}, "alpha"))
	require.True(t, xslices.ContainsFold([]string{"alpha", "beta"}, "BETA"))
	// Greek final sigma folds to sigma under simple case-folding.
	require.True(t, xslices.ContainsFold([]string{"ΟΔΟΣ"}, "οδος"))

	require.False(t, xslices.ContainsFold([]string{"alpha", "beta"}, "gamma"))
	require.False(t, xslices.ContainsFold([]string(nil), "alpha"))
	require.False(t, xslices.ContainsFold([]string{}, ""))

	// Named string types satisfy the constraint.
	type tag string
	require.True(t, xslices.ContainsFold([]tag{"Latest"}, tag("latest")))
}
