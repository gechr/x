package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestTruncateRight(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		n      int
		marker string
		want   string
	}{
		{"shorter than limit", "hi", 8, "…", "hi"},
		{"exactly at limit", "hello!!!", 8, "…", "hello!!!"},
		{"truncated", "hello world", 8, "…", "hello w…"},
		{"truncated no marker", "hello world", 5, "", "hello"},
		{"marker longer than n", "hello world", 2, "...", ".."},
		{"zero n", "abc", 0, "…", ""},
		{"unicode", "café latte", 5, "…", "café…"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(
				t,
				tc.want,
				string(xbytes.TruncateRight([]byte(tc.in), tc.n, []byte(tc.marker))),
			)
		})
	}
}

func TestTruncateLeft(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		n      int
		marker string
		want   string
	}{
		{"shorter than limit", "hi", 8, "…", "hi"},
		{"exactly at limit", "hello!!!", 8, "…", "hello!!!"},
		{"truncated", "hello world", 8, "…", "…o world"},
		{"truncated no marker", "hello world", 5, "", "world"},
		{"marker longer than n", "hello world", 2, "...", ".."},
		{"zero n", "abc", 0, "…", ""},
		{"unicode", "café latte", 5, "…", "…atte"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(
				t,
				tc.want,
				string(xbytes.TruncateLeft([]byte(tc.in), tc.n, []byte(tc.marker))),
			)
		})
	}
}

func TestTruncateMiddle(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		in     string
		n      int
		marker string
		want   string
	}{
		{"shorter than limit", "hi", 8, "…", "hi"},
		{"exactly at limit", "hello!!!", 8, "…", "hello!!!"},
		{"hash both ends", "0123456789abcdef", 7, "…", "012…def"},
		{"odd budget favours head", "0123456789abcdef", 6, "…", "012…ef"},
		{"truncated no marker", "hello world", 4, "", "held"},
		{"marker longer than n", "hello world", 2, "...", ".."},
		{"zero n", "abc", 0, "…", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(
				t,
				tc.want,
				string(xbytes.TruncateMiddle([]byte(tc.in), tc.n, []byte(tc.marker))),
			)
		})
	}
}
