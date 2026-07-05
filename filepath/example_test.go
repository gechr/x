package filepath_test

import (
	"fmt"

	xfilepath "github.com/gechr/x/filepath"
)

func ExampleIsWithin() {
	fmt.Println(xfilepath.IsWithin("src", "src/foo.go"))
	fmt.Println(xfilepath.IsWithin("src", "src"))
	fmt.Println(xfilepath.IsWithin("src", "lib/foo.go"))
	fmt.Println(xfilepath.IsWithin("src"))
	// Output:
	// true
	// true
	// false
	// false
}

// IsWithin only reports true when every target is contained within the base.
func ExampleIsWithin_multipleTargets() {
	fmt.Println(xfilepath.IsWithin(".", "src/foo.go", "lib/bar.go"))
	fmt.Println(xfilepath.IsWithin("src", "src/foo.go", "lib/bar.go"))
	// Output:
	// true
	// false
}

func ExampleMerge() {
	fmt.Println(xfilepath.Merge([]string{".", "./sub"}))
	fmt.Println(xfilepath.Merge([]string{"a/b", "a"}))
	fmt.Println(xfilepath.Merge([]string{"a", "b"}))
	// Output:
	// [.]
	// [a]
	// [a b]
}

// Exact duplicates are merged; the first occurrence survives in its
// original spelling.
func ExampleMerge_duplicates() {
	fmt.Println(xfilepath.Merge([]string{"a", "./a", "a/"}))
	// Output:
	// [a]
}
