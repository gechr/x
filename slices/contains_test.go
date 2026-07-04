package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestContainsAny(t *testing.T) {
	t.Parallel()

	require.True(t, xslices.ContainsAny("beta", []string{"alpha", "beta"}))
	require.True(t, xslices.ContainsAny("beta", []string{"alpha"}, []string{"beta"}))
	require.True(t, xslices.ContainsAny(2, []int{1}, []int{2, 3}))

	require.False(t, xslices.ContainsAny("gamma", []string{"alpha"}, []string{"beta"}))
	require.False(t, xslices.ContainsAny("alpha", []string(nil), []string{}))
	require.False(t, xslices.ContainsAny[[]string]("alpha"))
}

func TestContainsAll(t *testing.T) {
	t.Parallel()

	require.True(t, xslices.ContainsAll("beta", []string{"alpha", "beta"}))
	require.True(t, xslices.ContainsAll("beta", []string{"beta"}, []string{"beta", "gamma"}))
	require.True(t, xslices.ContainsAll[[]string]("alpha")) // vacuously true

	require.False(t, xslices.ContainsAll("beta", []string{"beta"}, []string{"alpha"}))
	require.False(t, xslices.ContainsAll("beta", []string(nil), []string{"beta"}))
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
