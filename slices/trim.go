package slices

import "github.com/gechr/x/set"

// Trim returns `items` with all leading and trailing elements contained in
// `cutset` removed. The result is a subslice of `items`, sharing its backing
// array.
func Trim[S ~[]E, E comparable](items, cutset S) S {
	return TrimLeft(TrimRight(items, cutset), cutset)
}

// TrimLeft returns `items` with all leading elements contained in `cutset`
// removed. The result is a subslice of `items`, sharing its backing array.
func TrimLeft[S ~[]E, E comparable](items, cutset S) S {
	drop := set.New(cutset...)
	start := 0
	for start < len(items) {
		if !drop.Contains(items[start]) {
			break
		}
		start++
	}
	return items[start:]
}

// TrimRight returns `items` with all trailing elements contained in `cutset`
// removed. The result is a subslice of `items`, sharing its backing array.
func TrimRight[S ~[]E, E comparable](items, cutset S) S {
	drop := set.New(cutset...)
	end := len(items)
	for end > 0 {
		if !drop.Contains(items[end-1]) {
			break
		}
		end--
	}
	return items[:end]
}
