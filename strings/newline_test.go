package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
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
		require.Equal(t, want, xstrings.EnsureTrailingNewline(in), "EnsureTrailingNewline(%q)", in)
	}
}
