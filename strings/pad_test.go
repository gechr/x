package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestPadLeft(t *testing.T) {
	t.Parallel()

	require.Equal(t, "   hi", xstrings.PadLeft("hi", 5))
	require.Equal(t, "hello", xstrings.PadLeft("hello", 5))
	require.Equal(t, "hello!", xstrings.PadLeft("hello!", 5))
	require.Equal(t, "   ", xstrings.PadLeft("", 3))
	require.Equal(t, "hi", xstrings.PadLeft("hi", 0))
	require.Equal(t, "hi", xstrings.PadLeft("hi", -1))
	// Width counts runes, not bytes.
	require.Equal(t, "  héé", xstrings.PadLeft("héé", 5))
}

func TestPadRight(t *testing.T) {
	t.Parallel()

	require.Equal(t, "hi   ", xstrings.PadRight("hi", 5))
	require.Equal(t, "hello", xstrings.PadRight("hello", 5))
	require.Equal(t, "hello!", xstrings.PadRight("hello!", 5))
	require.Equal(t, "   ", xstrings.PadRight("", 3))
	require.Equal(t, "hi", xstrings.PadRight("hi", 0))
	require.Equal(t, "héé  ", xstrings.PadRight("héé", 5))
}

func TestPadCenter(t *testing.T) {
	t.Parallel()

	require.Equal(t, " hi  ", xstrings.PadCenter("hi", 5))
	require.Equal(t, " hi ", xstrings.PadCenter("hi", 4))
	require.Equal(t, "hello", xstrings.PadCenter("hello", 5))
	require.Equal(t, "hello!", xstrings.PadCenter("hello!", 5))
	require.Equal(t, "   ", xstrings.PadCenter("", 3))
	require.Equal(t, "hi", xstrings.PadCenter("hi", -1))
	require.Equal(t, " héé ", xstrings.PadCenter("héé", 5))
}
