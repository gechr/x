package strings_test

import (
	"fmt"
	"slices"

	xstrings "github.com/gechr/x/strings"
)

func ExampleClosest() {
	fmt.Printf("%q\n", xstrings.Closest("verfiy", []string{"verify", "deep"}))
	fmt.Printf("%q\n", xstrings.Closest("xyzzy", []string{"verify", "deep"}))
	// Output:
	// "verify"
	// ""
}

func ExampleCompareNatural() {
	versions := []string{"v10", "v2", "v1"}
	slices.SortFunc(versions, xstrings.CompareNatural)
	fmt.Println(versions)
	// Output:
	// [v1 v2 v10]
}

func ExampleDedent() {
	fmt.Println(xstrings.Dedent("    foo\n      bar\n    baz"))
	// Output:
	// foo
	//   bar
	// baz
}

func ExampleIndent() {
	fmt.Println(xstrings.Indent("foo\nbar", "> "))
	// Output:
	// > foo
	// > bar
}

func ExampleSplitCSV() {
	fmt.Printf("%q\n", xstrings.SplitCSV(" a, b ,, c "))
	// Output:
	// ["a" "b" "c"]
}

func ExampleTruncateMiddle() {
	fmt.Println(xstrings.TruncateMiddle("0123456789abcdef", 7, "…"))
	// Output:
	// 012…def
}

func ExampleUnwrap() {
	fmt.Println(xstrings.Unwrap(`"quoted"`, `"`, `"`))
	fmt.Println(xstrings.Unwrap(`"one-sided`, `"`, `"`))
	// Output:
	// quoted true
	// "one-sided false
}

// PadCenter places the odd rune of padding on the right.
func ExamplePadCenter() {
	fmt.Printf("%q\n", xstrings.PadCenter("hi", 5))
	// Output:
	// " hi  "
}

func ExampleAppendCSV() {
	fmt.Printf("%q\n", xstrings.AppendCSV([]string{"x"}, " a, b ,, c "))
	// Output:
	// ["x" "a" "b" "c"]
}

func ExampleAnyEmpty() {
	fmt.Println(xstrings.AnyEmpty("alpha", "", "beta"))
	fmt.Println(xstrings.AnyEmpty("alpha", "beta"))
	// Output:
	// true
	// false
}

func ExampleAnyNonEmpty() {
	fmt.Println(xstrings.AnyNonEmpty("", "alpha", ""))
	fmt.Println(xstrings.AnyNonEmpty("", ""))
	// Output:
	// true
	// false
}

func ExampleAllEmpty() {
	fmt.Println(xstrings.AllEmpty("", ""))
	fmt.Println(xstrings.AllEmpty("", "alpha"))
	// Output:
	// true
	// false
}

func ExampleAllNonEmpty() {
	fmt.Println(xstrings.AllNonEmpty("alpha", "beta", "charlie"))
	fmt.Println(xstrings.AllNonEmpty("alpha", ""))
	// Output:
	// true
	// false
}

func ExampleCompactLines() {
	fmt.Println(xstrings.CompactLines("  foo \n\nbar\nfoo\n", ", "))
	// Output:
	// foo, bar
}

func ExampleCompareFold() {
	fmt.Println(xstrings.CompareFold("Go", "go"))
	fmt.Println(xstrings.CompareFold("abc", "ABD"))
	fmt.Println(xstrings.CompareFold("B", "a"))
	// Output:
	// 0
	// -1
	// 1
}

func ExampleContainsAll() {
	fmt.Println(xstrings.ContainsAll("hello world", "hello", "world"))
	fmt.Println(xstrings.ContainsAll("hello world", "hello", "moon"))
	// Output:
	// true
	// false
}

func ExampleContainsAny() {
	fmt.Println(xstrings.ContainsAny("hello world", "moon", "world"))
	fmt.Println(xstrings.ContainsAny("hello world", "moon", "sun"))
	// Output:
	// true
	// false
}

func ExampleCountAny() {
	fmt.Println(xstrings.CountAny("hello world", "lo"))
	// Output:
	// 5
}

func ExampleEnsureTrailingNewline() {
	fmt.Printf("%q\n", xstrings.EnsureTrailingNewline("hello\n\n"))
	fmt.Printf("%q\n", xstrings.EnsureTrailingNewline("hello"))
	// Output:
	// "hello\n"
	// "hello\n"
}

// Leading zeros are ignored when more text follows the numeric run.
func ExampleEqualNatural() {
	fmt.Println(xstrings.EqualNatural("a00b00", "a0b00"))
	fmt.Println(xstrings.EqualNatural("a1", "a2"))
	// Output:
	// true
	// false
}

func ExampleHexEqual() {
	fmt.Println(xstrings.HexEqual("0xDEADbeef", "deadbeef"))
	fmt.Println(xstrings.HexEqual("0x1234", "0x5678"))
	// Output:
	// true
	// false
}

func ExampleIsBlank() {
	fmt.Println(xstrings.IsBlank(" \t\n"))
	fmt.Println(xstrings.IsBlank("x"))
	// Output:
	// true
	// false
}

func ExampleIsDigits() {
	fmt.Println(xstrings.IsDigits("12345"))
	fmt.Println(xstrings.IsDigits("12a45"))
	fmt.Println(xstrings.IsDigits(""))
	// Output:
	// true
	// false
	// false
}

func ExampleIsGitCommit() {
	fmt.Println(xstrings.IsGitCommit("3b18e512dba79e4c8300dd08aeb37f8e728b8dad"))
	fmt.Println(xstrings.IsGitCommit("deadbeef"))
	// Output:
	// true
	// false
}

func ExampleIsHex() {
	fmt.Println(xstrings.IsHex("deadBEEF42"))
	fmt.Println(xstrings.IsHex("xyz"))
	// Output:
	// true
	// false
}

func ExampleIsHexChar() {
	fmt.Println(xstrings.IsHexChar('f'))
	fmt.Println(xstrings.IsHexChar('F'))
	fmt.Println(xstrings.IsHexChar('9'))
	fmt.Println(xstrings.IsHexChar('g'))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsSHA256() {
	fmt.Println(
		xstrings.IsSHA256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
	)
	fmt.Println(xstrings.IsSHA256("deadbeef"))
	// Output:
	// true
	// false
}

func ExampleLessNatural() {
	fmt.Println(xstrings.LessNatural("v2", "v10"))
	fmt.Println(xstrings.LessNatural("v10", "v2"))
	// Output:
	// true
	// false
}

func ExamplePadLeft() {
	fmt.Printf("%q\n", xstrings.PadLeft("hi", 5))
	// Output:
	// "   hi"
}

func ExamplePadRight() {
	fmt.Printf("%q\n", xstrings.PadRight("hi", 5))
	// Output:
	// "hi   "
}

// Empty segments between adjacent separators are preserved.
func ExampleSplitAny() {
	fmt.Printf("%q\n", xstrings.SplitAny("a,b;;c", ",;"))
	// Output:
	// ["a" "b" "" "c"]
}

func ExampleSplitBy() {
	fmt.Printf("%q\n", xstrings.SplitBy(" a | b || c ", "|"))
	// Output:
	// ["a" "b" "c"]
}

func ExampleSplitLines() {
	fmt.Printf("%q\n", xstrings.SplitLines("foo\n\n  bar \n"))
	// Output:
	// ["foo" "bar"]
}

func ExampleSplitLinesRaw() {
	fmt.Printf("%q\n", xstrings.SplitLinesRaw("foo\r\nbar\n"))
	// Output:
	// ["foo" "bar" ""]
}

func ExampleTruncate() {
	fmt.Println(xstrings.Truncate("hello world", 8, "…"))
	// Output:
	// hello w…
}

func ExampleTruncateLeft() {
	fmt.Println(xstrings.TruncateLeft("hello world", 8, "…"))
	// Output:
	// …o world
}

func ExampleTruncateRight() {
	fmt.Println(xstrings.TruncateRight("hello world", 8, "…"))
	fmt.Println(xstrings.TruncateRight("hi", 8, "…"))
	// Output:
	// hello w…
	// hi
}
