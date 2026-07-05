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

func ExampleKeysSlice() {
	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
	keys := xmaps.KeysSlice(m)
	slices.Sort(keys)
	fmt.Println(strings.Join(keys, ", "))
	// Output:
	// alpha, beta, charlie
}

func ExampleValuesSlice() {
	m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
	values := xmaps.ValuesSlice(m)
	slices.Sort(values)
	fmt.Println(values)
	// Output:
	// [1 2 3]
}
