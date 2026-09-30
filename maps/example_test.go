package maps_test

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	xmaps "github.com/gechr/x/maps"
)

func ExampleSorted() {
	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
	for k, v := range xmaps.Sorted(m) {
		fmt.Println(k, v)
	}
	// Output:
	// alpha 1
	// beta 2
	// charlie 3
}

// SortedFunc accepts any comparison following the [cmp.Compare] convention,
// such as a descending key order.
func ExampleSortedFunc() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	descending := func(x, y int) int { return cmp.Compare(y, x) }
	for k, v := range xmaps.SortedFunc(m, descending) {
		fmt.Println(k, v)
	}
	// Output:
	// 3 three
	// 2 two
	// 1 one
}

func ExampleGroup() {
	words := []string{"apple", "banana", "avocado", "blueberry", "cherry"}
	byInitial := xmaps.Group(func(yield func(byte, string) bool) {
		for _, w := range words {
			if !yield(w[0], w) {
				return
			}
		}
	})
	for initial, group := range xmaps.Sorted(byInitial) {
		fmt.Printf("%c: %v\n", initial, group)
	}
	// Output:
	// a: [apple avocado]
	// b: [banana blueberry]
	// c: [cherry]
}

func ExampleGroupFunc() {
	words := []string{"go", "rust", "zig", "java", "c"}
	byLength := xmaps.GroupFunc(slices.Values(words), func(w string) int {
		return len(w)
	})
	for length, group := range xmaps.Sorted(byLength) {
		fmt.Println(length, group)
	}
	// Output:
	// 1 [c]
	// 2 [go]
	// 3 [zig]
	// 4 [rust java]
}

func ExampleInvert() {
	codes := map[string]int{"a": 1, "b": 2, "c": 3}
	letters := xmaps.Invert(codes)
	for code, letter := range xmaps.Sorted(letters) {
		fmt.Println(code, letter)
	}
	// Output:
	// 1 a
	// 2 b
	// 3 c
}

func ExampleMerge() {
	defaults := map[string]string{"host": "localhost", "port": "8080"}
	overrides := map[string]string{"port": "9090"}
	for k, v := range xmaps.Sorted(xmaps.Merge(defaults, overrides)) {
		fmt.Println(k, v)
	}
	// Output:
	// host localhost
	// port 9090
}

// MergeFunc resolves duplicate keys with a custom function, such as summing
// the values.
func ExampleMergeFunc() {
	monday := map[string]int{"apples": 3, "pears": 1}
	tuesday := map[string]int{"apples": 2, "plums": 5}
	sum := func(_ string, existing, incoming int) int { return existing + incoming }
	for k, v := range xmaps.Sorted(xmaps.MergeFunc(sum, monday, tuesday)) {
		fmt.Println(k, v)
	}
	// Output:
	// apples 5
	// pears 1
	// plums 5
}

func ExampleMergeDeep() {
	defaults := map[string]any{
		"server": map[string]any{"host": "localhost", "port": 8080},
		"debug":  false,
	}
	overrides := map[string]any{
		"server": map[string]any{"port": 9090},
	}
	merged := xmaps.MergeDeep(defaults, overrides)
	fmt.Println(merged["debug"])
	server, _ := merged["server"].(map[string]any)
	for k, v := range xmaps.Sorted(server) {
		fmt.Println(k, v)
	}
	// Output:
	// false
	// host localhost
	// port 9090
}

func ExampleKeys() {
	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
	keys := xmaps.Keys(m)
	slices.Sort(keys)
	fmt.Println(strings.Join(keys, ", "))
	// Output:
	// alpha, beta, charlie
}

func ExampleKeysNatural() {
	m := map[string]int{"item10": 10, "item2": 2, "item1": 1}
	fmt.Println(strings.Join(xmaps.KeysNatural(m), ", "))
	// Output:
	// item1, item2, item10
}

func ExampleValues() {
	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
	values := xmaps.Values(m)
	slices.Sort(values)
	fmt.Println(values)
	// Output:
	// [1 2 3]
}

func ExampleValuesNatural() {
	m := map[int]string{10: "item10", 2: "item2", 1: "item1"}
	fmt.Println(strings.Join(xmaps.ValuesNatural(m), ", "))
	// Output:
	// item1, item2, item10
}
