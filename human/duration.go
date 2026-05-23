package human

import (
	"math"
	"strconv"
	"strings"
	"time"
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

	const (
		day  = HoursPerDay * time.Hour
		week = DaysPerWeek * day
		year = DaysPerYear * day
	)

	if d >= time.Second {
		d = d.Round(time.Second)
	}

	var parts []string
	switch {
	case d >= year:
		years := int(d / year)
		d -= time.Duration(years) * year
		parts = append(parts, strconv.Itoa(years)+"y")
		if weeks := int(d / week); weeks > 0 {
			parts = append(parts, strconv.Itoa(weeks)+"w")
		}
	case d >= week:
		weeks := int(d / week)
		d -= time.Duration(weeks) * week
		parts = append(parts, strconv.Itoa(weeks)+"w")
		if days := int(d / day); days > 0 {
			parts = append(parts, strconv.Itoa(days)+"d")
		}
	case d >= day:
		days := int(d / day)
		d -= time.Duration(days) * day
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
