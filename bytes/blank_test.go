package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
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
		require.Equal(t, want, xbytes.IsBlank([]byte(in)), "IsBlank(%q)", in)
	}
}

func TestAnyNonEmpty(t *testing.T) {
	t.Parallel()

	require.False(t, xbytes.AnyNonEmpty())
	require.False(t, xbytes.AnyNonEmpty([]byte(""), []byte(""), []byte("")))
	require.True(t, xbytes.AnyNonEmpty([]byte(""), []byte("x"), []byte("")))
	require.True(t, xbytes.AnyNonEmpty([]byte(" ")))
}

func TestAnyEmpty(t *testing.T) {
	t.Parallel()

	require.False(t, xbytes.AnyEmpty())
	require.False(t, xbytes.AnyEmpty([]byte("x"), []byte("y")))
	require.True(t, xbytes.AnyEmpty([]byte("x"), []byte(""), []byte("y")))
	require.False(t, xbytes.AnyEmpty([]byte(" ")))
}

func TestAllEmpty(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.AllEmpty())
	require.True(t, xbytes.AllEmpty([]byte(""), []byte(""), []byte("")))
	require.False(t, xbytes.AllEmpty([]byte(""), []byte("x"), []byte("")))
	require.False(t, xbytes.AllEmpty([]byte(" ")))
}

func TestAllNonEmpty(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.AllNonEmpty())
	require.True(t, xbytes.AllNonEmpty([]byte("x"), []byte("y")))
	require.False(t, xbytes.AllNonEmpty([]byte("x"), []byte(""), []byte("y")))
	require.True(t, xbytes.AllNonEmpty([]byte(" ")))
}
