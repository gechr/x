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
		// 0x8000 is 127.502 of 255 and 0x7fff is 127.498, so these only
		// pass if components round to the nearest value rather than truncate.
		"four-digit components round to nearest": {
			response: "\x1b]11;rgb:8000/7fff/ffff\a",
			red:      128,
			green:    127,
			blue:     255,
		},
		"three-digit components round to nearest": {
			response: "\x1b]11;rgb:800/fff/000\a",
			red:      128,
			green:    255,
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
		"\x1b]11;rgb:ffff/ffff/ffff\x1b",
		"\x1b]11;rgb:ffff/ffff/ffff\x1b[0m",
		"\x1b]11;rgb:gggg/ffff/ffff\a",
		"\x1b]11;rgb:ffff/ffff\a",
		"\x1b]11;rgb:fffff/ffff/ffff\a",
	} {
		_, _, _, ok := parseBackgroundResponse(response)
		require.False(t, ok, response)
	}
}

func TestQueryResponse_DeviceAttributes(t *testing.T) {
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
		"aborted introducer": {
			response: "\x1b[?\x1b[0mc",
		},
		"intermediate byte": {
			response: "\x1b[?62 qc",
		},
		"other CSI final": {
			response: "\x1b[?25hc",
		},
		"color after aborted introducer": {
			response: "\x1b[?\x1b]11;rgb:9abc/9abc/9abc\a",
		},
		"secondary attributes": {
			response: "\x1b[>0;1c",
		},
		"tertiary attributes": {
			response: "\x1b[=0;1c",
		},
		"cancelled attributes": {
			response: "\x1b[?62\x18c",
		},
		"substituted attributes": {
			response: "\x1b[?62\x1ac",
		},
		"recovered attributes": {
			response: "\x1b[?62\x1b[?1;2c",
			complete: true,
		},
		"attributes after control string": {
			response: "\x1bPignored[?62c\x1b\\\x1b[?1;2c",
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
			var parser queryResponse
			var complete bool
			for i := range len(test.response) {
				complete = parser.consume(test.response[i])
				if i < len(test.response)-1 {
					require.False(t, complete)
				}
			}
			require.Equal(t, test.complete, complete)
		})
	}
}
