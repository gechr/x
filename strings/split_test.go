package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestSplitBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		sep  string
		want []string
	}{
		{name: "empty", in: "", sep: "\n", want: []string{}},
		{name: "lines", in: " alpha \n\n bravo \n", sep: "\n", want: []string{"alpha", "bravo"}},
		{name: "csv", in: " alpha, ,bravo,", sep: ",", want: []string{"alpha", "bravo"}},
		{
			name: "multi character separator",
			in:   "alpha :: bravo :: ",
			sep:  "::",
			want: []string{"alpha", "bravo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xstrings.SplitBy(tt.in, tt.sep))
		})
	}
}
