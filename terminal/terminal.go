// Package terminal provides terminal detection and size queries.
package terminal

import (
	"os"

	"golang.org/x/term"
)

// Is returns true if the given file is a terminal.
// Returns false for nil files.
func Is(f *os.File) bool {
	if f == nil {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}

// Width returns the width of the terminal connected to f.
// Returns 0 if f is nil or not a terminal.
func Width(f *os.File) int {
	w, _ := Size(f)
	return w
}

// Height returns the height of the terminal connected to f, or 0 if f is nil
// or not a terminal.
func Height(f *os.File) int {
	_, h := Size(f)
	return h
}

// Size returns the (width, height) of the terminal connected to f in cells,
// or (0, 0) if f is nil or not a terminal.
func Size(f *os.File) (int, int) {
	if f == nil {
		return 0, 0
	}
	w, h, err := term.GetSize(int(f.Fd()))
	if err != nil {
		return 0, 0
	}
	return w, h
}
