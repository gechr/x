package human

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	xtime "github.com/gechr/x/time"
)

// FormatDuration formats d as up to two adjacent units with no separator
// (e.g. "2h15m", "1w2d", "1y5w"). Years are 365 days, weeks are 7 days.
// Durations >= 1s are rounded to the nearest second.
//
//	FormatDuration(90 * time.Second)             // "1m30s"
//	FormatDuration(2*time.Hour + 15*time.Minute) // "2h15m"
//	FormatDuration(8 * 24 * time.Hour)           // "1w1d"
//	FormatDuration(400 * 24 * time.Hour)         // "1y5w"
//	FormatDuration(50 * time.Millisecond)        // "50ms"
func FormatDuration(d time.Duration) string {
	if d == 0 {
		return "0s"
	}

	neg := d < 0
	if neg {
		if d == math.MinInt64 {
			d = math.MaxInt64
		} else {
			d = -d
		}
	}

	if d >= time.Second {
		d = d.Round(time.Second)
	}

	var parts []string
	switch {
	case d >= xtime.Year:
		years := int(d / xtime.Year)
		d -= time.Duration(years) * xtime.Year
		parts = append(parts, strconv.Itoa(years)+"y")
		if weeks := int(d / xtime.Week); weeks > 0 {
			parts = append(parts, strconv.Itoa(weeks)+"w")
		}
	case d >= xtime.Week:
		weeks := int(d / xtime.Week)
		d -= time.Duration(weeks) * xtime.Week
		parts = append(parts, strconv.Itoa(weeks)+"w")
		if days := int(d / xtime.Day); days > 0 {
			parts = append(parts, strconv.Itoa(days)+"d")
		}
	case d >= xtime.Day:
		days := int(d / xtime.Day)
		d -= time.Duration(days) * xtime.Day
		parts = append(parts, strconv.Itoa(days)+"d")
		if hours := int(d / time.Hour); hours > 0 {
			parts = append(parts, strconv.Itoa(hours)+"h")
		}
	case d >= time.Hour:
		hours := int(d / time.Hour)
		d -= time.Duration(hours) * time.Hour
		parts = append(parts, strconv.Itoa(hours)+"h")
		if mins := int(d / time.Minute); mins > 0 {
			parts = append(parts, strconv.Itoa(mins)+"m")
		}
	case d >= time.Minute:
		mins := int(d / time.Minute)
		d -= time.Duration(mins) * time.Minute
		parts = append(parts, strconv.Itoa(mins)+"m")
		if secs := int(d / time.Second); secs > 0 {
			parts = append(parts, strconv.Itoa(secs)+"s")
		}
	case d >= time.Second:
		parts = append(parts, strconv.Itoa(int(d/time.Second))+"s")
	case d >= time.Millisecond:
		parts = append(parts, strconv.Itoa(int(d/time.Millisecond))+"ms")
	case d >= time.Microsecond:
		parts = append(parts, strconv.Itoa(int(d/time.Microsecond))+"µs")
	default:
		parts = append(parts, strconv.FormatInt(int64(d), 10)+"ns")
	}

	out := strings.Join(parts, "")
	if neg {
		return "-" + out
	}
	return out
}

// ParseDuration parses a human duration string into a time.Duration. It is the
// inverse of FormatDuration, accepting the units that function emits: y, w, d,
// h, m, s, ms, µs (or us), and ns, where a year is 365 days and a week is 7
// days. Units may be combined but each may appear at most once and must run in
// descending order of size, so "1y2w", "2h15m", and "90s" are valid while a
// repeated ("5w5w") or out-of-order ("1w1y") unit is an error. An optional
// leading - negates the result, and "0" parses to zero.
//
//	ParseDuration("2h15m")  // 2*time.Hour + 15*time.Minute
//	ParseDuration("1w2d")   // 9 * 24 * time.Hour
//	ParseDuration("-1m30s") // -90 * time.Second
func ParseDuration(s string) (time.Duration, error) {
	orig := s
	if s == "" {
		return 0, errors.New("invalid duration: empty string")
	}

	neg := false
	switch s[0] {
	case '-':
		neg, s = true, s[1:]
	case '+':
		s = s[1:]
	}
	if s == "" {
		return 0, fmt.Errorf("invalid duration %q: no value after sign", orig)
	}
	if s == "0" {
		return 0, nil
	}

	rs := []rune(s)
	pos, end := 0, len(rs)
	largest := time.Duration(math.MaxInt64)

	var total time.Duration
	for pos < end {
		numStart := pos
		for pos < end && rs[pos] >= '0' && rs[pos] <= '9' {
			pos++
		}
		if pos == numStart {
			return 0, fmt.Errorf("invalid duration %q: expected a number", orig)
		}
		n, err := strconv.Atoi(string(rs[numStart:pos]))
		if err != nil {
			return 0, fmt.Errorf("invalid duration %q: %w", orig, err)
		}

		unitStart := pos
		for pos < end && (rs[pos] < '0' || rs[pos] > '9') {
			pos++
		}
		unit := string(rs[unitStart:pos])
		if unit == "" {
			return 0, fmt.Errorf("invalid duration %q: missing unit", orig)
		}

		size, ok := unitSize(unit)
		if !ok {
			return 0, fmt.Errorf("invalid duration %q: unknown unit %q", orig, unit)
		}
		if size >= largest {
			return 0, fmt.Errorf(
				"invalid duration %q: unit %q is repeated or out of order",
				orig,
				unit,
			)
		}
		largest = size
		total += time.Duration(n) * size
	}

	if neg {
		return -total, nil
	}
	return total, nil
}

// unitSize maps a duration unit to its size. Because every unit has a distinct
// size, callers enforce descending order (and reject repeats) by requiring each
// successive size to be strictly smaller than the last.
func unitSize(unit string) (time.Duration, bool) {
	switch unit {
	case "y":
		return xtime.Year, true
	case "w":
		return xtime.Week, true
	case "d":
		return xtime.Day, true
	case "h":
		return time.Hour, true
	case "m":
		return time.Minute, true
	case "s":
		return time.Second, true
	case "ms":
		return time.Millisecond, true
	case "us", "µs", "μs":
		return time.Microsecond, true
	case "ns":
		return time.Nanosecond, true
	}
	return 0, false
}
