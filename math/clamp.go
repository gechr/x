// Package math provides numeric helpers.
package math

import "cmp"

// Clamp restricts v to the [lo, hi] range. NaN clamps to lo; infinities clamp
// to the nearest bound. Unlike min(max(v, lo), hi), NaN does not propagate.
func Clamp[T cmp.Ordered](v, lo, hi T) T {
	if v != v { //nolint:gocritic // NaN is the only value that compares unequal to itself
		return lo
	}
	return max(lo, min(hi, v))
}

// Clamp01 restricts v to the [0, 1] range. NaN clamps to 0.
func Clamp01(v float64) float64 {
	return Clamp(v, 0, 1)
}
