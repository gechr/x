package human

import (
	"strconv"
	"strings"
)

const thousandsGroup = 3

// Plural returns `singular` when `n == 1`, otherwise `plural`. Unlike
// [Pluralize], it omits the count.
//
//	Plural(1, "file", "files") // "file"
//	Plural(3, "file", "files") // "files"
func Plural(n int, singular, plural string) string {
	if n == 1 {
		return singular
	}
	return plural
}

// Pluralize returns "1 singular" or "n plural".
//
//	Pluralize(1, "file", "files") // "1 file"
//	Pluralize(3, "file", "files") // "3 files"
func Pluralize(n int, singular, plural string) string {
	return strconv.Itoa(n) + " " + Plural(n, singular, plural)
}

// FormatNumber groups `n`'s digits in threes from the right, joined with `sep`.
// Not locale-aware: pick a separator suited to your output.
//
//	FormatNumber(1234567, ",") // "1,234,567"
//	FormatNumber(1234567, ".") // "1.234.567"
//	FormatNumber(1234567, " ") // "1 234 567"
//	FormatNumber(-42, ",")     // "-42"
func FormatNumber(n int64, sep string) string {
	s := strconv.FormatInt(n, 10)
	neg := false
	if strings.HasPrefix(s, "-") {
		neg = true
		s = s[1:]
	}
	if len(s) <= thousandsGroup {
		if neg {
			return "-" + s
		}
		return s
	}

	first := len(s) % thousandsGroup
	var b strings.Builder
	b.Grow(len(s) + len(sep)*(len(s)/thousandsGroup) + 1)
	if neg {
		b.WriteByte('-')
	}
	if first > 0 {
		b.WriteString(s[:first])
	}
	for i := first; i < len(s); i += thousandsGroup {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(s[i : i+thousandsGroup])
	}
	return b.String()
}

// compactBase is the scaling factor between successive compact-count units.
const compactBase = 1000

// compactUnits are the suffixes applied by [FormatNumberCompact], in ascending
// order: thousand, million, billion, trillion.
var compactUnits = []string{"K", "M", "B", "T"}

// FormatNumberCompact renders `n` in a compact, abbreviated form using K, M, B,
// and T suffixes (powers of 1000), with up to one decimal place and a trailing
// `.0` trimmed. Values whose magnitude is below 1000 are returned verbatim.
// Values that round up to the next unit are promoted (e.g. 999999 → `1M`), and
// magnitudes beyond a trillion stay in `T`.
//
//	FormatNumberCompact(950)      // "950"
//	FormatNumberCompact(1234)     // "1.2K"
//	FormatNumberCompact(1000000)  // "1M"
//	FormatNumberCompact(9999999)  // "10M"
//	FormatNumberCompact(-1500000) // "-1.5M"
func FormatNumberCompact(n int64) string {
	if n > -compactBase && n < compactBase {
		return strconv.FormatInt(n, 10)
	}

	neg := n < 0
	v := float64(n)
	if neg {
		v = -v
	}

	// Scale into the largest unit whose rounded value stays below the base,
	// so a value just under a boundary promotes (`1M`) instead of rendering
	// as `1000K`.
	v /= compactBase
	unit := compactUnits[0]
	for i := range compactUnits[:len(compactUnits)-1] {
		if rounded, _ := strconv.ParseFloat(
			strconv.FormatFloat(v, 'f', 1, 64),
			64,
		); rounded < compactBase {
			unit = compactUnits[i]
			break
		}
		v /= compactBase
		unit = compactUnits[i+1]
	}

	s := strings.TrimSuffix(strconv.FormatFloat(v, 'f', 1, 64), ".0")
	if neg {
		s = "-" + s
	}
	return s + unit
}

// FormatOrdinal returns `n` with its English ordinal suffix.
//
//	FormatOrdinal(1)   // "1st"
//	FormatOrdinal(22)  // "22nd"
//	FormatOrdinal(113) // "113th"
func FormatOrdinal(n int) string {
	abs := n
	if abs < 0 {
		abs = -abs
	}
	suffix := "th"
	if abs%100 < 11 || abs%100 > 13 {
		switch abs % 10 {
		case 1:
			suffix = "st"
		case 2: //nolint:mnd // 2 → "nd"
			suffix = "nd"
		case 3: //nolint:mnd // 3 → "rd"
			suffix = "rd"
		}
	}
	return strconv.Itoa(n) + suffix
}
