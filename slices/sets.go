package slices

import "github.com/gechr/x/set"

// Difference returns the elements of `items` not present in any of `others`,
// preserving order and duplicates from `items`.
func Difference[S ~[]E, E comparable](items S, others ...S) S {
	drop := make(set.Set[E])
	for _, other := range others {
		drop.Add(other...)
	}
	diff := make(S, 0, len(items))
	for _, item := range items {
		if !drop.Contains(item) {
			diff = append(diff, item)
		}
	}
	return diff
}

// Intersect returns the elements of `items` also present in every one of
// `others`, preserving order and duplicates from `items`.
func Intersect[S ~[]E, E comparable](items S, others ...S) S {
	sets := make([]set.Set[E], len(others))
	for i, other := range others {
		sets[i] = set.New(other...)
	}
	both := make(S, 0, len(items))
	for _, item := range items {
		inAll := true
		for _, s := range sets {
			if !s.Contains(item) {
				inAll = false
				break
			}
		}
		if inAll {
			both = append(both, item)
		}
	}
	return both
}

// Union returns the elements of `items` followed by the elements of `others`, in
// first-seen order with duplicates removed.
func Union[S ~[]E, E comparable](items S, others ...S) S {
	all := make(S, 0, len(items))
	all = append(all, items...)
	for _, other := range others {
		all = append(all, other...)
	}
	return Unique(all)
}
