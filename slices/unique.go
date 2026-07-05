package slices

import (
	"strings"
	"unicode"
)

// Unique returns `items` in first-seen order with duplicates removed.
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

// UniqueFunc returns `items` in first-seen order with duplicates removed,
// where two items are duplicates when `key` reports the same value for both.
func UniqueFunc[S ~[]E, E any, K comparable](items S, key func(E) K) S {
	seen := make(map[K]struct{}, len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		k := key(item)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		unique = append(unique, item)
	}
	return unique
}

// UniqueFold returns strings in first-seen order with duplicates removed
// case-insensitively, using the same simple case-folding as
// [strings.EqualFold].
func UniqueFold[S ~[]E, E ~string](items S) S {
	seen := make(map[string]struct{}, len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		key := foldKey(string(item))
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		unique = append(unique, item)
	}
	return unique
}

// foldKey maps each rune to the canonical (smallest) member of its case-fold
// orbit, so two strings have equal keys iff [strings.EqualFold] reports them
// equal. ToLower alone misses orbit members with distinct lowercase forms,
// e.g. Greek final sigma 'ς' vs 'σ'.
func foldKey(s string) string {
	return strings.Map(func(r rune) rune {
		key := r
		for f := unicode.SimpleFold(r); f != r; f = unicode.SimpleFold(f) {
			key = min(key, f)
		}
		return key
	}, s)
}
