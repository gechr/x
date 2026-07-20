package strings_test

import (
	"crypto/sha256"
	"strings"
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestHexEqual(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.HexEqual("0x1234", "0x1234"), "equal with prefix")
	require.True(t, xstrings.HexEqual("1234", "1234"), "equal without prefix")
	require.True(t, xstrings.HexEqual("0x1234aBcD", "1234abcd"), "prefix and case differ")
	require.True(t, xstrings.HexEqual("  0xDEAD ", "dead"), "surrounding whitespace")
	require.True(t, xstrings.HexEqual("", ""), "both blank")
	require.False(t, xstrings.HexEqual("0x1234", "0x5678"), "different values")
	require.False(t, xstrings.HexEqual("0x1234", "0x123456"), "different lengths")
	require.False(t, xstrings.HexEqual("", "0x1234"), "blank vs non-blank")
}

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

func TestDecodeSHA256(t *testing.T) {
	t.Parallel()

	want := sha256.Sum256(nil)
	for _, input := range []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
	} {
		got, err := xstrings.DecodeSHA256(input)
		require.NoError(t, err)
		require.Equal(t, want, got)
	}

	for _, input := range []string{
		"deadbeef",
		strings.Repeat("z", 64),
	} {
		got, err := xstrings.DecodeSHA256(input)
		require.Error(t, err)
		require.Zero(t, got)
	}
}
