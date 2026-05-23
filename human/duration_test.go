package human_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/gechr/x/human"
	"github.com/stretchr/testify/require"
)

func TestFormatDuration(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   time.Duration
		want string
	}{
		{0, "0s"},
		{500 * time.Nanosecond, "500ns"},
		{2 * time.Microsecond, "2µs"},
		{50 * time.Millisecond, "50ms"},
		{1 * time.Second, "1s"},
		{59 * time.Second, "59s"},
		{time.Minute, "1m"},
		{90 * time.Second, "1m30s"},
		{2*time.Hour + 15*time.Minute, "2h15m"},
		{3 * time.Hour, "3h"},
		{25 * time.Hour, "1d1h"},
		{72 * time.Hour, "3d"},
		{7 * 24 * time.Hour, "1w"},
		{8 * 24 * time.Hour, "1w1d"},
		{14 * 24 * time.Hour, "2w"},
		{2*7*24*time.Hour + 3*24*time.Hour, "2w3d"},
		{365 * 24 * time.Hour, "1y"},
		{400 * 24 * time.Hour, "1y5w"},
		{2*365*24*time.Hour + 30*7*24*time.Hour, "2y30w"},
		{-90 * time.Second, "-1m30s"},
	}
	for _, tc := range cases {
		require.Equal(t, tc.want, human.FormatDuration(tc.in), "FormatDuration(%s)", tc.in)
	}
}

func TestFormatDuration_MinInt64NoOverflow(t *testing.T) {
	t.Parallel()

	got := human.FormatDuration(time.Duration(math.MinInt64))
	require.True(t, strings.HasPrefix(got, "-"), "expected leading -, got %q", got)
	require.True(
		t,
		strings.HasSuffix(got, "y") || strings.Contains(got, "y"),
		"expected years unit, got %q",
		got,
	)
}
