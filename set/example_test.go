package set_test

import (
	"fmt"
	"slices"

	"github.com/gechr/x/set"
)

func ExampleNew() {
	s := set.New("a", "b", "a")
	fmt.Println(s.Len())
	fmt.Println(s.Contains("b"))
	fmt.Println(s.Contains("c"))
	// Output:
	// 2
	// true
	// false
}

// Collect builds a Set from any iter.Seq, such as slices.Values.
func ExampleCollect() {
	s := set.Collect(slices.Values([]int{3, 1, 3, 2}))
	fmt.Println(set.Sorted(s))
	// Output:
	// [1 2 3]
}

func ExampleSet_Union() {
	a := set.New(1, 2)
	b := set.New(2, 3)
	fmt.Println(set.Sorted(a.Union(b)))
	// Output:
	// [1 2 3]
}

func ExampleSet_Intersect() {
	a := set.New(1, 2, 3)
	b := set.New(2, 3, 4)
	fmt.Println(set.Sorted(a.Intersect(b)))
	// Output:
	// [2 3]
}

func ExampleSet_Difference() {
	a := set.New(1, 2, 3)
	b := set.New(2, 4)
	fmt.Println(set.Sorted(a.Difference(b)))
	// Output:
	// [1 3]
}

func ExampleSet_SubsetOf() {
	a := set.New("x", "y")
	b := set.New("x", "y", "z")
	fmt.Println(a.SubsetOf(b))
	fmt.Println(b.SubsetOf(a))
	// Output:
	// true
	// false
}

// Sorted turns a Set's indeterminate iteration order into a stable one.
func ExampleSorted() {
	s := set.New("banana", "apple", "cherry")
	fmt.Println(set.Sorted(s))
	// Output:
	// [apple banana cherry]
}

// SortedNatural orders embedded numbers by value, so "item2" sorts before
// "item10".
func ExampleSortedNatural() {
	s := set.New("item10", "item2", "item1")
	fmt.Println(set.Sorted(s))
	fmt.Println(set.SortedNatural(s))
	// Output:
	// [item1 item10 item2]
	// [item1 item2 item10]
}

// SortedSet keeps its items in ascending order at all times, so iteration
// is deterministic.
func ExampleNewSorted() {
	s := set.NewSorted(3, 1, 2, 1)
	s.Add(0)
	s.Delete(2)
	fmt.Println(s.Slice())
	// Output:
	// [0 1 3]
}

// CollectSorted builds a SortedSet from any iter.Seq, such as slices.Values.
func ExampleCollectSorted() {
	s := set.CollectSorted(slices.Values([]int{3, 1, 3, 2}))
	fmt.Println(s.Slice())
	// Output:
	// [1 2 3]
}

func ExampleSortedSet_All() {
	s := set.NewSorted("banana", "apple", "cherry")
	for item := range s.All() {
		fmt.Println(item)
	}
	// Output:
	// apple
	// banana
	// cherry
}
