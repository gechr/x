package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsAlpha(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":       false,
		"abc":    true,
		"ABC":    true,
		"aBcD":   true,
		"abc123": false,
		"ab c":   false,
		"ab-c":   false,
		"café":   false, // non-ASCII letter
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsAlpha(in), "IsAlpha(%q)", in)
	}
}

func TestIsAlphaChar(t *testing.T) {
	t.Parallel()

	cases := map[rune]bool{
		'a': true,
		'z': true,
		'A': true,
		'Z': true,
		'0': false,
		'-': false,
		' ': false,
		'é': false,
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsAlphaChar(in), "IsAlphaChar(%q)", in)
	}
}

func TestIsAlphanumeric(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":       false,
		"abc":    true,
		"ABC123": true,
		"0":      true,
		"a1b2":   true,
		"a b":    false,
		"a-b":    false,
		"a_b":    false,
		"café":   false, // non-ASCII letter
		"١٢٣":    false, // non-ASCII digits
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsAlphanumeric(in), "IsAlphanumeric(%q)", in)
	}
}

func TestIsAlphanumericChar(t *testing.T) {
	t.Parallel()

	cases := map[rune]bool{
		'a': true,
		'Z': true,
		'0': true,
		'9': true,
		'-': false,
		'_': false,
		' ': false,
		'é': false,
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsAlphanumericChar(in), "IsAlphanumericChar(%q)", in)
	}
}

func TestIsASCII(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":                false,
		"abc":             true,
		"a1!~":            true,
		"whitespace \t\n": true,
		"café":            false,
		"١٢٣":             false,
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsASCII(in), "IsASCII(%q)", in)
	}
}
