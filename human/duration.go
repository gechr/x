package human

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	xstrings "github.com/gechr/x/strings"
	xtime "github.com/gechr/x/time"
)

const (
	defaultDurationDecimalMaximum = 10 * time.Second
	defaultDurationDecimalRound   = 100 * time.Millisecond
)

// DurationFormatOptions controls duration rounding and fractional display.
type DurationFormatOptions struct {
	// Precision is the number of decimal places used for durations below one
	// minute.
	Precision int
	// Round is the rounding granularity. Zero disables pre-format rounding.
	Round time.Duration
	// TrimTrailingZeros removes trailing zeroes from the fractional part.
	TrimTrailingZeros bool
}

// FormatDuration formats d as up to two adjacent units with no separator
// (e.g. "5.2s", "2h15m", "1w2d"). Years are 365 days, weeks are 7 days.
// With no options it uses magnitude-based resolution: the largest exact
// sub-second unit, one decimal place below ten seconds, and whole seconds
// thereafter. When multiple options values are supplied, the last applies.
//
//	FormatDuration(90 * time.Second)             // "1m30s"
//	FormatDuration(2*time.Hour + 15*time.Minute) // "2h15m"
//	FormatDuration(8 * 24 * time.Hour)           // "1w1d"
//	FormatDuration(400 * 24 * time.Hour)         // "1y5w"
//	FormatDuration(50 * time.Millisecond)        // "50ms"
//	FormatDuration(5200 * time.Millisecond)      // "5.2s"
func FormatDuration(d time.Duration, options ...DurationFormatOptions) string {
	if d == 0 {
		return "0s"
	}

	opts := defaultDurationFormatOptions(d)
	if len(options) > 0 {
		opts = options[len(options)-1]
	}

	neg := d < 0
	if neg {
		if d == math.MinInt64 {
			d = math.MaxInt64
		} else {
			d = -d
		}
	}

	if opts.Round > 0 {
		d = d.Round(opts.Round)
	}
	if d == 0 {
		return "0s"
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
		parts = append(parts, formatDurationUnit(d, time.Second, "s", opts))
	case d >= time.Millisecond:
		parts = append(parts, formatDurationUnit(d, time.Millisecond, "ms", opts))
	case d >= time.Microsecond:
		parts = append(parts, formatDurationUnit(d, time.Microsecond, "µs", opts))
	default:
		parts = append(parts, formatDurationUnit(d, time.Nanosecond, "ns", opts))
	}

	out := strings.Join(parts, "")
	if neg {
		return "-" + out
	}
	return out
}

// defaultDurationFormatOptions selects the standard display scale by absolute
// magnitude.
func defaultDurationFormatOptions(d time.Duration) DurationFormatOptions {
	if d < 0 {
		if d == math.MinInt64 {
			d = math.MaxInt64
		} else {
			d = -d
		}
	}

	switch {
	case d < time.Microsecond:
		return DurationFormatOptions{Round: time.Nanosecond}
	case d < time.Millisecond:
		return DurationFormatOptions{Round: time.Microsecond}
	case d < time.Second:
		return DurationFormatOptions{Round: time.Millisecond}
	case d < defaultDurationDecimalMaximum:
		return DurationFormatOptions{
			Precision:         1,
			Round:             defaultDurationDecimalRound,
			TrimTrailingZeros: true,
		}
	default:
		return DurationFormatOptions{Round: time.Second}
	}
}

// formatDurationUnit formats d in a single unit.
func formatDurationUnit(
	d, unit time.Duration,
	suffix string,
	opts DurationFormatOptions,
) string {
	precision := max(opts.Precision, 0)
	value := float64(d) / float64(unit)
	formatted := strconv.FormatFloat(value, 'f', precision, 64)
	if opts.TrimTrailingZeros && strings.Contains(formatted, ".") {
		formatted = strings.TrimRight(formatted, "0")
		formatted = strings.TrimSuffix(formatted, ".")
	}
	return formatted + suffix
}

// ParseDuration parses a human duration string into a [time.Duration]. It is the
// inverse of [FormatDuration], accepting decimal values and the units y, w, d,
// h, m, s, ms, µs (or us), and ns, where a year is 365 days and a week is 7
// days. Units may be combined but each may appear at most once and must run in
// descending order of size. A decimal value must be the final component. An
// optional leading - negates the result, and `0` parses to zero.
//
//	ParseDuration("2h15m")  // 2*time.Hour + 15*time.Minute
//	ParseDuration("1w2d")   // 9 * 24 * time.Hour
//	ParseDuration("5.2s")   // 5200 * time.Millisecond
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
		for pos < end && xstrings.IsDigitChar(rs[pos]) {
			pos++
		}
		if pos == numStart {
			return 0, fmt.Errorf("invalid duration %q: expected a number", orig)
		}
		if pos < end && rs[pos] == '.' {
			pos++
			fractionStart := pos
			for pos < end && xstrings.IsDigitChar(rs[pos]) {
				pos++
			}
			if pos == fractionStart {
				return 0, fmt.Errorf(
					"invalid duration %q: expected digits after decimal point",
					orig,
				)
			}
		}
		number := string(rs[numStart:pos])

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

		part, err := parseDurationPart(number, size)
		if err != nil {
			return 0, fmt.Errorf("invalid duration %q: %w", orig, err)
		}
		if part > time.Duration(math.MaxInt64)-total {
			return 0, fmt.Errorf("invalid duration %q: value overflows time.Duration", orig)
		}
		total += part
		if strings.Contains(number, ".") && pos < end {
			return 0, fmt.Errorf("invalid duration %q: decimal value must be final", orig)
		}
	}

	if neg {
		return -total, nil
	}
	return total, nil
}

// parseDurationPart converts an exact decimal number of units to nanoseconds.
func parseDurationPart(number string, unit time.Duration) (time.Duration, error) {
	value, ok := new(big.Rat).SetString(number)
	if !ok {
		return 0, fmt.Errorf("invalid number %q", number)
	}
	value.Mul(value, new(big.Rat).SetInt64(int64(unit)))
	if !value.IsInt() {
		return 0, fmt.Errorf("value %q is more precise than a nanosecond", number)
	}
	if !value.Num().IsInt64() {
		return 0, fmt.Errorf("value %q overflows time.Duration", number)
	}
	return time.Duration(value.Num().Int64()), nil
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
