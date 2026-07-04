// Package slices provides slice helpers.
package slices

import (
	"slices"
	"strings"
)

// ContainsFold reports whether items contains target case-insensitively,
// using the same simple case-folding as [strings.EqualFold].
func ContainsFold[S ~[]E, E ~string](items S, target E) bool {
	return slices.ContainsFunc(items, func(item E) bool {
		return strings.EqualFold(string(item), string(target))
	})
}
