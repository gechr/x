package terminal_test

import (
	"fmt"
	"os"

	"github.com/gechr/x/terminal"
)

// A pipe is not a terminal, and nil files are always reported as non-terminals.
func ExampleIs() {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer r.Close()
	defer w.Close()

	fmt.Println(terminal.Is(r))
	fmt.Println(terminal.Is(w))
	fmt.Println(terminal.Is(nil))
	// Output:
	// false
	// false
	// false
}

// Size returns (0, 0) when the file is nil or not connected to a terminal.
func ExampleSize() {
	w, h := terminal.Size(nil)
	fmt.Println(w, h)
	// Output:
	// 0 0
}

// Width returns 0 when the file is nil or not connected to a terminal.
func ExampleWidth() {
	fmt.Println(terminal.Width(nil))
	// Output:
	// 0
}

// Height returns 0 when the file is nil or not connected to a terminal.
func ExampleHeight() {
	fmt.Println(terminal.Height(nil))
	// Output:
	// 0
}
