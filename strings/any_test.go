package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestSplitAny(t *testing.T) {
	t.Parallel()

	require.Equal(t, []string{"a", "b", "c"}, xstrings.SplitAny("a,b;c", ",;"))

	// Empty segments between adjacent separators are preserved.
	require.Equal(t, []string{"a", "", "b"}, xstrings.SplitAny("a,;b", ",;"))
	require.Equal(t, []string{"", "a", ""}, xstrings.SplitAny(",a;", ",;"))

	// No separators present.
	require.Equal(t, []string{"abc"}, xstrings.SplitAny("abc", ",;"))
	require.Equal(t, []string{""}, xstrings.SplitAny("", ",;"))

	// Empty cutset returns s whole.
	require.Equal(t, []string{"a,b"}, xstrings.SplitAny("a,b", ""))

	// Multi-byte separators split on full code points.
	require.Equal(t, []string{"a", "b"}, xstrings.SplitAny("a→b", "→"))
}

func TestCountAny(t *testing.T) {
	t.Parallel()

	require.Equal(t, 2, xstrings.CountAny("a,b;c", ",;"))
	require.Equal(t, 0, xstrings.CountAny("abc", ",;"))
	require.Equal(t, 0, xstrings.CountAny("", ",;"))
	require.Equal(t, 0, xstrings.CountAny("a,b", ""))

	// Multi-byte code points count once each.
	require.Equal(t, 2, xstrings.CountAny("a→b→c", "→"))
}
