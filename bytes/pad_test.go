package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestPadLeft(t *testing.T) {
	t.Parallel()

	require.Equal(t, "   hi", string(xbytes.PadLeft([]byte("hi"), 5)))
	require.Equal(t, "hello", string(xbytes.PadLeft([]byte("hello"), 5)))
	require.Equal(t, "hello!", string(xbytes.PadLeft([]byte("hello!"), 5)))
	require.Equal(t, "   ", string(xbytes.PadLeft([]byte(""), 3)))
	require.Equal(t, "hi", string(xbytes.PadLeft([]byte("hi"), 0)))
	require.Equal(t, "hi", string(xbytes.PadLeft([]byte("hi"), -1)))
	// Width counts runes, not bytes.
	require.Equal(t, "  héé", string(xbytes.PadLeft([]byte("héé"), 5)))
}

func TestPadRight(t *testing.T) {
	t.Parallel()

	require.Equal(t, "hi   ", string(xbytes.PadRight([]byte("hi"), 5)))
	require.Equal(t, "hello", string(xbytes.PadRight([]byte("hello"), 5)))
	require.Equal(t, "hello!", string(xbytes.PadRight([]byte("hello!"), 5)))
	require.Equal(t, "   ", string(xbytes.PadRight([]byte(""), 3)))
	require.Equal(t, "hi", string(xbytes.PadRight([]byte("hi"), 0)))
	require.Equal(t, "héé  ", string(xbytes.PadRight([]byte("héé"), 5)))
}

func TestPadCenter(t *testing.T) {
	t.Parallel()

	require.Equal(t, " hi  ", string(xbytes.PadCenter([]byte("hi"), 5)))
	require.Equal(t, " hi ", string(xbytes.PadCenter([]byte("hi"), 4)))
	require.Equal(t, "hello", string(xbytes.PadCenter([]byte("hello"), 5)))
	require.Equal(t, "hello!", string(xbytes.PadCenter([]byte("hello!"), 5)))
	require.Equal(t, "   ", string(xbytes.PadCenter([]byte(""), 3)))
	require.Equal(t, "hi", string(xbytes.PadCenter([]byte("hi"), -1)))
	require.Equal(t, " héé ", string(xbytes.PadCenter([]byte("héé"), 5)))
}

// The pad helpers guarantee the returned slice never aliases `s`, so padding
// must not write through the input's backing array or spare capacity.
func TestPadNoAlias(t *testing.T) {
	t.Parallel()

	for _, pad := range []func([]byte, int) []byte{
		xbytes.PadLeft,
		xbytes.PadRight,
		xbytes.PadCenter,
	} {
		in := make([]byte, 2, 8) // spare capacity is the aliasing hazard
		copy(in, "hi")
		_ = pad(in, 5)
		require.Equal(t, []byte("hi"), in, "input must be left untouched")
	}
}
