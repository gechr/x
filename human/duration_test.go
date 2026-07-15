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
		{1200 * time.Millisecond, "1.2s"},
		{5200 * time.Millisecond, "5.2s"},
		{9940 * time.Millisecond, "9.9s"},
		{9960 * time.Millisecond, "10s"},
		{12400 * time.Millisecond, "12s"},
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

func TestFormatDurationOptions(t *testing.T) {
	t.Parallel()

	d := 5200 * time.Millisecond
	require.Equal(t, "5.20s", human.FormatDuration(d, human.DurationFormatOptions{Precision: 2}))
	require.Equal(t, "5.2s", human.FormatDuration(d, human.DurationFormatOptions{
		Precision:         2,
		TrimTrailingZeros: true,
	}))
	require.Equal(t, "5s", human.FormatDuration(d, human.DurationFormatOptions{Round: time.Second}))
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

func TestParseDuration(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   string
		want time.Duration
	}{
		{"0", 0},
		{"0s", 0},
		{"90s", 90 * time.Second},
		{"5.2s", 5200 * time.Millisecond},
		{"1m30.5s", 90*time.Second + 500*time.Millisecond},
		{"1.5h", 90 * time.Minute},
		{"1m30s", 90 * time.Second},
		{"2h15m", 2*time.Hour + 15*time.Minute},
		{"1h30m45s", time.Hour + 30*time.Minute + 45*time.Second},
		{"1w2d", 9 * 24 * time.Hour},
		{"1y", 365 * 24 * time.Hour},
		{"1y2w", 365*24*time.Hour + 2*7*24*time.Hour},
		{"50ms", 50 * time.Millisecond},
		{"2µs", 2 * time.Microsecond},
		{"2us", 2 * time.Microsecond},
		{"500ns", 500 * time.Nanosecond},
		{"1m5ms", time.Minute + 5*time.Millisecond},
		{"-1m30s", -90 * time.Second},
		{"+2h", 2 * time.Hour},
	}
	for _, tc := range cases {
		got, err := human.ParseDuration(tc.in)
		require.NoError(t, err, "ParseDuration(%q)", tc.in)
		require.Equal(t, tc.want, got, "ParseDuration(%q)", tc.in)
	}
}

func TestParseDurationErrors(t *testing.T) {
	t.Parallel()

	cases := []string{
		"",         // empty
		"-",        // sign only
		"5",        // missing unit
		"5x",       // unknown unit
		"5w5w",     // repeated unit
		"1w1y",     // out of order
		"1m1h",     // out of order
		"5ms5m",    // out of order across sub-second
		"abc",      // no number
		".5s",      // leading digit required
		"1.s",      // fractional digit required
		"1.5h30m",  // decimal must be final
		"0.0001ns", // finer than nanosecond precision
	}
	for _, in := range cases {
		_, err := human.ParseDuration(in)
		require.Error(t, err, "ParseDuration(%q)", in)
	}
}

func TestParseDurationRoundTrip(t *testing.T) {
	t.Parallel()

	// FormatDuration output parses back to an equal duration for values it
	// represents exactly (<= two units, second precision).
	cases := []time.Duration{
		90 * time.Second,
		2*time.Hour + 15*time.Minute,
		9 * 24 * time.Hour,
		400 * 24 * time.Hour,
		-90 * time.Second,
		50 * time.Millisecond,
		5200 * time.Millisecond,
	}
	for _, d := range cases {
		got, err := human.ParseDuration(human.FormatDuration(d))
		require.NoError(t, err, "round-trip %s", d)
		require.Equal(t, d, got, "round-trip %s", d)
	}
}
