package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestSplitLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want []string
	}{
		{name: "empty", in: "", want: []string{}},
		{name: "single line", in: "alpha", want: []string{"alpha"}},
		{name: "lf", in: "alpha\nbravo", want: []string{"alpha", "bravo"}},
		{name: "crlf", in: "alpha\r\nbravo", want: []string{"alpha", "bravo"}},
		{name: "trims input", in: " \n alpha \n ", want: []string{"alpha"}},
		{name: "trailing lf", in: "alpha\n", want: []string{"alpha"}},
		{name: "trailing crlf", in: "alpha\r\n", want: []string{"alpha"}},
		{name: "blank line", in: "alpha\n\nbravo", want: []string{"alpha", "bravo"}},
		{name: "space only line", in: "alpha\n \t \nbravo", want: []string{"alpha", "bravo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xstrings.SplitLines(tt.in))
		})
	}
}
