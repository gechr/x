package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsBlank(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":          true,
		" ":         true,
		"\t\n  \t":  true,
		"x":         false,
		"  hello  ": false,
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsBlank(in), "IsBlank(%q)", in)
	}
}

func TestAnyNonEmpty(t *testing.T) {
	t.Parallel()

	require.False(t, xstrings.AnyNonEmpty())
	require.False(t, xstrings.AnyNonEmpty("", "", ""))
	require.True(t, xstrings.AnyNonEmpty("", "x", ""))
	require.True(t, xstrings.AnyNonEmpty(" "))
}

func TestAnyEmpty(t *testing.T) {
	t.Parallel()

	require.False(t, xstrings.AnyEmpty())
	require.False(t, xstrings.AnyEmpty("x", "y"))
	require.True(t, xstrings.AnyEmpty("x", "", "y"))
	require.False(t, xstrings.AnyEmpty(" "))
}

func TestAllEmpty(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.AllEmpty())
	require.True(t, xstrings.AllEmpty("", "", ""))
	require.False(t, xstrings.AllEmpty("", "x", ""))
	require.False(t, xstrings.AllEmpty(" "))
}

func TestAllNonEmpty(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.AllNonEmpty())
	require.True(t, xstrings.AllNonEmpty("x", "y"))
	require.False(t, xstrings.AllNonEmpty("x", "", "y"))
	require.True(t, xstrings.AllNonEmpty(" "))
}
