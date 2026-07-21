package bytes_test

import (
	"fmt"

	xbytes "github.com/gechr/x/bytes"
)

func ExampleDedent() {
	fmt.Printf("%s\n", xbytes.Dedent([]byte("    foo\n      bar\n    baz")))
	// Output:
	// foo
	//   bar
	// baz
}

func ExampleIndent() {
	fmt.Printf("%s\n", xbytes.Indent([]byte("foo\nbar"), []byte("> ")))
	// Output:
	// > foo
	// > bar
}

func ExampleTruncateMiddle() {
	fmt.Printf("%s\n", xbytes.TruncateMiddle([]byte("0123456789abcdef"), 7, []byte("…")))
	// Output:
	// 012…def
}

func ExampleUnwrap() {
	quoted, ok := xbytes.Unwrap([]byte(`"quoted"`), []byte(`"`), []byte(`"`))
	fmt.Println(string(quoted), ok)
	oneSided, ok := xbytes.Unwrap([]byte(`"one-sided`), []byte(`"`), []byte(`"`))
	fmt.Println(string(oneSided), ok)
	// Output:
	// quoted true
	// "one-sided false
}

// PadCenter places the odd rune of padding on the right.
func ExamplePadCenter() {
	fmt.Printf("%q\n", xbytes.PadCenter([]byte("hi"), 5))
	// Output:
	// " hi  "
}

func ExampleAnyEmpty() {
	fmt.Println(xbytes.AnyEmpty([]byte("alpha"), []byte(""), []byte("beta")))
	fmt.Println(xbytes.AnyEmpty([]byte("alpha"), []byte("beta")))
	// Output:
	// true
	// false
}

func ExampleAnyNonEmpty() {
	fmt.Println(xbytes.AnyNonEmpty([]byte(""), []byte("alpha"), []byte("")))
	fmt.Println(xbytes.AnyNonEmpty([]byte(""), []byte("")))
	// Output:
	// true
	// false
}

func ExampleAllEmpty() {
	fmt.Println(xbytes.AllEmpty([]byte(""), []byte("")))
	fmt.Println(xbytes.AllEmpty([]byte(""), []byte("alpha")))
	// Output:
	// true
	// false
}

func ExampleAllNonEmpty() {
	fmt.Println(xbytes.AllNonEmpty([]byte("alpha"), []byte("beta"), []byte("charlie")))
	fmt.Println(xbytes.AllNonEmpty([]byte("alpha"), []byte("")))
	// Output:
	// true
	// false
}

func ExampleCompareFold() {
	fmt.Println(xbytes.CompareFold([]byte("Go"), []byte("go")))
	fmt.Println(xbytes.CompareFold([]byte("abc"), []byte("ABD")))
	fmt.Println(xbytes.CompareFold([]byte("B"), []byte("a")))
	// Output:
	// 0
	// -1
	// 1
}

func ExampleContainsAll() {
	fmt.Println(xbytes.ContainsAll([]byte("hello world"), []byte("hello"), []byte("world")))
	fmt.Println(xbytes.ContainsAll([]byte("hello world"), []byte("hello"), []byte("moon")))
	// Output:
	// true
	// false
}

func ExampleContainsAny() {
	fmt.Println(xbytes.ContainsAny([]byte("hello world"), []byte("moon"), []byte("world")))
	fmt.Println(xbytes.ContainsAny([]byte("hello world"), []byte("moon"), []byte("sun")))
	// Output:
	// true
	// false
}

func ExampleContainsFold() {
	fmt.Println(xbytes.ContainsFold([]byte("Hello, World"), []byte("WORLD")))
	fmt.Println(xbytes.ContainsFold([]byte("Hello, World"), []byte("moon")))
	// Output:
	// true
	// false
}

func ExampleCountAny() {
	fmt.Println(xbytes.CountAny([]byte("hello world"), "lo"))
	// Output:
	// 5
}

func ExampleEnsureTrailingNewline() {
	fmt.Printf("%q\n", xbytes.EnsureTrailingNewline([]byte("hello\n\n")))
	fmt.Printf("%q\n", xbytes.EnsureTrailingNewline([]byte("hello")))
	// Output:
	// "hello\n"
	// "hello\n"
}

func ExampleHexEqual() {
	fmt.Println(xbytes.HexEqual([]byte("0xDEADbeef"), []byte("deadbeef")))
	fmt.Println(xbytes.HexEqual([]byte("0x1234"), []byte("0x5678")))
	// Output:
	// true
	// false
}

func ExampleHasPrefixFold() {
	fmt.Println(xbytes.HasPrefixFold([]byte("Hello, World"), []byte("HELLO")))
	fmt.Println(xbytes.HasPrefixFold([]byte("Hello, World"), []byte("world")))
	// Output:
	// true
	// false
}

func ExampleHasSuffixFold() {
	fmt.Println(xbytes.HasSuffixFold([]byte("Hello, World"), []byte("WORLD")))
	fmt.Println(xbytes.HasSuffixFold([]byte("Hello, World"), []byte("hello")))
	// Output:
	// true
	// false
}

func ExampleIsBlank() {
	fmt.Println(xbytes.IsBlank([]byte(" \t\n")))
	fmt.Println(xbytes.IsBlank([]byte("x")))
	// Output:
	// true
	// false
}

func ExampleIsDigits() {
	fmt.Println(xbytes.IsDigits([]byte("12345")))
	fmt.Println(xbytes.IsDigits([]byte("12a45")))
	fmt.Println(xbytes.IsDigits([]byte("")))
	// Output:
	// true
	// false
	// false
}

func ExampleIsGitCommit() {
	fmt.Println(xbytes.IsGitCommit([]byte("3b18e512dba79e4c8300dd08aeb37f8e728b8dad")))
	fmt.Println(xbytes.IsGitCommit([]byte("deadbeef")))
	// Output:
	// true
	// false
}

func ExampleIsHex() {
	fmt.Println(xbytes.IsHex([]byte("deadBEEF42")))
	fmt.Println(xbytes.IsHex([]byte("xyz")))
	// Output:
	// true
	// false
}

func ExampleIsHexChar() {
	fmt.Println(xbytes.IsHexChar('f'))
	fmt.Println(xbytes.IsHexChar('F'))
	fmt.Println(xbytes.IsHexChar('9'))
	fmt.Println(xbytes.IsHexChar('g'))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsSHA256() {
	fmt.Println(
		xbytes.IsSHA256([]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")),
	)
	fmt.Println(xbytes.IsSHA256([]byte("deadbeef")))
	// Output:
	// true
	// false
}

func ExampleDecodeSHA256() {
	digest, err := xbytes.DecodeSHA256(
		[]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
	)
	fmt.Printf("%x\n", digest[:4])
	fmt.Println(err)
	// Output:
	// e3b0c442
	// <nil>
}

func ExamplePadLeft() {
	fmt.Printf("%q\n", xbytes.PadLeft([]byte("hi"), 5))
	// Output:
	// "   hi"
}

func ExamplePadRight() {
	fmt.Printf("%q\n", xbytes.PadRight([]byte("hi"), 5))
	// Output:
	// "hi   "
}

// Empty segments between adjacent separators are preserved.
func ExampleSplitAny() {
	fmt.Printf("%q\n", xbytes.SplitAny([]byte("a,b;;c"), ",;"))
	// Output:
	// ["a" "b" "" "c"]
}

func ExampleSplitBy() {
	fmt.Printf("%q\n", xbytes.SplitBy([]byte(" a | b || c "), []byte("|")))
	// Output:
	// ["a" "b" "c"]
}

func ExampleSplitLines() {
	fmt.Printf("%q\n", xbytes.SplitLines([]byte("foo\n\n  bar \n")))
	// Output:
	// ["foo" "bar"]
}

func ExampleSplitLinesRaw() {
	fmt.Printf("%q\n", xbytes.SplitLinesRaw([]byte("foo\r\nbar\n")))
	// Output:
	// ["foo" "bar" ""]
}
