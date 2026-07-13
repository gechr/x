package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestContainsAll(t *testing.T) {
	t.Parallel()

	require.True(t, xslices.ContainsAll([]string{"alpha", "beta"}, "alpha"))
	require.True(t, xslices.ContainsAll([]string{"alpha", "beta"}, "alpha", "beta"))
	require.True(t, xslices.ContainsAll([]string{"alpha", "beta"})) // vacuously true

	require.False(t, xslices.ContainsAll([]string{"alpha", "beta"}, "alpha", "gamma"))
	require.False(t, xslices.ContainsAll([]string(nil), "alpha"))
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	require.True(t, xslices.ContainsAny([]string{"alpha", "beta"}, "beta"))
	require.True(t, xslices.ContainsAny([]string{"alpha", "beta"}, "gamma", "beta"))
	require.True(t, xslices.ContainsAny([]int{1, 2, 3}, 2, 4))

	require.False(t, xslices.ContainsAny([]string{"alpha", "beta"}, "gamma", "delta"))
	require.False(t, xslices.ContainsAny([]string(nil), "alpha"))
	require.False(t, xslices.ContainsAny([]string{"alpha"}))
}

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
