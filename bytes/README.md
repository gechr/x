# bytes

```go
import "github.com/gechr/x/bytes"
```

Package `bytes` provides byte-slice helpers mirroring [strings](<../strings/README.md>): split, contains, indent/dedent, truncate, and blank checks.

## Index

- [func AllEmpty(values ...\[\]byte) bool](<#AllEmpty>)
- [func AllNonEmpty(values ...\[\]byte) bool](<#AllNonEmpty>)
- [func AnyEmpty(values ...\[\]byte) bool](<#AnyEmpty>)
- [func AnyNonEmpty(values ...\[\]byte) bool](<#AnyNonEmpty>)
- [func CompareFold(a, b \[\]byte) int](<#CompareFold>)
- [func ContainsAll(s \[\]byte, subslices ...\[\]byte) bool](<#ContainsAll>)
- [func ContainsAny(s \[\]byte, subslices ...\[\]byte) bool](<#ContainsAny>)
- [func ContainsFold(s, subslice \[\]byte) bool](<#ContainsFold>)
- [func CountAny(s \[\]byte, chars string) int](<#CountAny>)
- [func DecodeSHA256(s \[\]byte) (\[sha256.Size\]byte, error)](<#DecodeSHA256>)
- [func Dedent(s \[\]byte) \[\]byte](<#Dedent>)
- [func EnsureTrailingNewline(s \[\]byte) \[\]byte](<#EnsureTrailingNewline>)
- [func HasPrefixFold(s, prefix \[\]byte) bool](<#HasPrefixFold>)
- [func HasSuffixFold(s, suffix \[\]byte) bool](<#HasSuffixFold>)
- [func HexEqual(a, b \[\]byte) bool](<#HexEqual>)
- [func Indent(s, prefix \[\]byte) \[\]byte](<#Indent>)
- [func IsASCII(s \[\]byte) bool](<#IsASCII>)
- [func IsAlpha(s \[\]byte) bool](<#IsAlpha>)
- [func IsAlphaChar(c byte) bool](<#IsAlphaChar>)
- [func IsAlphanumeric(s \[\]byte) bool](<#IsAlphanumeric>)
- [func IsAlphanumericChar(c byte) bool](<#IsAlphanumericChar>)
- [func IsBlank(s \[\]byte) bool](<#IsBlank>)
- [func IsDigitChar(c byte) bool](<#IsDigitChar>)
- [func IsDigits(s \[\]byte) bool](<#IsDigits>)
- [func IsGitCommit(s \[\]byte) bool](<#IsGitCommit>)
- [func IsHex(s \[\]byte) bool](<#IsHex>)
- [func IsHexChar(c byte) bool](<#IsHexChar>)
- [func IsSHA256(s \[\]byte) bool](<#IsSHA256>)
- [func PadCenter(s \[\]byte, width int) \[\]byte](<#PadCenter>)
- [func PadLeft(s \[\]byte, width int) \[\]byte](<#PadLeft>)
- [func PadRight(s \[\]byte, width int) \[\]byte](<#PadRight>)
- [func SplitAny(s \[\]byte, chars string) \[\]\[\]byte](<#SplitAny>)
- [func SplitBy(s, sep \[\]byte) \[\]\[\]byte](<#SplitBy>)
- [func SplitLines(s \[\]byte) \[\]\[\]byte](<#SplitLines>)
- [func SplitLinesRaw(s \[\]byte) \[\]\[\]byte](<#SplitLinesRaw>)
- [func TrimPrefixes(s \[\]byte, prefixes ...\[\]byte) \[\]byte](<#TrimPrefixes>)
- [func TrimSuffixes(s \[\]byte, suffixes ...\[\]byte) \[\]byte](<#TrimSuffixes>)
- [func Truncate(s \[\]byte, n int, marker \[\]byte) \[\]byte](<#Truncate>)
- [func TruncateLeft(s \[\]byte, n int, marker \[\]byte) \[\]byte](<#TruncateLeft>)
- [func TruncateMiddle(s \[\]byte, n int, marker \[\]byte) \[\]byte](<#TruncateMiddle>)
- [func TruncateRight(s \[\]byte, n int, marker \[\]byte) \[\]byte](<#TruncateRight>)
- [func Unwrap(s, prefix, suffix \[\]byte) (\[\]byte, bool)](<#Unwrap>)

<a name="AllEmpty"></a>

## func [AllEmpty](<https://github.com/gechr/x/blob/main/bytes/blank.go#L31>)

```go
func AllEmpty(values ...[]byte) bool
```

**AllEmpty** reports whether every given slice is empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.AllEmpty([]byte(""), []byte("")))
fmt.Println(xbytes.AllEmpty([]byte(""), []byte("alpha")))
```

Output:

```text
true
false
```

</details>

<a name="AllNonEmpty"></a>

## func [AllNonEmpty](<https://github.com/gechr/x/blob/main/bytes/blank.go#L41>)

```go
func AllNonEmpty(values ...[]byte) bool
```

**AllNonEmpty** reports whether every given slice is non-empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.AllNonEmpty([]byte("alpha"), []byte("beta"), []byte("charlie")))
fmt.Println(xbytes.AllNonEmpty([]byte("alpha"), []byte("")))
```

Output:

```text
true
false
```

</details>

<a name="AnyEmpty"></a>

## func [AnyEmpty](<https://github.com/gechr/x/blob/main/bytes/blank.go#L11>)

```go
func AnyEmpty(values ...[]byte) bool
```

**AnyEmpty** reports whether any of the given slices is empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.AnyEmpty([]byte("alpha"), []byte(""), []byte("beta")))
fmt.Println(xbytes.AnyEmpty([]byte("alpha"), []byte("beta")))
```

Output:

```text
true
false
```

</details>

<a name="AnyNonEmpty"></a>

## func [AnyNonEmpty](<https://github.com/gechr/x/blob/main/bytes/blank.go#L21>)

```go
func AnyNonEmpty(values ...[]byte) bool
```

**AnyNonEmpty** reports whether any of the given slices is non-empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.AnyNonEmpty([]byte(""), []byte("alpha"), []byte("")))
fmt.Println(xbytes.AnyNonEmpty([]byte(""), []byte("")))
```

Output:

```text
true
false
```

</details>

<a name="CompareFold"></a>

## func [CompareFold](<https://github.com/gechr/x/blob/main/bytes/fold.go#L9>)

```go
func CompareFold(a, b []byte) int
```

**CompareFold** compares `a` and `b` case-insensitively, using the same simple case-folding as [bytes.EqualFold](<https://pkg.go.dev/bytes#EqualFold>), and returns -1, 0, or 1 following the [cmp.Compare](<https://pkg.go.dev/cmp#Compare>) convention. `CompareFold(a, b) == 0` iff `bytes.EqualFold(a, b)`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.CompareFold([]byte("Go"), []byte("go")))
fmt.Println(xbytes.CompareFold([]byte("abc"), []byte("ABD")))
fmt.Println(xbytes.CompareFold([]byte("B"), []byte("a")))
```

Output:

```text
0
-1
1
```

</details>

<a name="ContainsAll"></a>

## func [ContainsAll](<https://github.com/gechr/x/blob/main/bytes/contains.go#L6>)

```go
func ContainsAll(s []byte, subslices ...[]byte) bool
```

**ContainsAll** reports whether `s` contains all of the given `subslices`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.ContainsAll([]byte("hello world"), []byte("hello"), []byte("world")))
fmt.Println(xbytes.ContainsAll([]byte("hello world"), []byte("hello"), []byte("moon")))
```

Output:

```text
true
false
```

</details>

<a name="ContainsAny"></a>

## func [ContainsAny](<https://github.com/gechr/x/blob/main/bytes/contains.go#L16>)

```go
func ContainsAny(s []byte, subslices ...[]byte) bool
```

**ContainsAny** reports whether `s` contains any of the given `subslices`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.ContainsAny([]byte("hello world"), []byte("moon"), []byte("world")))
fmt.Println(xbytes.ContainsAny([]byte("hello world"), []byte("moon"), []byte("sun")))
```

Output:

```text
true
false
```

</details>

<a name="ContainsFold"></a>

## func [ContainsFold](<https://github.com/gechr/x/blob/main/bytes/fold.go#L15>)

```go
func ContainsFold(s, subslice []byte) bool
```

**ContainsFold** reports whether `s` contains `subslice`, case-insensitively using the same simple case-folding as [bytes.EqualFold](<https://pkg.go.dev/bytes#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.ContainsFold([]byte("Hello, World"), []byte("WORLD")))
fmt.Println(xbytes.ContainsFold([]byte("Hello, World"), []byte("moon")))
```

Output:

```text
true
false
```

</details>

<a name="CountAny"></a>

## func [CountAny](<https://github.com/gechr/x/blob/main/bytes/any.go#L35>)

```go
func CountAny(s []byte, chars string) int
```

**CountAny** returns the number of Unicode code points in `s` that are contained in `chars`, following the cutset convention of [bytes.IndexAny](<https://pkg.go.dev/bytes#IndexAny>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.CountAny([]byte("hello world"), "lo"))
```

Output:

```text
5
```

</details>

<a name="DecodeSHA256"></a>

## func [DecodeSHA256](<https://github.com/gechr/x/blob/main/bytes/hex.go#L62>)

```go
func DecodeSHA256(s []byte) ([sha256.Size]byte, error)
```

**DecodeSHA256** decodes a 64-digit hexadecimal sha256 digest. It returns the zero digest and an error if `s` has the wrong length or contains a non-hexadecimal character.

<details><summary><b>Example</b></summary>

```go
digest, err := xbytes.DecodeSHA256(
    []byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
)
fmt.Printf("%x\n", digest[:4])
fmt.Println(err)
```

Output:

```text
e3b0c442
<nil>
```

</details>

<a name="Dedent"></a>

## func [Dedent](<https://github.com/gechr/x/blob/main/bytes/indent.go#L36>)

```go
func Dedent(s []byte) []byte
```

**Dedent** strips the longest common leading-whitespace prefix from non-empty lines. Whitespace-only lines are normalized to empty (Python textwrap.dedent) and CRLF line endings to LF.

```go
Dedent([]byte("    foo\n      bar\n    baz")) // "foo\n  bar\nbaz"
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%s\n", xbytes.Dedent([]byte("    foo\n      bar\n    baz")))
```

Output:

```text
foo
  bar
baz
```

</details>

<a name="EnsureTrailingNewline"></a>

## func [EnsureTrailingNewline](<https://github.com/gechr/x/blob/main/bytes/newline.go#L8>)

```go
func EnsureTrailingNewline(s []byte) []byte
```

**EnsureTrailingNewline** trims any trailing newlines from `s` and appends exactly one, so the result always ends in a single `\n`. An empty slice becomes `\n`. The returned slice never aliases `s`.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.EnsureTrailingNewline([]byte("hello\n\n")))
fmt.Printf("%q\n", xbytes.EnsureTrailingNewline([]byte("hello")))
```

Output:

```text
"hello\n"
"hello\n"
```

</details>

<a name="HasPrefixFold"></a>

## func [HasPrefixFold](<https://github.com/gechr/x/blob/main/bytes/fold.go#L21>)

```go
func HasPrefixFold(s, prefix []byte) bool
```

**HasPrefixFold** reports whether `s` begins with `prefix`, case-insensitively using the same simple case-folding as [bytes.EqualFold](<https://pkg.go.dev/bytes#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.HasPrefixFold([]byte("Hello, World"), []byte("HELLO")))
fmt.Println(xbytes.HasPrefixFold([]byte("Hello, World"), []byte("world")))
```

Output:

```text
true
false
```

</details>

<a name="HasSuffixFold"></a>

## func [HasSuffixFold](<https://github.com/gechr/x/blob/main/bytes/fold.go#L27>)

```go
func HasSuffixFold(s, suffix []byte) bool
```

**HasSuffixFold** reports whether `s` ends with `suffix`, case-insensitively using the same simple case-folding as [bytes.EqualFold](<https://pkg.go.dev/bytes#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.HasSuffixFold([]byte("Hello, World"), []byte("WORLD")))
fmt.Println(xbytes.HasSuffixFold([]byte("Hello, World"), []byte("hello")))
```

Output:

```text
true
false
```

</details>

<a name="HexEqual"></a>

## func [HexEqual](<https://github.com/gechr/x/blob/main/bytes/hex.go#L15>)

```go
func HexEqual(a, b []byte) bool
```

**HexEqual** reports whether `a` and `b` denote the same hexadecimal value, ignoring surrounding whitespace, an optional `0x` (or `0X`) prefix, and case. Two blank slices are equal; a blank slice never equals a non-blank one.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.HexEqual([]byte("0xDEADbeef"), []byte("deadbeef")))
fmt.Println(xbytes.HexEqual([]byte("0x1234"), []byte("0x5678")))
```

Output:

```text
true
false
```

</details>

<a name="Indent"></a>

## func [Indent](<https://github.com/gechr/x/blob/main/bytes/indent.go#L11>)

```go
func Indent(s, prefix []byte) []byte
```

**Indent** prefixes every non-blank line of `s` with `prefix`. Blank and whitespace-only lines are normalized to empty, and CRLF line endings to LF.

```go
Indent([]byte("foo\nbar"), []byte("  "))      // "  foo\n  bar"
Indent([]byte("foo\n\nbar"), []byte("> "))    // "> foo\n\n> bar"
Indent([]byte("foo\n   \nbar"), []byte("> ")) // "> foo\n\n> bar"
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%s\n", xbytes.Indent([]byte("foo\nbar"), []byte("> ")))
```

Output:

```text
> foo
> bar
```

</details>

<a name="IsASCII"></a>

## func [IsASCII](<https://github.com/gechr/x/blob/main/bytes/alpha.go#L7>)

```go
func IsASCII(s []byte) bool
```

**IsASCII** reports whether `s` is non-empty and consists entirely of ASCII characters (code points 0-127). An empty slice is not ASCII.

<a name="IsAlpha"></a>

## func [IsAlpha](<https://github.com/gechr/x/blob/main/bytes/alpha.go#L21>)

```go
func IsAlpha(s []byte) bool
```

**IsAlpha** reports whether `s` is non-empty and consists entirely of ASCII letters (a-z, A-Z). An empty slice is not alpha.

<a name="IsAlphaChar"></a>

## func [IsAlphaChar](<https://github.com/gechr/x/blob/main/bytes/alpha.go#L34>)

```go
func IsAlphaChar(c byte) bool
```

**IsAlphaChar** reports whether `c` is an ASCII letter (a-z, A-Z).

<a name="IsAlphanumeric"></a>

## func [IsAlphanumeric](<https://github.com/gechr/x/blob/main/bytes/alpha.go#L41>)

```go
func IsAlphanumeric(s []byte) bool
```

**IsAlphanumeric** reports whether `s` is non-empty and consists entirely of ASCII letters (a-z, A-Z) or digits (0-9). An empty slice is not alphanumeric.

<a name="IsAlphanumericChar"></a>

## func [IsAlphanumericChar](<https://github.com/gechr/x/blob/main/bytes/alpha.go#L55>)

```go
func IsAlphanumericChar(c byte) bool
```

**IsAlphanumericChar** reports whether `c` is an ASCII letter (a-z, A-Z) or digit (0-9).

<a name="IsBlank"></a>

## func [IsBlank](<https://github.com/gechr/x/blob/main/bytes/blank.go#L6>)

```go
func IsBlank(s []byte) bool
```

**IsBlank** reports whether `s` is empty or consists only of whitespace.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.IsBlank([]byte(" \t\n")))
fmt.Println(xbytes.IsBlank([]byte("x")))
```

Output:

```text
true
false
```

</details>

<a name="IsDigitChar"></a>

## func [IsDigitChar](<https://github.com/gechr/x/blob/main/bytes/digits.go#L18>)

```go
func IsDigitChar(c byte) bool
```

**IsDigitChar** reports whether `c` is an ASCII digit (0-9).

<a name="IsDigits"></a>

## func [IsDigits](<https://github.com/gechr/x/blob/main/bytes/digits.go#L5>)

```go
func IsDigits(s []byte) bool
```

**IsDigits** reports whether `s` is non-empty and consists entirely of ASCII digits (0-9). An empty slice is not digits.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.IsDigits([]byte("12345")))
fmt.Println(xbytes.IsDigits([]byte("12a45")))
fmt.Println(xbytes.IsDigits([]byte("")))
```

Output:

```text
true
false
false
```

</details>

<a name="IsGitCommit"></a>

## func [IsGitCommit](<https://github.com/gechr/x/blob/main/bytes/hex.go#L50>)

```go
func IsGitCommit(s []byte) bool
```

**IsGitCommit** reports whether `s` is 40 hexadecimal digits (a Git commit hash).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.IsGitCommit([]byte("3b18e512dba79e4c8300dd08aeb37f8e728b8dad")))
fmt.Println(xbytes.IsGitCommit([]byte("deadbeef")))
```

Output:

```text
true
false
```

</details>

<a name="IsHex"></a>

## func [IsHex](<https://github.com/gechr/x/blob/main/bytes/hex.go#L32>)

```go
func IsHex(s []byte) bool
```

**IsHex** reports whether `s` is non-empty and consists entirely of hexadecimal digits. An empty slice is not hex.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.IsHex([]byte("deadBEEF42")))
fmt.Println(xbytes.IsHex([]byte("xyz")))
```

Output:

```text
true
false
```

</details>

<a name="IsHexChar"></a>

## func [IsHexChar](<https://github.com/gechr/x/blob/main/bytes/hex.go#L45>)

```go
func IsHexChar(c byte) bool
```

**IsHexChar** reports whether `c` is a valid hexadecimal digit (0-9, a-f, A-F).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xbytes.IsHexChar('f'))
fmt.Println(xbytes.IsHexChar('F'))
fmt.Println(xbytes.IsHexChar('9'))
fmt.Println(xbytes.IsHexChar('g'))
```

Output:

```text
true
true
true
false
```

</details>

<a name="IsSHA256"></a>

## func [IsSHA256](<https://github.com/gechr/x/blob/main/bytes/hex.go#L55>)

```go
func IsSHA256(s []byte) bool
```

**IsSHA256** reports whether `s` is 64 hexadecimal digits (a sha256 digest).

<details><summary><b>Example</b></summary>

```go
fmt.Println(
    xbytes.IsSHA256([]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")),
)
fmt.Println(xbytes.IsSHA256([]byte("deadbeef")))
```

Output:

```text
true
false
```

</details>

<a name="PadCenter"></a>

## func [PadCenter](<https://github.com/gechr/x/blob/main/bytes/pad.go#L38>)

```go
func PadCenter(s []byte, width int) []byte
```

**PadCenter** pads `s` with spaces on both sides to `width` runes, centring it. An odd rune of padding goes on the right. Slices already `width` runes or longer are returned unchanged. The returned slice never aliases `s`.

```go
PadCenter([]byte("hi"), 5) // " hi  "
```

<details><summary><b>Example</b></summary>

**PadCenter** places the odd rune of padding on the right.

```go
fmt.Printf("%q\n", xbytes.PadCenter([]byte("hi"), 5))
```

Output:

```text
" hi  "
```

</details>

<a name="PadLeft"></a>

## func [PadLeft](<https://github.com/gechr/x/blob/main/bytes/pad.go#L14>)

```go
func PadLeft(s []byte, width int) []byte
```

**PadLeft** pads `s` with spaces on the left to `width` runes, right-aligning it. Slices already `width` runes or longer are returned unchanged. Width is counted in runes; for display-width-aware handling of ANSI text use the [ansi](<../ansi/README.md>) package. The returned slice never aliases `s`.

```go
PadLeft([]byte("hi"), 5) // "   hi"
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.PadLeft([]byte("hi"), 5))
```

Output:

```text
"   hi"
```

</details>

<a name="PadRight"></a>

## func [PadRight](<https://github.com/gechr/x/blob/main/bytes/pad.go#L26>)

```go
func PadRight(s []byte, width int) []byte
```

**PadRight** pads `s` with spaces on the right to `width` runes, left-aligning it. Slices already `width` runes or longer are returned unchanged. The returned slice never aliases `s`.

```go
PadRight([]byte("hi"), 5) // "hi   "
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.PadRight([]byte("hi"), 5))
```

Output:

```text
"hi   "
```

</details>

<a name="SplitAny"></a>

## func [SplitAny](<https://github.com/gechr/x/blob/main/bytes/any.go#L17>)

```go
func SplitAny(s []byte, chars string) [][]byte
```

**SplitAny** splits `s` around each occurrence of any Unicode code point in `chars`, following the cutset convention of [bytes.IndexAny](<https://pkg.go.dev/bytes#IndexAny>). Empty segments between adjacent separators are preserved, matching [bytes.Split](<https://pkg.go.dev/bytes#Split>) semantics. If `chars` is empty, [SplitAny](<#SplitAny>) returns a single-element slice containing `s`.

<details><summary><b>Example</b></summary>

Empty segments between adjacent separators are preserved.

```go
fmt.Printf("%q\n", xbytes.SplitAny([]byte("a,b;;c"), ",;"))
```

Output:

```text
["a" "b" "" "c"]
```

</details>

<a name="SplitBy"></a>

## func [SplitBy](<https://github.com/gechr/x/blob/main/bytes/split.go#L7>)

```go
func SplitBy(s, sep []byte) [][]byte
```

**SplitBy** splits `s` by `sep`, trims whitespace from each part, and drops empty values.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.SplitBy([]byte(" a | b || c "), []byte("|")))
```

Output:

```text
["a" "b" "c"]
```

</details>

<a name="SplitLines"></a>

## func [SplitLines](<https://github.com/gechr/x/blob/main/bytes/lines.go#L6>)

```go
func SplitLines(s []byte) [][]byte
```

**SplitLines** splits `s` into non-empty trimmed lines.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.SplitLines([]byte("foo\n\n  bar \n")))
```

Output:

```text
["foo" "bar"]
```

</details>

<a name="SplitLinesRaw"></a>

## func [SplitLinesRaw](<https://github.com/gechr/x/blob/main/bytes/lines.go#L13>)

```go
func SplitLinesRaw(s []byte) [][]byte
```

**SplitLinesRaw** splits `s` into lines losslessly, normalizing CRLF to LF: every line is kept verbatim - empty lines and the trailing empty element included - so the result joins back with `"\n"` without losing content or line numbers.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xbytes.SplitLinesRaw([]byte("foo\r\nbar\n")))
```

Output:

```text
["foo" "bar" ""]
```

</details>

<a name="TrimPrefixes"></a>

## func [TrimPrefixes](<https://github.com/gechr/x/blob/main/bytes/trim.go#L8>)

```go
func TrimPrefixes(s []byte, prefixes ...[]byte) []byte
```

**TrimPrefixes** returns `s` with the first matching prefix in `prefixes` removed. At most one prefix is removed; if none match, `s` is returned unchanged.

<a name="TrimSuffixes"></a>

## func [TrimSuffixes](<https://github.com/gechr/x/blob/main/bytes/trim.go#L20>)

```go
func TrimSuffixes(s []byte, suffixes ...[]byte) []byte
```

**TrimSuffixes** returns `s` with the first matching suffix in `suffixes` removed. At most one suffix is removed; if none match, `s` is returned unchanged.

<a name="Truncate"></a>

## func [Truncate](<https://github.com/gechr/x/blob/main/bytes/truncate.go#L50>)

```go
func Truncate(s []byte, n int, marker []byte) []byte
```

**Truncate** is an alias for [TruncateRight](<#TruncateRight>), the most common form: it keeps the head and trims the tail.

<a name="TruncateLeft"></a>

## func [TruncateLeft](<https://github.com/gechr/x/blob/main/bytes/truncate.go#L26>)

```go
func TruncateLeft(s []byte, n int, marker []byte) []byte
```

**TruncateLeft** shortens `s` to at most `n` runes (including `marker`) by removing characters from the left, prepending `marker` when truncation occurs. The tail is kept.

```go
TruncateLeft([]byte("hello world"), 8, []byte("…")) // "…o world"
```

<a name="TruncateMiddle"></a>

## func [TruncateMiddle](<https://github.com/gechr/x/blob/main/bytes/truncate.go#L39>)

```go
func TruncateMiddle(s []byte, n int, marker []byte) []byte
```

**TruncateMiddle** shortens `s` to at most `n` runes (including `marker`) by removing characters from the middle, inserting `marker` between the kept head and tail so both ends stay visible. This suits hashes and paths, where the start and end are the recognisable parts.

```go
TruncateMiddle([]byte("0123456789abcdef"), 7, []byte("…")) // "012…def"
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%s\n", xbytes.TruncateMiddle([]byte("0123456789abcdef"), 7, []byte("…")))
```

Output:

```text
012…def
```

</details>

<a name="TruncateRight"></a>

## func [TruncateRight](<https://github.com/gechr/x/blob/main/bytes/truncate.go#L15>)

```go
func TruncateRight(s []byte, n int, marker []byte) []byte
```

**TruncateRight** shortens `s` to at most `n` runes (including `marker`) by removing characters from the right, appending `marker` when truncation occurs. The head is kept. For display-width-aware truncation of ANSI text use [ansi.Truncate](<../ansi/README.md#Truncate>).

```go
TruncateRight([]byte("hello world"), 8, []byte("…")) // "hello w…"
TruncateRight([]byte("hi"), 8, []byte("…"))          // "hi"
```

<a name="Unwrap"></a>

## func [Unwrap](<https://github.com/gechr/x/blob/main/bytes/unwrap.go#L9>)

```go
func Unwrap(s, prefix, suffix []byte) ([]byte, bool)
```

**Unwrap** returns `s` with the leading `prefix` and trailing `suffix` removed and reports whether both were present. Unlike a [bytes.TrimPrefix](<https://pkg.go.dev/bytes#TrimPrefix>) + [bytes.TrimSuffix](<https://pkg.go.dev/bytes#TrimSuffix>) chain, nothing is removed unless `s` starts with `prefix` AND ends with `suffix`, so a one-sided match is returned unchanged.

<details><summary><b>Example</b></summary>

```go
quoted, ok := xbytes.Unwrap([]byte(`"quoted"`), []byte(`"`), []byte(`"`))
fmt.Println(string(quoted), ok)
oneSided, ok := xbytes.Unwrap([]byte(`"one-sided`), []byte(`"`), []byte(`"`))
fmt.Println(string(oneSided), ok)
```

Output:

```text
quoted true
"one-sided false
```

</details>
