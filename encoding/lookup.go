package encoding

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	xmaps "github.com/gechr/x/maps"
	xslices "github.com/gechr/x/slices"
	xstrings "github.com/gechr/x/strings"
)

// Lookup resolves `p` against `doc`, a document decoded into generic Go
// values (map[string]any, map[any]any, []any), and returns the value it
// addresses. It reports false when a segment is missing, an index is out of
// range, or the path contains a wildcard segment (use [Path.LookupAll]).
func (p *Path) Lookup(doc any) (any, bool) {
	value := doc
	for _, seg := range p.segments() {
		next, ok := seg.lookup(value)
		if !ok {
			return nil, false
		}
		value = next
	}
	return value, true
}

// LookupAll resolves `p` against `doc`, fanning out at wildcard segments:
// each [Path.Wildcard] matches every element of an array in order, or every
// value of a map in natural key order (see
// [github.com/gechr/x/strings.CompareNatural], with ties broken lexically and
// then by the key's Go type). Values missing a later segment are skipped.
// Without wildcards the result has at most one value. Returns nil when
// nothing matches.
func (p *Path) LookupAll(doc any) []any {
	values := []any{doc}
	for _, seg := range p.segments() {
		var next []any
		for _, value := range values {
			if seg.kind == segmentWildcard {
				next = append(next, elements(value)...)
				continue
			}
			if v, ok := seg.lookup(value); ok {
				next = append(next, v)
			}
		}
		if len(next) == 0 {
			return nil
		}
		values = next
	}
	return values
}

// lookup resolves a single non-wildcard segment against `value`.
func (p *Path) lookup(value any) (any, bool) {
	switch p.kind {
	case segmentChild, segmentKey:
		return lookupKey(value, p.name)
	case segmentIndex:
		list, ok := value.([]any)
		if !ok || p.index < 0 || p.index >= len(list) {
			return nil, false
		}
		return list[p.index], true
	case segmentWildcard:
		return nil, false
	}
	return nil, false
}

// lookupKey returns the value of key `name` in `value` when it is a map. For
// map[any]any, an exact string key wins; otherwise any key whose string form
// (fmt.Sprint) equals `name` matches, so numeric keys from legacy YAML
// decoders ("2: two") are addressable as Child("2"). Colliding stringified
// keys resolve to the lexically smallest Go type name, for determinism.
func lookupKey(value any, name string) (any, bool) {
	switch m := value.(type) {
	case map[string]any:
		v, ok := m[name]
		return v, ok
	case map[any]any:
		if v, ok := m[name]; ok {
			return v, ok
		}
		var (
			best     any
			bestType string
			found    bool
		)
		for k, v := range m {
			if fmt.Sprint(k) != name {
				continue
			}
			keyType := fmt.Sprintf("%T", k)
			if !found || keyType < bestType {
				best, bestType, found = v, keyType, true
			}
		}
		return best, found
	}
	return nil, false
}

// elements returns the wildcard fan-out of `value`: array elements in order,
// or map values in natural key order (ties broken lexically, then by key
// type) so results are deterministic.
func elements(value any) []any {
	switch v := value.(type) {
	case []any:
		return v
	case map[string]any:
		keys := xmaps.Keys(v)
		slices.SortFunc(keys, compareKeys)
		return xslices.Map(keys, func(k string) any { return v[k] })
	case map[any]any:
		// Pair each key with its value during iteration: re-fetching by key
		// after sorting would miss keys that are not equal to themselves
		// (math.NaN()).
		type entry struct {
			key     string
			keyType string
			value   any
		}
		entries := make([]entry, 0, len(v))
		for k, val := range v {
			entries = append(entries, entry{fmt.Sprint(k), fmt.Sprintf("%T", k), val})
		}
		slices.SortFunc(entries, func(a, b entry) int {
			return cmp.Or(compareKeys(a.key, b.key), strings.Compare(a.keyType, b.keyType))
		})
		return xslices.Map(entries, func(e entry) any { return e.value })
	}
	return nil
}

// compareKeys orders map keys naturally, falling back to lexical order when
// distinct keys compare naturally equal (e.g. "a0" vs "a00") so fan-out order
// does not inherit randomized map iteration.
func compareKeys(a, b string) int {
	return cmp.Or(xstrings.CompareNatural(a, b), strings.Compare(a, b))
}
