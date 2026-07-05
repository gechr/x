package ptr_test

import (
	"fmt"

	"github.com/gechr/x/ptr"
)

func ExampleDeref() {
	s := "hello"
	fmt.Println(ptr.Deref(&s))
	// Output:
	// hello
}

// A nil pointer dereferences to the zero value instead of panicking.
func ExampleDeref_nil() {
	var s *string
	var n *int
	fmt.Printf("%q\n", ptr.Deref(s))
	fmt.Println(ptr.Deref(n))
	// Output:
	// ""
	// 0
}

// Deref is handy for optional struct fields modelled as pointers.
func ExampleDeref_optionalField() {
	type Config struct {
		Retries *int
	}

	retries := 3
	fmt.Println(ptr.Deref(Config{Retries: &retries}.Retries))
	fmt.Println(ptr.Deref(Config{}.Retries))
	// Output:
	// 3
	// 0
}
