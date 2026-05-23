package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestTruncate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		n      int
		suffix string
		want   string
	}{
		{"shorter than limit", "hi", 8, "…", "hi"},
		{"exactly at limit", "hello!!!", 8, "…", "hello!!!"},
		{"truncated", "hello world", 8, "…", "hello w…"},
		{"truncated no suffix", "hello world", 5, "", "hello"},
		{"suffix longer than n", "hello world", 2, "...", ".."},
		{"zero n", "abc", 0, "…", ""},
		{"unicode", "café latte", 5, "…", "café…"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, xstrings.Truncate(tc.in, tc.n, tc.suffix))
		})
	}
}
