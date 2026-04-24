package slices

import "strings"

// Unique returns items in first-seen order with duplicates removed.
func Unique[S ~[]E, E comparable](items S) S {
	seen := make(map[E]struct{}, len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		unique = append(unique, item)
	}
	return unique
}

// UniqueFold returns strings in first-seen order with duplicates removed
// case-insensitively.
func UniqueFold[S ~[]E, E ~string](items S) S {
	seen := make(map[string]struct{}, len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		key := strings.ToLower(string(item))
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		unique = append(unique, item)
	}
	return unique
}
