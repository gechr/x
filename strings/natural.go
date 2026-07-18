package strings

import "github.com/gechr/x/internal/natural"

// CompareNatural orders `a` and `b` the way a human reads them, treating each run of
// digits as a single decimal number so `x2` sorts before `x10`. It returns -1,
// 0, or +1 and allocates nothing, handling numbers of any length without
// overflow.
func CompareNatural(a, b string) int {
	return natural.Compare(a, b)
}

// LessNatural reports whether `a` sorts before `b` in natural order, as decided by
// [CompareNatural]. It reads cleanly at call sites that want a boolean rather
// than a three-way result, such as sort predicates and conditionals.
func LessNatural(a, b string) bool {
	return CompareNatural(a, b) < 0
}

// EqualNatural reports whether `a` and `b` compare equal in natural order, as
// decided by [CompareNatural]. This can differ from `a == b`, since a numeric
// run followed by more to compare matches regardless of leading zeros (for
// example `a00b00` and `a0b00`).
func EqualNatural(a, b string) bool {
	return CompareNatural(a, b) == 0
}
