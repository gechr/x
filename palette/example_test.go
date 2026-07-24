package palette_test

import (
	"fmt"

	"github.com/gechr/x/palette"
)

func ExamplePalette_Color() {
	p := palette.DefaultDark()

	// The same entity always resolves to the same color.
	fmt.Println(p.Color("alpha") == p.Color("alpha"))
	// Output:
	// true
}

func ExamplePalette_Color_empty() {
	// An empty palette resolves every input to nil, leaving text uncolored.
	fmt.Println(palette.Palette(nil).Color("alpha"))
	// Output:
	// <nil>
}

func ExampleNewAssigner() {
	// An Assigner gives distinct keys distinct colors (no duplicates until the
	// palette is exhausted) while keeping each key stable. With no colors it
	// defaults to Auto for the current terminal.
	a := palette.NewAssigner()

	fmt.Println(a.Assign("web") == a.Assign("web"))
	// Output:
	// true
}

func ExampleAuto() {
	// Auto detects the terminal background and true-color support, then colors
	// entities so identical strings share a stable color.
	p := palette.Auto()

	for _, entity := range []string{"alpha", "beta", "alpha"} {
		_ = p.Color(entity)
	}

	// Force the true-color Glasbey palette regardless of detection.
	p = palette.Auto(palette.WithTrueColor())
	fmt.Println(len(p))
	// Output:
	// 50
}
