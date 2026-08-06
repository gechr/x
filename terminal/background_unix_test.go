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
		"trailing device attributes reply": {
			response: "\x1b]11;rgb:0000/0000/0000\x1b\\\x1b[?62;4c",
			red:      0,
			green:    0,
			blue:     0,
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

func TestDeviceAttributesComplete(t *testing.T) {
	tests := map[string]struct {
		response string
		complete bool
	}{
		"empty": {
			response: "",
		},
		"reply alone": {
			response: "\x1b[?62;4c",
			complete: true,
		},
		"reply after background response": {
			response: "\x1b]11;rgb:0000/0000/0000\x1b\\\x1b[?1;2c",
			complete: true,
		},
		"minimal reply": {
			response: "\x1b[?c",
			complete: true,
		},
		"introducer without final byte": {
			response: "\x1b[?62;4",
		},
		// The final byte must be looked for after the introducer: hex digits
		// in a background response would otherwise terminate the reply early.
		"c in background response only": {
			response: "\x1b]11;rgb:9abc/9abc/9abc\x1b\\",
		},
		"background response still arriving": {
			response: "\x1b]11;rgb:0000/0000/00",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.complete, deviceAttributesComplete(test.response))
		})
	}
}
