package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
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
		require.Equal(t, want, xbytes.IsDigits([]byte(in)), "IsDigits(%q)", in)
	}
}

func TestIsDigitChar(t *testing.T) {
	t.Parallel()

	cases := map[byte]bool{
		'0': true,
		'9': true,
		'a': false,
		'-': false,
		' ': false,
	}
	for in, want := range cases {
		require.Equal(t, want, xbytes.IsDigitChar(in), "IsDigitChar(%q)", in)
	}
}
