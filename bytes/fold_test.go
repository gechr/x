package bytes_test

import (
	"bytes"
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestCompareFold(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0, xbytes.CompareFold([]byte("go"), []byte("GO")))
	require.Equal(t, 0, xbytes.CompareFold([]byte(""), []byte("")))
	require.Negative(t, xbytes.CompareFold([]byte("alpha"), []byte("BETA")))
	require.Positive(t, xbytes.CompareFold([]byte("BETA"), []byte("alpha")))

	// A prefix sorts before its extension.
	require.Negative(t, xbytes.CompareFold([]byte("go"), []byte("GOPHER")))
	require.Positive(t, xbytes.CompareFold([]byte("GOPHER"), []byte("go")))

	// Greek final sigma folds to sigma under simple case-folding.
	require.Equal(t, 0, xbytes.CompareFold([]byte("ΟΔΟΣ"), []byte("οδος")))

	// Zero iff EqualFold, including for orbit members ToLower would miss.
	for _, pair := range [][2]string{
		{"kelvin", "KELVIN"},
		{"ς", "Σ"},
		{"abc", "abd"},
		{"", "a"},
	} {
		a, b := []byte(pair[0]), []byte(pair[1])
		require.Equal(t, bytes.EqualFold(a, b), xbytes.CompareFold(a, b) == 0,
			"EqualFold(%q, %q) mismatch", a, b)
	}
}

func TestContainsFold(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.ContainsFold([]byte("Hello, World"), []byte("WORLD")))
	require.True(t, xbytes.ContainsFold([]byte("aKb"), []byte("KB")))
	require.True(t, xbytes.ContainsFold([]byte("ΟΔΟΣ"), []byte("δος")))
	require.True(t, xbytes.ContainsFold([]byte("anything"), []byte("")))

	require.False(t, xbytes.ContainsFold([]byte("Hello, World"), []byte("moon")))
	require.False(t, xbytes.ContainsFold([]byte("Straße"), []byte("STRASSE")))
	require.False(t, xbytes.ContainsFold([]byte(""), []byte("x")))
}

func TestHasPrefixFold(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.HasPrefixFold([]byte("Hello, World"), []byte("HELLO")))
	require.True(t, xbytes.HasPrefixFold([]byte("Kelvin"), []byte("kel")))
	require.True(t, xbytes.HasPrefixFold([]byte("anything"), []byte("")))

	require.False(t, xbytes.HasPrefixFold([]byte("Hello, World"), []byte("world")))
	require.False(t, xbytes.HasPrefixFold([]byte("short"), []byte("shorter")))
}

func TestHasSuffixFold(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.HasSuffixFold([]byte("Hello, World"), []byte("WORLD")))
	require.True(t, xbytes.HasSuffixFold([]byte("ΟΔΟΣ"), []byte("δος")))
	require.True(t, xbytes.HasSuffixFold([]byte("anything"), []byte("")))

	require.False(t, xbytes.HasSuffixFold([]byte("Hello, World"), []byte("hello")))
	require.False(t, xbytes.HasSuffixFold([]byte("short"), []byte("longer-short")))
}
