package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestContainsAll(t *testing.T) {
	t.Parallel()

	require.True(
		t,
		xbytes.ContainsAll([]byte("foo bar baz"), []byte("foo"), []byte("bar"), []byte("baz")),
	)
	require.True(t, xbytes.ContainsAll([]byte("foo"), []byte("foo")))
	require.True(t, xbytes.ContainsAll([]byte("anything")))
	require.False(t, xbytes.ContainsAll([]byte("foo bar"), []byte("foo"), []byte("baz")))
	require.False(t, xbytes.ContainsAll([]byte(""), []byte("x")))
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.ContainsAny([]byte("foo bar baz"), []byte("baz"), []byte("qux")))
	require.True(t, xbytes.ContainsAny([]byte("foo"), []byte("foo"), []byte("bar")))
	require.False(t, xbytes.ContainsAny([]byte("foo bar")))
	require.False(t, xbytes.ContainsAny([]byte("foo bar"), []byte("baz"), []byte("qux")))
	require.False(t, xbytes.ContainsAny([]byte(""), []byte("x")))
}
