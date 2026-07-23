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

func ExampleAuto() {
	// Auto detects the terminal background and true-color support, then colors
	// entities so identical strings share a stable color.
	p := palette.Auto()

	for _, entity := range []string{"alpha", "beta", "alpha"} {
		_ = p.Color(entity)
	}

	// Force the 256-color Glasbey palette regardless of detection.
	p = palette.Auto(palette.WithTrueColor())
	fmt.Println(len(p))
	// Output:
	// 256
}
