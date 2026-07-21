package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestIndent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		prefix string
		want   string
	}{
		{"empty", "", "  ", ""},
		{"empty prefix no blanks", "foo\nbar", "", "foo\nbar"},
		{"empty prefix normalizes blanks", "foo\n   \nbar", "", "foo\n\nbar"},
		{"single line", "foo", "  ", "  foo"},
		{"multi line", "foo\nbar", "  ", "  foo\n  bar"},
		{"preserves blank lines", "foo\n\nbar", "> ", "> foo\n\n> bar"},
		{"normalizes whitespace-only lines", "foo\n   \nbar", "> ", "> foo\n\n> bar"},
		{"trailing newline", "foo\nbar\n", "  ", "  foo\n  bar\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, string(xbytes.Indent([]byte(tc.in), []byte(tc.prefix))))
		})
	}
}

func TestDedent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"no leading", "foo\nbar", "foo\nbar"},
		{"uniform", "  foo\n  bar", "foo\nbar"},
		{"uneven keeps relative", "    foo\n      bar\n    baz", "foo\n  bar\nbaz"},
		{"ignores blank lines for prefix", "  foo\n\n  bar", "foo\n\nbar"},
		{"mixed tabs and spaces fall back", "\tfoo\n  bar", "\tfoo\n  bar"},
		{"whitespace-only line normalized", "  foo\n   \n  bar", "foo\n\nbar"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, string(xbytes.Dedent([]byte(tc.in))))
		})
	}
}

func TestIndentDedentRoundTrip(t *testing.T) {
	t.Parallel()

	original := []byte("foo\n  bar\n\nbaz")
	require.Equal(t, original, xbytes.Dedent(xbytes.Indent(original, []byte("    "))))
}
