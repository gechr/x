package human_test

import (
	"testing"

	"github.com/gechr/x/human"
	"github.com/stretchr/testify/require"
)

func TestPlural(t *testing.T) {
	t.Parallel()

	cases := []struct {
		n        int
		singular string
		plural   string
		want     string
	}{
		{0, "file", "files", "files"},
		{1, "file", "files", "file"},
		{2, "file", "files", "files"},
		{42, "match", "matches", "matches"},
	}
	for _, tc := range cases {
		require.Equal(t, tc.want, human.Plural(tc.n, tc.singular, tc.plural))
	}
}

func TestPluralize(t *testing.T) {
	t.Parallel()

	cases := []struct {
		n        int
		singular string
		plural   string
		want     string
	}{
		{0, "file", "files", "0 files"},
		{1, "file", "files", "1 file"},
		{2, "file", "files", "2 files"},
		{42, "match", "matches", "42 matches"},
	}
	for _, tc := range cases {
		require.Equal(t, tc.want, human.Pluralize(tc.n, tc.singular, tc.plural))
	}
}

func TestFormatNumber(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   int64
		sep  string
		want string
	}{
		{0, ",", "0"},
		{7, ",", "7"},
		{42, ",", "42"},
		{999, ",", "999"},
		{1000, ",", "1,000"},
		{1234, ",", "1,234"},
		{12345, ",", "12,345"},
		{123456, ",", "123,456"},
		{1234567, ",", "1,234,567"},
		{-1, ",", "-1"},
		{-1000, ",", "-1,000"},
		{-1234567, ",", "-1,234,567"},
		{1234567, ".", "1.234.567"},
		{1234567, " ", "1 234 567"},
		{1234567, "_", "1_234_567"},
		{1234567, "", "1234567"},
	}
	for _, tc := range cases {
		require.Equal(
			t,
			tc.want,
			human.FormatNumber(tc.in, tc.sep),
			"FormatNumber(%d, %q)",
			tc.in,
			tc.sep,
		)
	}
}

func TestFormatOrdinal(t *testing.T) {
	t.Parallel()

	cases := map[int]string{
		0:   "0th",
		1:   "1st",
		2:   "2nd",
		3:   "3rd",
		4:   "4th",
		10:  "10th",
		11:  "11th",
		12:  "12th",
		13:  "13th",
		14:  "14th",
		20:  "20th",
		21:  "21st",
		22:  "22nd",
		23:  "23rd",
		101: "101st",
		111: "111th",
		112: "112th",
		113: "113th",
		121: "121st",
		-1:  "-1st",
	}
	for in, want := range cases {
		require.Equal(t, want, human.FormatOrdinal(in), "FormatOrdinal(%d)", in)
	}
}
