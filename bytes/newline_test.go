package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestEnsureTrailingNewline(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"":              "\n",
		"\n":            "\n",
		"\n\n\n":        "\n",
		"hello":         "hello\n",
		"hello\n":       "hello\n",
		"hello\n\n\n":   "hello\n",
		"a\nb":          "a\nb\n",
		"a\nb\n\n":      "a\nb\n",
		"trailing  \t ": "trailing  \t \n", //nolint:gocritic // trailing whitespace is the case under test
	}
	for in, want := range cases {
		require.Equal(
			t,
			want,
			string(xbytes.EnsureTrailingNewline([]byte(in))),
			"EnsureTrailingNewline(%q)",
			in,
		)
	}
}

// EnsureTrailingNewline must not clobber the caller's backing array when the
// input has trailing newlines to trim (the append-aliasing hazard).
func TestEnsureTrailingNewlineNoAlias(t *testing.T) {
	t.Parallel()

	in := []byte("hi\n\n")
	_ = xbytes.EnsureTrailingNewline(in)
	require.Equal(t, []byte("hi\n\n"), in, "input must be left untouched")
}
