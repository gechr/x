package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestSurround(t *testing.T) {
	t.Parallel()

	require.Equal(t,
		[]string{`"a"`, `"b"`, `"c"`},
		xslices.Surround([]string{"a", "b", "c"}, `"`, `"`),
	)

	// Different prefix and suffix.
	require.Equal(t,
		[]string{"(a)", "(b)"},
		xslices.Surround([]string{"a", "b"}, "(", ")"),
	)

	// Empty prefix/suffix leaves elements unchanged.
	require.Equal(t, []string{"a", "b"}, xslices.Surround([]string{"a", "b"}, "", ""))

	require.Empty(t, xslices.Surround([]string(nil), "(", ")"))

	// Named slice types are accepted; the result is a plain []string.
	type tags []string
	require.Equal(t, []string{"[x]"}, xslices.Surround(tags{"x"}, "[", "]"))
}
