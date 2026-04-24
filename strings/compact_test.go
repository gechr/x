package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestCompactLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		sep  string
		want string
	}{
		{name: "empty", in: "", sep: " | ", want: ""},
		{name: "single", in: "hello", sep: " | ", want: "hello"},
		{name: "trims whitespace", in: "  a  \n  b  ", sep: " | ", want: "a | b"},
		{name: "drops blank lines", in: "a\n\n\nb", sep: " | ", want: "a | b"},
		{name: "dedupes", in: "a\nb\na\nb\nc", sep: " | ", want: "a | b | c"},
		{name: "custom separator", in: "a\nb", sep: ", ", want: "a, b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.want, xstrings.CompactLines(tt.in, tt.sep))
		})
	}
}
