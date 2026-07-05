package slices_test

import (
	"fmt"

	xslices "github.com/gechr/x/slices"
)

func ExampleContainsFold() {
	tags := []string{"Latest", "Stable"}
	fmt.Println(xslices.ContainsFold(tags, "latest"))
	fmt.Println(xslices.ContainsFold(tags, "STABLE"))
	fmt.Println(xslices.ContainsFold(tags, "beta"))
	// Output:
	// true
	// true
	// false
}

func ExamplePartition() {
	isEven := func(n int) bool { return n%2 == 0 }
	even, odd := xslices.Partition([]int{1, 2, 3, 4, 5, 6}, isEven)
	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [2 4 6]
	// [1 3 5]
}

// Duplicates and order in the first slice are preserved.
func ExampleDifference() {
	fmt.Println(xslices.Difference([]int{1, 1, 2, 3}, []int{2}))
	fmt.Println(xslices.Difference([]int{1, 2, 3}, []int{2}, []int{3}))
	// Output:
	// [1 1 3]
	// [1]
}

func ExampleIntersect() {
	fmt.Println(xslices.Intersect([]string{"a", "b", "c"}, []string{"c", "b", "d"}))
	fmt.Println(xslices.Intersect([]int{1, 2, 3}, []int{2, 3}, []int{2}))
	// Output:
	// [b c]
	// [2]
}

func ExampleUnion() {
	fmt.Println(xslices.Union([]int{1, 2}, []int{2, 3}, []int{3, 4, 1}))
	// Output:
	// [1 2 3 4]
}

// Embedded numbers compare by value, so "item2" sorts before "item10".
func ExampleSortNatural() {
	items := []string{"item10", "item2", "item1", "item20", "item3"}
	xslices.SortNatural(items)
	fmt.Println(items)
	// Output:
	// [item1 item2 item3 item10 item20]
}

func ExampleTrim() {
	fmt.Println(xslices.Trim([]int{0, 0, 1, 2, 0}, []int{0}))
	fmt.Println(xslices.Trim([]string{"a", "b", "c", "b", "a"}, []string{"a", "b"}))
	// Output:
	// [1 2]
	// [c]
}

func ExampleUnique() {
	fmt.Println(xslices.Unique([]string{"a", "b", "a", "A", "c", "b"}))
	// Output:
	// [a b A c]
}

// The first-seen item wins for each key.
func ExampleUniqueFunc() {
	type user struct {
		name string
		age  int
	}
	users := []user{{"alice", 30}, {"bob", 25}, {"alice", 40}}
	fmt.Println(xslices.UniqueFunc(users, func(u user) string { return u.name }))
	// Output:
	// [{alice 30} {bob 25}]
}

// Every list must contain the target; no lists reports true.
func ExampleContainsAll() {
	fmt.Println(xslices.ContainsAll("a", []string{"a", "b"}, []string{"c", "a"}))
	fmt.Println(xslices.ContainsAll("b", []string{"a", "b"}, []string{"c", "a"}))
	fmt.Println(xslices.ContainsAll[[]string]("a"))
	// Output:
	// true
	// false
	// true
}

// A single list containing the target suffices; no lists reports false.
func ExampleContainsAny() {
	fmt.Println(xslices.ContainsAny("b", []string{"a", "b"}, []string{"c", "d"}))
	fmt.Println(xslices.ContainsAny("z", []string{"a", "b"}, []string{"c", "d"}))
	fmt.Println(xslices.ContainsAny[[]string]("a"))
	// Output:
	// true
	// false
	// false
}

func ExampleCount() {
	fmt.Println(xslices.Count([]string{"a", "b", "a", "c", "a"}, "a"))
	fmt.Println(xslices.Count([]string{"a", "b"}, "z"))
	// Output:
	// 3
	// 0
}

func ExampleCountFunc() {
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println(xslices.CountFunc([]int{1, 2, 3, 4, 5, 6}, isEven))
	// Output:
	// 3
}

func ExampleLastIndex() {
	fmt.Println(xslices.LastIndex([]string{"a", "b", "a", "c"}, "a"))
	fmt.Println(xslices.LastIndex([]string{"a", "b"}, "z"))
	// Output:
	// 2
	// -1
}

func ExampleLastIndexFunc() {
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println(xslices.LastIndexFunc([]int{2, 3, 4, 5}, isEven))
	fmt.Println(xslices.LastIndexFunc([]int{1, 3, 5}, isEven))
	// Output:
	// 2
	// -1
}

func ExampleTrimLeft() {
	fmt.Println(xslices.TrimLeft([]int{0, 0, 1, 2, 0}, []int{0}))
	// Output:
	// [1 2 0]
}

func ExampleTrimRight() {
	fmt.Println(xslices.TrimRight([]int{0, 0, 1, 2, 0}, []int{0}))
	// Output:
	// [0 0 1 2]
}

// The first-seen spelling wins; folding matches strings.EqualFold, so Greek
// final sigma "ς", medial "σ", and capital "Σ" are duplicates.
func ExampleUniqueFold() {
	fmt.Println(xslices.UniqueFold([]string{"Go", "GO", "go", "Rust"}))
	fmt.Println(xslices.UniqueFold([]string{"ς", "σ", "Σ"}))
	// Output:
	// [Go Rust]
	// [ς]
}
