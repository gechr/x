package bytes_test

import (
	"bytes"
	"crypto/sha256"
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestHexEqual(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.HexEqual([]byte("0x1234"), []byte("0x1234")), "equal with prefix")
	require.True(t, xbytes.HexEqual([]byte("1234"), []byte("1234")), "equal without prefix")
	require.True(
		t,
		xbytes.HexEqual([]byte("0x1234aBcD"), []byte("1234abcd")),
		"prefix and case differ",
	)
	require.True(t, xbytes.HexEqual([]byte("  0xDEAD "), []byte("dead")), "surrounding whitespace")
	require.True(t, xbytes.HexEqual([]byte(""), []byte("")), "both blank")
	require.False(t, xbytes.HexEqual([]byte("0x1234"), []byte("0x5678")), "different values")
	require.False(t, xbytes.HexEqual([]byte("0x1234"), []byte("0x123456")), "different lengths")
	require.False(t, xbytes.HexEqual([]byte(""), []byte("0x1234")), "blank vs non-blank")
}

func TestIsHex(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.IsHex([]byte("deadbeef")))
	require.True(t, xbytes.IsHex([]byte("DEADBEEF")))
	require.True(t, xbytes.IsHex([]byte("0123456789abcdefABCDEF")))
	require.False(t, xbytes.IsHex([]byte("")), "empty is not hex")
	require.False(t, xbytes.IsHex([]byte("xyz")))
	require.False(t, xbytes.IsHex([]byte("deadbeeg")))
	require.False(t, xbytes.IsHex([]byte("dead beef")))
}

func TestIsHexChar(t *testing.T) {
	t.Parallel()

	for _, c := range []byte("0123456789abcdefABCDEF") {
		require.True(t, xbytes.IsHexChar(c))
	}
	for _, c := range []byte("ghijklmnopqrstuvwxyzGHIJKLMNOP !@#") {
		require.False(t, xbytes.IsHexChar(c))
	}
}

func TestIsGitCommit(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.IsGitCommit([]byte("a0dfaeb072753c3d48cd4df5fdacfd035b2281bf")))
	require.False(t, xbytes.IsGitCommit([]byte("a0dfaeb")), "abbreviated")
	require.False(t, xbytes.IsGitCommit(bytes.Repeat([]byte("a"), 64)), "that is sha256 length")
	require.False(t, xbytes.IsGitCommit(bytes.Repeat([]byte("g"), 40)), "not hex")
}

func TestIsSHA256(t *testing.T) {
	t.Parallel()

	require.True(t, xbytes.IsSHA256(bytes.Repeat([]byte("a"), 64)))
	require.False(t, xbytes.IsSHA256(bytes.Repeat([]byte("a"), 40)), "that is git commit length")
	require.False(t, xbytes.IsSHA256(bytes.Repeat([]byte("z"), 64)), "not hex")
}

func TestDecodeSHA256(t *testing.T) {
	t.Parallel()

	want := sha256.Sum256(nil)
	for _, input := range []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
	} {
		got, err := xbytes.DecodeSHA256([]byte(input))
		require.NoError(t, err)
		require.Equal(t, want, got)
	}

	for _, input := range []string{
		"deadbeef",
		string(bytes.Repeat([]byte("z"), 64)),
	} {
		got, err := xbytes.DecodeSHA256([]byte(input))
		require.Error(t, err)
		require.Zero(t, got)
	}
}
