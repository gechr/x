package shell_test

import (
	"strings"
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func TestSplit(t *testing.T) {
	t.Parallel()

	input := "one two \"three four\" \"five \\\"six\\\"\" seven#eight # nine # ten\n eleven 'twelve\\' thirteen=13 fourteen/14"

	got, err := shell.Split(input)
	require.NoError(t, err)
	require.Equal(t, []string{
		"one",
		"two",
		"three four",
		"five \"six\"",
		"seven#eight",
		"eleven",
		"twelve\\",
		"thirteen=13",
		"fourteen/14",
	}, got)
}

func TestSplitCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want []string
	}{
		{name: "empty", in: "", want: nil},
		{name: "spaces", in: " \t\r\n ", want: nil},
		{name: "escaped space", in: `one\ two`, want: []string{"one two"}},
		{name: "adjacent quotes", in: `"one"'two' three`, want: []string{"onetwo", "three"}},
		{name: "empty quoted word", in: `"" ''`, want: []string{"", ""}},
		{name: "comment at token boundary", in: "one # two\nthree", want: []string{"one", "three"}},
		{name: "hash inside word", in: "one#two three", want: []string{"one#two", "three"}},
		{name: "double quoted escaped backslash", in: `"one\\two"`, want: []string{`one\two`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := shell.Split(tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSplitErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		in      string
		wantErr string
	}{
		{name: "trailing escape", in: `one\`, wantErr: "EOF found after escape character"},
		{
			name:    "trailing double quote escape",
			in:      `"one\`,
			wantErr: "EOF found after escape character",
		},
		{
			name:    "unterminated single quote",
			in:      `'one`,
			wantErr: "EOF found when expecting closing quote",
		},
		{
			name:    "unterminated double quote",
			in:      `"one`,
			wantErr: "EOF found when expecting closing quote",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := shell.Split(tt.in)
			require.Equal(t, []string(nil), got)
			require.EqualError(t, err, tt.wantErr)
		})
	}
}

func TestQuoteSplitRoundTrip(t *testing.T) {
	t.Parallel()

	want := []string{
		"",
		"plain",
		"two words",
		`"double"`,
		`'single'`,
		";${}",
		"`echo hello`",
		"line\nbreak",
	}
	quoted := make([]string, len(want))
	for i, word := range want {
		quoted[i] = shell.Quote(word)
	}

	got, err := shell.Split(strings.Join(quoted, " "))
	require.NoError(t, err)
	require.Equal(t, want, got)
}
