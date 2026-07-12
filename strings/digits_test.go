package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsDigits(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":       false,
		"0":      true,
		"012345": true,
		"12a34":  false,
		"-1":     false,
		"1.5":    false,
		" 1":     false, //nolint:gocritic // leading space is the case under test
		"١٢٣":    false, // non-ASCII digits
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsDigits(in), "IsDigits(%q)", in)
	}
}

func TestIsDigitChar(t *testing.T) {
	t.Parallel()

	cases := map[rune]bool{
		'0': true,
		'9': true,
		'a': false,
		'-': false,
		' ': false,
		'٢': false, // non-ASCII digit
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsDigitChar(in), "IsDigitChar(%q)", in)
	}
}
