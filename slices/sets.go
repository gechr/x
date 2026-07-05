package slices

// Difference returns the elements of `items` not present in any of `others`,
// preserving order and duplicates from `items`.
func Difference[S ~[]E, E comparable](items S, others ...S) S {
	drop := make(map[E]struct{})
	for _, other := range others {
		for item := range toSet(other) {
			drop[item] = struct{}{}
		}
	}
	diff := make(S, 0, len(items))
	for _, item := range items {
		if _, ok := drop[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// Intersect returns the elements of `items` also present in every one of
// `others`, preserving order and duplicates from `items`.
func Intersect[S ~[]E, E comparable](items S, others ...S) S {
	sets := make([]map[E]struct{}, len(others))
	for i, other := range others {
		sets[i] = toSet(other)
	}
	both := make(S, 0, len(items))
	for _, item := range items {
		inAll := true
		for _, set := range sets {
			if _, ok := set[item]; !ok {
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

func toSet[S ~[]E, E comparable](items S) map[E]struct{} {
	set := make(map[E]struct{}, len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}
