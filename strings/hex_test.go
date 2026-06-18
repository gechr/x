package strings_test

import (
	"strings"
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsHex(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.IsHex("deadbeef"))
	require.True(t, xstrings.IsHex("DEADBEEF"))
	require.True(t, xstrings.IsHex("0123456789abcdefABCDEF"))
	require.False(t, xstrings.IsHex(""), "empty is not hex")
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

func TestIsGitCommit(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.IsGitCommit("a0dfaeb072753c3d48cd4df5fdacfd035b2281bf"))
	require.False(t, xstrings.IsGitCommit("a0dfaeb"), "abbreviated")
	require.False(t, xstrings.IsGitCommit(strings.Repeat("a", 64)), "that is sha256 length")
	require.False(t, xstrings.IsGitCommit(strings.Repeat("g", 40)), "not hex")
}

func TestIsSHA256(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.IsSHA256(strings.Repeat("a", 64)))
	require.False(t, xstrings.IsSHA256(strings.Repeat("a", 40)), "that is git commit length")
	require.False(t, xstrings.IsSHA256(strings.Repeat("z", 64)), "not hex")
}
