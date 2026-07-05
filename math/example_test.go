package math_test

import (
	"fmt"
	"math"

	xmath "github.com/gechr/x/math"
)

func ExampleClamp() {
	fmt.Println(xmath.Clamp(5, 0, 10))
	fmt.Println(xmath.Clamp(-3, 0, 10))
	fmt.Println(xmath.Clamp(42, 0, 10))
	// Output:
	// 5
	// 0
	// 10
}

// Clamp works with any ordered type, including strings.
func ExampleClamp_strings() {
	fmt.Println(xmath.Clamp("a", "b", "d"))
	fmt.Println(xmath.Clamp("c", "b", "d"))
	fmt.Println(xmath.Clamp("z", "b", "d"))
	// Output:
	// b
	// c
	// d
}

// Unlike min(max(v, lo), hi), NaN does not propagate - it clamps to `lo`.
func ExampleClamp_nan() {
	fmt.Println(xmath.Clamp(math.NaN(), 1.0, 2.0))
	fmt.Println(min(max(math.NaN(), 1.0), 2.0))
	// Output:
	// 1
	// NaN
}

func ExampleClamp01() {
	fmt.Println(xmath.Clamp01(0.5))
	fmt.Println(xmath.Clamp01(-0.1))
	fmt.Println(xmath.Clamp01(1.1))
	// Output:
	// 0.5
	// 0
	// 1
}
