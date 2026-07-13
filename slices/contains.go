// Package slices provides slice helpers.
package slices

import (
	"slices"
	"strings"
)

// ContainedByAll reports whether `target` occurs in every one of `lists`.
// It returns true when no `lists` are given.
func ContainedByAll[S ~[]E, E comparable](target E, lists ...S) bool {
	return !slices.ContainsFunc(lists, func(items S) bool {
		return !slices.Contains(items, target)
	})
}

// ContainedByAny reports whether `target` occurs in any one of `lists`.
func ContainedByAny[S ~[]E, E comparable](target E, lists ...S) bool {
	return slices.ContainsFunc(lists, func(items S) bool {
		return slices.Contains(items, target)
	})
}

// ContainsAll reports whether `items` contains every one of `targets`.
// It returns true when no `targets` are given.
func ContainsAll[S ~[]E, E comparable](items S, targets ...E) bool {
	return !slices.ContainsFunc(targets, func(target E) bool {
		return !slices.Contains(items, target)
	})
}

// ContainsAny reports whether `items` contains any one of `targets`.
func ContainsAny[S ~[]E, E comparable](items S, targets ...E) bool {
	return slices.ContainsFunc(targets, func(target E) bool {
		return slices.Contains(items, target)
	})
}

// ContainsFold reports whether `items` contains `target` case-insensitively,
// using the same simple case-folding as [strings.EqualFold].
func ContainsFold[S ~[]E, E ~string](items S, target E) bool {
	return slices.ContainsFunc(items, func(item E) bool {
		return strings.EqualFold(string(item), string(target))
	})
}
