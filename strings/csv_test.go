package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestSplitCSV(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want []string
	}{
		{name: "empty", in: "", want: []string{}},
		{name: "single", in: "alpha", want: []string{"alpha"}},
		{name: "multiple", in: "alpha,bravo,charlie", want: []string{"alpha", "bravo", "charlie"}},
		{
			name: "trims spaces",
			in:   " alpha, bravo ,charlie ",
			want: []string{"alpha", "bravo", "charlie"},
		},
		{name: "drops empty", in: "alpha,, ,bravo,", want: []string{"alpha", "bravo"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xstrings.SplitCSV(tt.in))
		})
	}
}

func TestAppendCSV(t *testing.T) {
	t.Parallel()

	got := xstrings.AppendCSV([]string{"alpha"}, " bravo, ,charlie ")
	require.Equal(t, []string{"alpha", "bravo", "charlie"}, got)
}
