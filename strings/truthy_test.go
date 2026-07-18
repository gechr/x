package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsTruthy(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":       false,
		" 1 ":    true, //nolint:gocritic // whitespace is the point: trimming is under test
		"0":      false,
		"1":      true,
		"On":     true,
		"TRUE":   true,
		"Yes":    true,
		"banana": false,
		"false":  false,
		"on":     true,
		"true":   true,
		"yes":    true,
	}
	for in, want := range cases {
		t.Run(in, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, want, xstrings.IsTruthy(in))
		})
	}
}

func TestIsFalsy(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":       false,
		" 0 ":    true, //nolint:gocritic // whitespace is the point: trimming is under test
		"0":      true,
		"1":      false,
		"FALSE":  true,
		"No":     true,
		"Off":    true,
		"banana": false,
		"false":  true,
		"no":     true,
		"off":    true,
		"true":   false,
	}
	for in, want := range cases {
		t.Run(in, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, want, xstrings.IsFalsy(in))
		})
	}
}
