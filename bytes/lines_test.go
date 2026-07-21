package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestSplitLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want [][]byte
	}{
		{name: "empty", in: "", want: [][]byte{}},
		{name: "single line", in: "alpha", want: [][]byte{[]byte("alpha")}},
		{name: "lf", in: "alpha\nbravo", want: [][]byte{[]byte("alpha"), []byte("bravo")}},
		{name: "crlf", in: "alpha\r\nbravo", want: [][]byte{[]byte("alpha"), []byte("bravo")}},
		{name: "trims input", in: " \n alpha \n ", want: [][]byte{[]byte("alpha")}},
		{name: "trailing lf", in: "alpha\n", want: [][]byte{[]byte("alpha")}},
		{name: "trailing crlf", in: "alpha\r\n", want: [][]byte{[]byte("alpha")}},
		{
			name: "blank line",
			in:   "alpha\n\nbravo",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
		{
			name: "space only line",
			in:   "alpha\n \t \nbravo",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xbytes.SplitLines([]byte(tt.in)))
		})
	}
}

func TestSplitLinesRaw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want [][]byte
	}{
		{name: "empty", in: "", want: [][]byte{nil}},
		{name: "single line", in: "alpha", want: [][]byte{[]byte("alpha")}},
		{name: "lf", in: "alpha\nbravo", want: [][]byte{[]byte("alpha"), []byte("bravo")}},
		{
			name: "crlf normalized",
			in:   "alpha\r\nbravo",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
		{
			name: "keeps whitespace",
			in:   " alpha \n bravo ",
			want: [][]byte{[]byte(" alpha "), []byte(" bravo ")},
		},
		{name: "keeps trailing empty", in: "alpha\n", want: [][]byte{[]byte("alpha"), {}}},
		{
			name: "keeps blank lines",
			in:   "alpha\n\nbravo",
			want: [][]byte{[]byte("alpha"), {}, []byte("bravo")},
		},
		{name: "bare cr survives", in: "alpha\rbravo", want: [][]byte{[]byte("alpha\rbravo")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xbytes.SplitLinesRaw([]byte(tt.in)))
		})
	}
}
