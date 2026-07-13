package strings_test

import (
	"strings"
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestCompareFold(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0, xstrings.CompareFold("go", "GO"))
	require.Equal(t, 0, xstrings.CompareFold("", ""))
	require.Negative(t, xstrings.CompareFold("alpha", "BETA"))
	require.Positive(t, xstrings.CompareFold("BETA", "alpha"))

	// A prefix sorts before its extension.
	require.Negative(t, xstrings.CompareFold("go", "GOPHER"))
	require.Positive(t, xstrings.CompareFold("GOPHER", "go"))

	// Greek final sigma folds to sigma under simple case-folding.
	require.Equal(t, 0, xstrings.CompareFold("ΟΔΟΣ", "οδος"))

	// Zero iff EqualFold, including for orbit members ToLower would miss.
	for _, pair := range [][2]string{
		{"kelvin", "KELVIN"},
		{"ς", "Σ"},
		{"abc", "abd"},
		{"", "a"},
	} {
		a, b := pair[0], pair[1]
		require.Equal(t, strings.EqualFold(a, b), xstrings.CompareFold(a, b) == 0,
			"EqualFold(%q, %q) mismatch", a, b)
	}
}

func TestContainsFold(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.ContainsFold("Hello, World", "WORLD"))
	require.True(t, xstrings.ContainsFold("aKb", "KB"))
	require.True(t, xstrings.ContainsFold("ΟΔΟΣ", "δος"))
	require.True(t, xstrings.ContainsFold("anything", ""))

	require.False(t, xstrings.ContainsFold("Hello, World", "moon"))
	require.False(t, xstrings.ContainsFold("Straße", "STRASSE"))
	require.False(t, xstrings.ContainsFold("", "x"))
}

func TestHasPrefixFold(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.HasPrefixFold("Hello, World", "HELLO"))
	require.True(t, xstrings.HasPrefixFold("Kelvin", "kel"))
	require.True(t, xstrings.HasPrefixFold("anything", ""))

	require.False(t, xstrings.HasPrefixFold("Hello, World", "world"))
	require.False(t, xstrings.HasPrefixFold("short", "shorter"))
}

func TestHasSuffixFold(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.HasSuffixFold("Hello, World", "WORLD"))
	require.True(t, xstrings.HasSuffixFold("ΟΔΟΣ", "δος"))
	require.True(t, xstrings.HasSuffixFold("anything", ""))

	require.False(t, xstrings.HasSuffixFold("Hello, World", "hello"))
	require.False(t, xstrings.HasSuffixFold("short", "longer-short"))
}
