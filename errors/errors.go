// Package errors provides helpers for working with errors.
package errors

import "errors"

// IsAny reports whether any error in `targets` matches `err`'s error tree.
// Each target is compared using [errors.Is].
func IsAny(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
