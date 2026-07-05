package ansi_test

import (
	"fmt"

	"github.com/gechr/x/ansi"
)

func ExampleStrip() {
	fmt.Println(ansi.Strip("\x1b[1mbold\x1b[m and plain"))
	// Output:
	// bold and plain
}

func ExampleStringWidth() {
	fmt.Println(ansi.StringWidth("\x1b[31mred\x1b[m"))
	fmt.Println(ansi.StringWidth("こんにちは"))
	// Output:
	// 3
	// 10
}

func ExampleTruncate() {
	fmt.Println(ansi.Truncate("Hello, World!", 8, "…"))
	// Output:
	// Hello, …
}

func ExampleWrapSoft() {
	fmt.Println(ansi.WrapSoft("the quick brown fox jumps over the lazy dog", 12))
	// Output:
	// the quick
	// brown fox
	// jumps over
	// the lazy dog
}

func ExampleWrapHard() {
	fmt.Println(ansi.WrapHard("the quick brown fox", 6))
	// Output:
	// the qu
	// ick br
	// own fo
	// x
}

// Breakpoints add word-break opportunities beyond spaces, which is useful
// for wrapping paths or flags.
func ExampleNewWrapper() {
	w := ansi.NewWrapper(ansi.WithWidth(10), ansi.WithBreakpoints("-"))
	fmt.Println(w.Wrap("a-very-long-flag-name"))
	// Output:
	// a-very-
	// long-flag-
	// name
}

// When output is not a terminal, hyperlinks render as plain text using the
// configured fallback mode.
func ExampleANSI_Hyperlink() {
	w := ansi.Never()
	fmt.Println(w.Hyperlink("https://example.com", "example"))

	w = ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallbackMarkdown))
	fmt.Println(w.Hyperlink("https://example.com", "example"))
	// Output:
	// example (https://example.com)
	// [example](https://example.com)
}

func ExampleCursorUp() {
	fmt.Printf("%q\n", ansi.CursorUp(3))
	// Output:
	// "\x1b[3A"
}
