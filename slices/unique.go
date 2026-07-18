package slices

import (
	"strings"
	"unicode"

	"github.com/gechr/x/set"
)

// Unique returns `items` in first-seen order with duplicates removed.
func Unique[S ~[]E, E comparable](items S) S {
	seen := make(set.Set[E], len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		if seen.Contains(item) {
			continue
		}
		seen.Add(item)
		unique = append(unique, item)
	}
	return unique
}

// UniqueFunc returns `items` in first-seen order with duplicates removed,
// where two items are duplicates when `key` reports the same value for both.
func UniqueFunc[S ~[]E, E any, K comparable](items S, key func(E) K) S {
	seen := make(set.Set[K], len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		k := key(item)
		if seen.Contains(k) {
			continue
		}
		seen.Add(k)
		unique = append(unique, item)
	}
	return unique
}

// UniqueFold returns strings in first-seen order with duplicates removed
// case-insensitively, using the same simple case-folding as
// [strings.EqualFold].
func UniqueFold[S ~[]E, E ~string](items S) S {
	seen := make(set.Set[string], len(items))
	unique := make(S, 0, len(items))
	for _, item := range items {
		key := foldKey(string(item))
		if seen.Contains(key) {
			continue
		}
		seen.Add(key)
		unique = append(unique, item)
	}
	return unique
}

// foldKey maps each rune to the canonical (smallest) member of its case-fold
// orbit, so two strings have equal keys iff [strings.EqualFold] reports them
// equal. [strings.ToLower] alone misses orbit members with distinct lowercase
// forms, e.g. Greek final sigma 'ς' vs 'σ'.
func foldKey(s string) string {
	return strings.Map(func(r rune) rune {
		key := r
		for f := unicode.SimpleFold(r); f != r; f = unicode.SimpleFold(f) {
			key = min(key, f)
		}
		return key
	}, s)
}
