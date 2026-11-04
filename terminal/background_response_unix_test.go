//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package terminal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryResponse_Background(t *testing.T) {
	tests := map[string]struct {
		input string
		valid bool
	}{
		"BEL":           {"\x1b]11;rgb:1234/5678/9abc\a", true},
		"ST":            {"\x1b]11;rgb:1234/5678/9abc\x1b\\", true},
		"unterminated":  {"\x1b]11;rgb:1234/5678/9abc", false},
		"partial ST":    {"\x1b]11;rgb:1234/5678/9abc\x1b", false},
		"aborted":       {"\x1b]11;rgb:1234/5678/9abc\x1b[0m", false},
		"cancelled":     {"\x1b]11;rgb:1234/5678/9abc\x18\a", false},
		"foreground":    {"\x1b]10;rgb:1234/5678/9abc\a", false},
		"invalid color": {"\x1b]11;rgb:zzzz/5678/9abc\a", false},
		"oversized":     {"\x1b]11;" + strings.Repeat("x", maxBackgroundResponse) + "\a", false},
		"recovery": {
			"\x1b]" + strings.Repeat(
				"x",
				maxBackgroundResponse+1,
			) + "\a\x1b]11;rgb:1234/5678/9abc\a",
			true,
		},
		"leading noise": {
			strings.Repeat("x", maxBackgroundResponse*2) + "\x1b]11;rgb:1234/5678/9abc\a",
			true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var parser queryResponse
			for i := range len(test.input) {
				require.False(t, parser.consume(test.input[i]))
				if i < len(test.input)-1 {
					require.Empty(t, parser.background)
				}
			}
			red, green, blue, valid := parseBackgroundResponse(parser.background)
			require.Equal(t, test.valid, valid)
			if valid {
				require.Equal(t, uint8(18), red)
				require.Equal(t, uint8(86), green)
				require.Equal(t, uint8(154), blue)
			}
		})
	}
}

func TestQueryResponse_StopsAtDA1(t *testing.T) {
	for _, prefix := range []string{
		strings.Repeat("x", maxBackgroundResponse*2),
		"\x1b]" + strings.Repeat("x", maxBackgroundResponse*2) + "\a",
		"\x1b]" + strings.Repeat("x", maxBackgroundResponse*2) + "\x1b\\",
		"\x1b]11;rgb:9abc/9abc/9abc\a",
		"\x1bPignored[?62c\x1b\\",
		"\x1bXignored[?62c\x1b\\",
		"\x1b^ignored[?62c\x1b\\",
		"\x1b_ignored[?62c\x1b\\",
		"\x1bPunterminated",
		"\x1b]unterminated",
	} {
		var parser queryResponse
		replies := prefix + "\x1b[?62;4c"
		input := replies + "USER"
		consumed := 0
		for i := range len(input) {
			if parser.consume(input[i]) {
				consumed = i + 1
				break
			}
		}
		require.Equal(t, len(replies), consumed)
	}
}

func FuzzQueryResponse(f *testing.F) {
	f.Add([]byte("\x1b]11;rgb:0000/0000/0000\x1b\\\x1b[?62;4c"))
	f.Add([]byte("\x1b[?\x1b[0mc"))
	f.Add([]byte("\x1b]" + strings.Repeat("x", maxBackgroundResponse*2) + "\a"))
	f.Fuzz(func(t *testing.T, input []byte) {
		var parser queryResponse
		for _, b := range input {
			parser.consume(b)
			require.LessOrEqual(t, parser.length, maxBackgroundResponse)
			require.LessOrEqual(t, len(parser.background), maxBackgroundResponse+len("\x1b]\a"))
		}
	})
}
