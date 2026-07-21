package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestSplitBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		sep  string
		want [][]byte
	}{
		{name: "empty", in: "", sep: "\n", want: [][]byte{}},
		{
			name: "lines",
			in:   " alpha \n\n bravo \n",
			sep:  "\n",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
		{
			name: "csv",
			in:   " alpha, ,bravo,",
			sep:  ",",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
		{
			name: "multi character separator",
			in:   "alpha :: bravo :: ",
			sep:  "::",
			want: [][]byte{[]byte("alpha"), []byte("bravo")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xbytes.SplitBy([]byte(tt.in), []byte(tt.sep)))
		})
	}
}
