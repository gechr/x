// Package slices provides slice helpers.
package slices

import (
	"slices"
	"strings"
)

// ContainsAny reports whether any of the given `lists` contains `target`.
func ContainsAny[S ~[]E, E comparable](target E, lists ...S) bool {
	return slices.ContainsFunc(lists, func(items S) bool {
		return slices.Contains(items, target)
	})
}

// ContainsAll reports whether every one of the given `lists` contains `target`.
// It returns true when no `lists` are given.
func ContainsAll[S ~[]E, E comparable](target E, lists ...S) bool {
	return !slices.ContainsFunc(lists, func(items S) bool {
		return !slices.Contains(items, target)
	})
}

// ContainsFold reports whether `items` contains `target` case-insensitively,
// using the same simple case-folding as [strings.EqualFold].
func ContainsFold[S ~[]E, E ~string](items S, target E) bool {
	return slices.ContainsFunc(items, func(item E) bool {
		return strings.EqualFold(string(item), string(target))
	})
}
