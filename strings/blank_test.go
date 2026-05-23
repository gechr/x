package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsBlank(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":          true,
		" ":         true,
		"\t\n  \t":  true,
		"x":         false,
		"  hello  ": false,
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsBlank(in), "IsBlank(%q)", in)
	}
}
