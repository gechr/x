package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestSplitAny(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		[][]byte{[]byte("a"), []byte("b"), []byte("c")},
		xbytes.SplitAny([]byte("a,b;c"), ",;"),
	)

	// Empty segments between adjacent separators are preserved.
	require.Equal(t, [][]byte{[]byte("a"), {}, []byte("b")}, xbytes.SplitAny([]byte("a,;b"), ",;"))
	require.Equal(t, [][]byte{{}, []byte("a"), {}}, xbytes.SplitAny([]byte(",a;"), ",;"))

	// No separators present.
	require.Equal(t, [][]byte{[]byte("abc")}, xbytes.SplitAny([]byte("abc"), ",;"))
	require.Equal(t, [][]byte{{}}, xbytes.SplitAny([]byte(""), ",;"))

	// Empty cutset returns s whole.
	require.Equal(t, [][]byte{[]byte("a,b")}, xbytes.SplitAny([]byte("a,b"), ""))

	// Multi-byte separators split on full code points.
	require.Equal(t, [][]byte{[]byte("a"), []byte("b")}, xbytes.SplitAny([]byte("a→b"), "→"))
}

func TestCountAny(t *testing.T) {
	t.Parallel()

	require.Equal(t, 2, xbytes.CountAny([]byte("a,b;c"), ",;"))
	require.Equal(t, 0, xbytes.CountAny([]byte("abc"), ",;"))
	require.Equal(t, 0, xbytes.CountAny([]byte(""), ",;"))
	require.Equal(t, 0, xbytes.CountAny([]byte("a,b"), ""))

	// Multi-byte code points count once each.
	require.Equal(t, 2, xbytes.CountAny([]byte("a→b→c"), "→"))
}
