package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsHex(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.IsHex("deadbeef"))
	require.True(t, xstrings.IsHex("DEADBEEF"))
	require.True(t, xstrings.IsHex("0123456789abcdefABCDEF"))
	require.True(t, xstrings.IsHex(""))
	require.False(t, xstrings.IsHex("xyz"))
	require.False(t, xstrings.IsHex("deadbeeg"))
	require.False(t, xstrings.IsHex("dead beef"))
}

func TestIsHexChar(t *testing.T) {
	t.Parallel()

	for _, c := range "0123456789abcdefABCDEF" {
		require.True(t, xstrings.IsHexChar(c))
	}
	for _, c := range "ghijklmnopqrstuvwxyzGHIJKLMNOP !@#" {
		require.False(t, xstrings.IsHexChar(c))
	}
}
