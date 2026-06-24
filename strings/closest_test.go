package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestClosest(t *testing.T) {
	t.Parallel()

	keys := []string{"verify", "prerelease", "downgrade", "deep", "output"}
	cases := []struct {
		name       string
		target     string
		candidates []string
		want       string
	}{
		{name: "exact match", target: "verify", candidates: keys, want: "verify"},
		{
			name:       "adjacent transposition counts as one edit",
			target:     "verfiy",
			candidates: keys,
			want:       "verify",
		},
		{name: "single deletion", target: "verfy", candidates: keys, want: "verify"},
		{name: "single insertion", target: "outout", candidates: keys, want: "output"},
		{name: "too far is unsuggested", target: "xyzzy", candidates: keys, want: ""},
		{name: "empty target suggests nothing", target: "", candidates: keys, want: ""},
		{name: "no candidates", target: "verify", candidates: nil, want: ""},
		{name: "short target tolerates one edit", target: "dep", candidates: keys, want: "deep"},
		{
			name:       "first of equidistant ties wins",
			target:     "x",
			candidates: []string{"a", "b"},
			want:       "a",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.want, xstrings.Closest(tc.target, tc.candidates))
		})
	}
}
