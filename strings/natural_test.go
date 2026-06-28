package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestCompareNatural(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{name: "empty before nonempty", a: "", b: "a", want: -1},
		{name: "lexical letters", a: "a", b: "b", want: -1},
		{name: "prefix is shorter", a: "a", b: "aa", want: -1},
		{name: "digit by value", a: "a0", b: "a1", want: -1},
		{name: "fewer digits is smaller", a: "a0", b: "a00", want: -1},
		{name: "leading zero magnitude", a: "a00", b: "a01", want: -1},
		{name: "leading zero versus bare", a: "a01", b: "a1", want: -1},
		{name: "value beats lexical", a: "a01", b: "a2", want: -1},
		{name: "value beats lexical with suffix", a: "a01x", b: "a2x", want: -1},
		{name: "only the last number matters", a: "a0b00", b: "a00b1", want: -1},
		{name: "ten after two", a: "10", b: "2", want: 1},
		{name: "two-x after one-x", a: "a2x", b: "a01x", want: 1},
		{name: "bare versus leading zero", a: "a2", b: "a01", want: 1},
		{name: "equal letters", a: "a", b: "a", want: 0},
		{name: "equal padded", a: "a01", b: "a01", want: 0},
		{name: "padded equals bare across groups", a: "a00b00", b: "a0b00", want: 0},
		// Numbers far beyond uint64 still compare by value, not lexically.
		{
			name: "huge number by value",
			a:    "a99999999999999999999",
			b:    "a100000000000000000000",
			want: -1,
		},
		{
			name: "huge equal with leading zeros and suffix",
			a:    "a099999999999999999999x",
			b:    "a99999999999999999999x",
			want: 0,
		},
		{
			name: "huge padded equals bare with suffix",
			a:    "a00000000000000000000001x",
			b:    "a1x",
			want: 0,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.want, xstrings.CompareNatural(tc.a, tc.b))
			require.Equal(
				t,
				-tc.want,
				xstrings.CompareNatural(tc.b, tc.a),
				"reversing the operands negates the result",
			)
		})
	}
}
