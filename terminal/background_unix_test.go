//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package terminal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBackgroundResponse(t *testing.T) {
	tests := map[string]struct {
		response string
		red      uint8
		green    uint8
		blue     uint8
	}{
		"bel terminator": {
			response: "\x1b]11;rgb:ffff/8080/0000\a",
			red:      255,
			green:    128,
			blue:     0,
		},
		"st terminator": {
			response: "noise\x1b]11;rgb:1234/5678/9abc\x1b\\",
			red:      18,
			green:    86,
			blue:     154,
		},
		"short components": {
			response: "\x1b]11;rgb:f/8/0\a",
			red:      255,
			green:    136,
			blue:     0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			red, green, blue, ok := parseBackgroundResponse(test.response)
			require.True(t, ok)
			require.Equal(t, test.red, red)
			require.Equal(t, test.green, green)
			require.Equal(t, test.blue, blue)
		})
	}
}

func TestParseBackgroundResponse_Invalid(t *testing.T) {
	for _, response := range []string{
		"",
		"\x1b]10;rgb:ffff/ffff/ffff\a",
		"\x1b]11;rgb:ffff/ffff/ffff",
		"\x1b]11;rgb:gggg/ffff/ffff\a",
		"\x1b]11;rgb:ffff/ffff\a",
		"\x1b]11;rgb:fffff/ffff/ffff\a",
	} {
		_, _, _, ok := parseBackgroundResponse(response)
		require.False(t, ok, response)
	}
}
