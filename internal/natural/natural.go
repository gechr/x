// Package natural implements natural-order string comparison. It sits below
// the public packages ([github.com/gechr/x/strings], [github.com/gechr/x/set],
// [github.com/gechr/x/slices]) so they can share it without import cycles.
package natural

import (
	"cmp"
	"strings"
)

// Compare orders `a` and `b` the way a human reads them, treating each run of
// digits as a single decimal number so `x2` sorts before `x10`. It returns -1,
// 0, or +1 and allocates nothing, handling numbers of any length without
// overflow.
func Compare(a, b string) int {
	for {
		if p := commonNonDigitPrefix(a, b); p != 0 {
			a, b = a[p:], b[p:]
		}
		if a == "" {
			return cmp.Compare(0, len(b))
		}

		da, db := digitRun(a), digitRun(b)
		if da == 0 || db == 0 {
			return strings.Compare(a, b)
		}
		if c := compareNumbers(a[:da], b[:db]); c != 0 {
			return c
		}

		// Equal numeric value: descend past both runs only when each has more to
		// compare, otherwise the leading zeros decide it (e.g. `01` < `1`).
		if da == len(a) || db == len(b) {
			return strings.Compare(a, b)
		}
		a, b = a[da:], b[db:]
	}
}

// commonNonDigitPrefix returns the length of the shared leading run of `a` and `b`,
// stopping at the first differing byte or the first digit on either side so the
// numbers that follow are compared by value.
func commonNonDigitPrefix(a, b string) int {
	n := min(len(a), len(b))
	for i := range n {
		if ca, cb := a[i], b[i]; isDigit(ca) || isDigit(cb) || ca != cb {
			return i
		}
	}
	return n
}

// digitRun returns the length of the leading run of ASCII digits in `s`.
func digitRun(s string) int {
	for i := range len(s) {
		if !isDigit(s[i]) {
			return i
		}
	}
	return len(s)
}

// compareNumbers compares two non-empty digit strings by numeric value: with
// leading zeros gone, the longer string is the larger number, and equal lengths
// compare lexically.
func compareNumbers(a, b string) int {
	a, b = trimLeadingZeros(a), trimLeadingZeros(b)
	if len(a) != len(b) {
		return cmp.Compare(len(a), len(b))
	}
	return strings.Compare(a, b)
}

// trimLeadingZeros drops leading `0` bytes, yielding `""` for an all-zero string
// so it ranks below any nonzero magnitude.
func trimLeadingZeros(s string) string {
	for i := range len(s) {
		if s[i] != '0' {
			return s[i:]
		}
	}
	return s[len(s):]
}

// isDigit reports whether `c` is an ASCII digit. This package sits below
// [github.com/gechr/x/strings], so it cannot use `IsDigitChar` from there.
func isDigit(c byte) bool { return c >= '0' && c <= '9' }
