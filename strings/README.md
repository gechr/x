# strings

```go
import "github.com/gechr/x/strings"
```

Package `strings` provides string helpers: split, contains, indent/dedent, truncate, and blank checks.

## Index

- [func AllEmpty(values ...string) bool](<#AllEmpty>)
- [func AllNonEmpty(values ...string) bool](<#AllNonEmpty>)
- [func AnyEmpty(values ...string) bool](<#AnyEmpty>)
- [func AnyNonEmpty(values ...string) bool](<#AnyNonEmpty>)
- [func AppendCSV(dst \[\]string, raw string) \[\]string](<#AppendCSV>)
- [func Closest(target string, candidates \[\]string) string](<#Closest>)
- [func CompactLines(s, sep string) string](<#CompactLines>)
- [func CompareFold(a, b string) int](<#CompareFold>)
- [func CompareNatural(a, b string) int](<#CompareNatural>)
- [func ContainsAll(s string, substrings ...string) bool](<#ContainsAll>)
- [func ContainsAny(s string, substrings ...string) bool](<#ContainsAny>)
- [func ContainsFold(s, substr string) bool](<#ContainsFold>)
- [func CountAny(s, chars string) int](<#CountAny>)
- [func Dedent(s string) string](<#Dedent>)
- [func EnsureTrailingNewline(s string) string](<#EnsureTrailingNewline>)
- [func EqualNatural(a, b string) bool](<#EqualNatural>)
- [func HasPrefixFold(s, prefix string) bool](<#HasPrefixFold>)
- [func HasSuffixFold(s, suffix string) bool](<#HasSuffixFold>)
- [func HexEqual(a, b string) bool](<#HexEqual>)
- [func Indent(s, prefix string) string](<#Indent>)
- [func IsASCII(s string) bool](<#IsASCII>)
- [func IsAlpha(s string) bool](<#IsAlpha>)
- [func IsAlphaChar(c rune) bool](<#IsAlphaChar>)
- [func IsAlphanumeric(s string) bool](<#IsAlphanumeric>)
- [func IsAlphanumericChar(c rune) bool](<#IsAlphanumericChar>)
- [func IsBlank(s string) bool](<#IsBlank>)
- [func IsDigitChar(c rune) bool](<#IsDigitChar>)
- [func IsDigits(s string) bool](<#IsDigits>)
- [func IsGitCommit(s string) bool](<#IsGitCommit>)
- [func IsHex(s string) bool](<#IsHex>)
- [func IsHexChar(c rune) bool](<#IsHexChar>)
- [func IsSHA256(s string) bool](<#IsSHA256>)
- [func IsSlug(s string) bool](<#IsSlug>)
- [func IsSlugLenient(s string) bool](<#IsSlugLenient>)
- [func LessNatural(a, b string) bool](<#LessNatural>)
- [func PadCenter(s string, width int) string](<#PadCenter>)
- [func PadLeft(s string, width int) string](<#PadLeft>)
- [func PadRight(s string, width int) string](<#PadRight>)
- [func SplitAny(s, chars string) \[\]string](<#SplitAny>)
- [func SplitBy(s, sep string) \[\]string](<#SplitBy>)
- [func SplitCSV(s string) \[\]string](<#SplitCSV>)
- [func SplitLines(s string) \[\]string](<#SplitLines>)
- [func SplitLinesRaw(s string) \[\]string](<#SplitLinesRaw>)
- [func TrimPrefixes(s string, prefixes ...string) string](<#TrimPrefixes>)
- [func TrimSuffixes(s string, suffixes ...string) string](<#TrimSuffixes>)
- [func Truncate(s string, n int, marker string) string](<#Truncate>)
- [func TruncateLeft(s string, n int, marker string) string](<#TruncateLeft>)
- [func TruncateMiddle(s string, n int, marker string) string](<#TruncateMiddle>)
- [func TruncateRight(s string, n int, marker string) string](<#TruncateRight>)
- [func Unwrap(s, prefix, suffix string) (string, bool)](<#Unwrap>)

<a name="AllEmpty"></a>

## func [AllEmpty](<https://github.com/gechr/x/blob/main/strings/blank.go#L29>)

```go
func AllEmpty(values ...string) bool
```

**AllEmpty** reports whether every given string is empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.AllEmpty("", ""))
fmt.Println(xstrings.AllEmpty("", "alpha"))
```

Output:

```text
true
false
```

</details>

<a name="AllNonEmpty"></a>

## func [AllNonEmpty](<https://github.com/gechr/x/blob/main/strings/blank.go#L39>)

```go
func AllNonEmpty(values ...string) bool
```

**AllNonEmpty** reports whether every given string is non-empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.AllNonEmpty("alpha", "beta", "charlie"))
fmt.Println(xstrings.AllNonEmpty("alpha", ""))
```

Output:

```text
true
false
```

</details>

<a name="AnyEmpty"></a>

## func [AnyEmpty](<https://github.com/gechr/x/blob/main/strings/blank.go#L14>)

```go
func AnyEmpty(values ...string) bool
```

**AnyEmpty** reports whether any of the given strings is empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.AnyEmpty("alpha", "", "beta"))
fmt.Println(xstrings.AnyEmpty("alpha", "beta"))
```

Output:

```text
true
false
```

</details>

<a name="AnyNonEmpty"></a>

## func [AnyNonEmpty](<https://github.com/gechr/x/blob/main/strings/blank.go#L19>)

```go
func AnyNonEmpty(values ...string) bool
```

**AnyNonEmpty** reports whether any of the given strings is non-empty.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.AnyNonEmpty("", "alpha", ""))
fmt.Println(xstrings.AnyNonEmpty("", ""))
```

Output:

```text
true
false
```

</details>

<a name="AppendCSV"></a>

## func [AppendCSV](<https://github.com/gechr/x/blob/main/strings/csv.go#L5>)

```go
func AppendCSV(dst []string, raw string) []string
```

**AppendCSV** splits `raw` on commas, trims whitespace, drops empty values, and appends the remaining values to `dst`.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.AppendCSV([]string{"x"}, " a, b ,, c "))
```

Output:

```text
["x" "a" "b" "c"]
```

</details>

<a name="Closest"></a>

## func [Closest](<https://github.com/gechr/x/blob/main/strings/closest.go#L18>)

```go
func Closest(target string, candidates []string) string
```

**Closest** returns the candidate nearest to `target`, suitable for a "did you mean?" suggestion. Distance is the Damerau-Levenshtein (optimal string alignment) edit distance, so an adjacent transposition like `verfiy` counts as one edit, not two - the common typo plain Levenshtein over-penalizes. It returns `""` when the nearest candidate is further than a third of `target`'s length in edits, so an unrelated word is never suggested. An empty `target` carries no signal and suggests nothing. Ties resolve to the first candidate.

```go
Closest("verfiy", []string{"verify", "deep"}) // "verify"
Closest("xyzzy", []string{"verify", "deep"})  // ""
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.Closest("verfiy", []string{"verify", "deep"}))
fmt.Printf("%q\n", xstrings.Closest("xyzzy", []string{"verify", "deep"}))
```

Output:

```text
"verify"
""
```

</details>

<a name="CompactLines"></a>

## func [CompactLines](<https://github.com/gechr/x/blob/main/strings/compact.go#L7>)

```go
func CompactLines(s, sep string) string
```

**CompactLines** trims lines, drops blank lines, removes duplicate lines while preserving first-seen order, and joins the remaining lines with `sep`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.CompactLines("  foo \n\nbar\nfoo\n", ", "))
```

Output:

```text
foo, bar
```

</details>

<a name="CompareFold"></a>

## func [CompareFold](<https://github.com/gechr/x/blob/main/strings/fold.go#L14>)

```go
func CompareFold(a, b string) int
```

**CompareFold** compares `a` and `b` case-insensitively, using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>), and returns -1, 0, or 1 following the [cmp.Compare](<https://pkg.go.dev/cmp#Compare>) convention. `CompareFold(a, b) == 0` iff `strings.EqualFold(a, b)`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.CompareFold("Go", "go"))
fmt.Println(xstrings.CompareFold("abc", "ABD"))
fmt.Println(xstrings.CompareFold("B", "a"))
```

Output:

```text
0
-1
1
```

</details>

<a name="CompareNatural"></a>

## func [CompareNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L12>)

```go
func CompareNatural(a, b string) int
```

**CompareNatural** orders `a` and `b` the way a human reads them, treating each run of digits as a single decimal number so `x2` sorts before `x10`. It returns -1, 0, or +1 and allocates nothing, handling numbers of any length without overflow.

<details><summary><b>Example</b></summary>

```go
versions := []string{"v10", "v2", "v1"}
slices.SortFunc(versions, xstrings.CompareNatural)
fmt.Println(versions)
```

Output:

```text
[v1 v2 v10]
```

</details>

<a name="ContainsAll"></a>

## func [ContainsAll](<https://github.com/gechr/x/blob/main/strings/contains.go#L6>)

```go
func ContainsAll(s string, substrings ...string) bool
```

**ContainsAll** reports whether `s` contains all of the given `substrings`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.ContainsAll("hello world", "hello", "world"))
fmt.Println(xstrings.ContainsAll("hello world", "hello", "moon"))
```

Output:

```text
true
false
```

</details>

<a name="ContainsAny"></a>

## func [ContainsAny](<https://github.com/gechr/x/blob/main/strings/contains.go#L16>)

```go
func ContainsAny(s string, substrings ...string) bool
```

**ContainsAny** reports whether `s` contains any of the given `substrings`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.ContainsAny("hello world", "moon", "world"))
fmt.Println(xstrings.ContainsAny("hello world", "moon", "sun"))
```

Output:

```text
true
false
```

</details>

<a name="ContainsFold"></a>

## func [ContainsFold](<https://github.com/gechr/x/blob/main/strings/fold.go#L28>)

```go
func ContainsFold(s, substr string) bool
```

**ContainsFold** reports whether `s` contains `substr`, case-insensitively using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.ContainsFold("Hello, World", "WORLD"))
fmt.Println(xstrings.ContainsFold("Hello, World", "moon"))
```

Output:

```text
true
false
```

</details>

<a name="CountAny"></a>

## func [CountAny](<https://github.com/gechr/x/blob/main/strings/any.go#L32>)

```go
func CountAny(s, chars string) int
```

**CountAny** returns the number of Unicode code points in `s` that are contained in `chars`, following the cutset convention of [strings.IndexAny](<https://pkg.go.dev/strings#IndexAny>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.CountAny("hello world", "lo"))
```

Output:

```text
5
```

</details>

<a name="Dedent"></a>

## func [Dedent](<https://github.com/gechr/x/blob/main/strings/indent.go#L35>)

```go
func Dedent(s string) string
```

**Dedent** strips the longest common leading-whitespace prefix from non-empty lines. Whitespace-only lines are normalized to empty (Python textwrap.dedent).

```go
Dedent("    foo\n      bar\n    baz") // "foo\n  bar\nbaz"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.Dedent("    foo\n      bar\n    baz"))
```

Output:

```text
foo
  bar
baz
```

</details>

<a name="EnsureTrailingNewline"></a>

## func [EnsureTrailingNewline](<https://github.com/gechr/x/blob/main/strings/newline.go#L8>)

```go
func EnsureTrailingNewline(s string) string
```

**EnsureTrailingNewline** trims any trailing newlines from `s` and appends exactly one, so the result always ends in a single `\n`. An empty string becomes `\n`.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.EnsureTrailingNewline("hello\n\n"))
fmt.Printf("%q\n", xstrings.EnsureTrailingNewline("hello"))
```

Output:

```text
"hello\n"
"hello\n"
```

</details>

<a name="EqualNatural"></a>

## func [EqualNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L49>)

```go
func EqualNatural(a, b string) bool
```

**EqualNatural** reports whether `a` and `b` compare equal in natural order, as decided by [CompareNatural](<#CompareNatural>). This can differ from `a == b`, since a numeric run followed by more to compare matches regardless of leading zeros (for example `a00b00` and `a0b00`).

<details><summary><b>Example</b></summary>

Leading zeros are ignored when more text follows the numeric run.

```go
fmt.Println(xstrings.EqualNatural("a00b00", "a0b00"))
fmt.Println(xstrings.EqualNatural("a1", "a2"))
```

Output:

```text
true
false
```

</details>

<a name="HasPrefixFold"></a>

## func [HasPrefixFold](<https://github.com/gechr/x/blob/main/strings/fold.go#L44>)

```go
func HasPrefixFold(s, prefix string) bool
```

**HasPrefixFold** reports whether `s` begins with `prefix`, case-insensitively using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.HasPrefixFold("Hello, World", "HELLO"))
fmt.Println(xstrings.HasPrefixFold("Hello, World", "world"))
```

Output:

```text
true
false
```

</details>

<a name="HasSuffixFold"></a>

## func [HasSuffixFold](<https://github.com/gechr/x/blob/main/strings/fold.go#L51>)

```go
func HasSuffixFold(s, suffix string) bool
```

**HasSuffixFold** reports whether `s` ends with `suffix`, case-insensitively using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.HasSuffixFold("Hello, World", "WORLD"))
fmt.Println(xstrings.HasSuffixFold("Hello, World", "hello"))
```

Output:

```text
true
false
```

</details>

<a name="HexEqual"></a>

## func [HexEqual](<https://github.com/gechr/x/blob/main/strings/hex.go#L8>)

```go
func HexEqual(a, b string) bool
```

**HexEqual** reports whether `a` and `b` denote the same hexadecimal value, ignoring surrounding whitespace, an optional `0x` (or `0X`) prefix, and case. Two blank strings are equal; a blank string never equals a non-blank one.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.HexEqual("0xDEADbeef", "deadbeef"))
fmt.Println(xstrings.HexEqual("0x1234", "0x5678"))
```

Output:

```text
true
false
```

</details>

<a name="Indent"></a>

## func [Indent](<https://github.com/gechr/x/blob/main/strings/indent.go#L11>)

```go
func Indent(s, prefix string) string
```

**Indent** prefixes every non-blank line of `s` with `prefix`. Blank and whitespace-only lines are normalized to empty.

```go
Indent("foo\nbar", "  ")      // "  foo\n  bar"
Indent("foo\n\nbar", "> ")    // "> foo\n\n> bar"
Indent("foo\n   \nbar", "> ") // "> foo\n\n> bar"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.Indent("foo\nbar", "> "))
```

Output:

```text
> foo
> bar
```

</details>

<a name="IsASCII"></a>

## func [IsASCII](<https://github.com/gechr/x/blob/main/strings/alpha.go#L7>)

```go
func IsASCII(s string) bool
```

**IsASCII** reports whether `s` is non-empty and consists entirely of ASCII characters (code points 0-127). An empty string is not ASCII.

<a name="IsAlpha"></a>

## func [IsAlpha](<https://github.com/gechr/x/blob/main/strings/alpha.go#L21>)

```go
func IsAlpha(s string) bool
```

**IsAlpha** reports whether `s` is non-empty and consists entirely of ASCII letters (a-z, A-Z). An empty string is not alpha.

<a name="IsAlphaChar"></a>

## func [IsAlphaChar](<https://github.com/gechr/x/blob/main/strings/alpha.go#L34>)

```go
func IsAlphaChar(c rune) bool
```

**IsAlphaChar** reports whether `c` is an ASCII letter (a-z, A-Z).

<a name="IsAlphanumeric"></a>

## func [IsAlphanumeric](<https://github.com/gechr/x/blob/main/strings/alpha.go#L41>)

```go
func IsAlphanumeric(s string) bool
```

**IsAlphanumeric** reports whether `s` is non-empty and consists entirely of ASCII letters (a-z, A-Z) or digits (0-9). An empty string is not alphanumeric.

<a name="IsAlphanumericChar"></a>

## func [IsAlphanumericChar](<https://github.com/gechr/x/blob/main/strings/alpha.go#L55>)

```go
func IsAlphanumericChar(c rune) bool
```

**IsAlphanumericChar** reports whether `c` is an ASCII letter (a-z, A-Z) or digit (0-9).

<a name="IsBlank"></a>

## func [IsBlank](<https://github.com/gechr/x/blob/main/strings/blank.go#L9>)

```go
func IsBlank(s string) bool
```

**IsBlank** reports whether `s` is empty or consists only of whitespace.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.IsBlank(" \t\n"))
fmt.Println(xstrings.IsBlank("x"))
```

Output:

```text
true
false
```

</details>

<a name="IsDigitChar"></a>

## func [IsDigitChar](<https://github.com/gechr/x/blob/main/strings/digits.go#L18>)

```go
func IsDigitChar(c rune) bool
```

**IsDigitChar** reports whether `c` is an ASCII digit (0-9).

<a name="IsDigits"></a>

## func [IsDigits](<https://github.com/gechr/x/blob/main/strings/digits.go#L5>)

```go
func IsDigits(s string) bool
```

**IsDigits** reports whether `s` is non-empty and consists entirely of ASCII digits (0-9). An empty string is not digits.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.IsDigits("12345"))
fmt.Println(xstrings.IsDigits("12a45"))
fmt.Println(xstrings.IsDigits(""))
```

Output:

```text
true
false
false
```

</details>

<a name="IsGitCommit"></a>

## func [IsGitCommit](<https://github.com/gechr/x/blob/main/strings/hex.go#L43>)

```go
func IsGitCommit(s string) bool
```

**IsGitCommit** reports whether `s` is 40 hexadecimal digits (a Git commit hash).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.IsGitCommit("3b18e512dba79e4c8300dd08aeb37f8e728b8dad"))
fmt.Println(xstrings.IsGitCommit("deadbeef"))
```

Output:

```text
true
false
```

</details>

<a name="IsHex"></a>

## func [IsHex](<https://github.com/gechr/x/blob/main/strings/hex.go#L25>)

```go
func IsHex(s string) bool
```

**IsHex** reports whether `s` is non-empty and consists entirely of hexadecimal digits. An empty string is not hex.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.IsHex("deadBEEF42"))
fmt.Println(xstrings.IsHex("xyz"))
```

Output:

```text
true
false
```

</details>

<a name="IsHexChar"></a>

## func [IsHexChar](<https://github.com/gechr/x/blob/main/strings/hex.go#L38>)

```go
func IsHexChar(c rune) bool
```

**IsHexChar** reports whether `c` is a valid hexadecimal digit (0-9, a-f, A-F).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.IsHexChar('f'))
fmt.Println(xstrings.IsHexChar('F'))
fmt.Println(xstrings.IsHexChar('9'))
fmt.Println(xstrings.IsHexChar('g'))
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

## func [IsSHA256](<https://github.com/gechr/x/blob/main/strings/hex.go#L48>)

```go
func IsSHA256(s string) bool
```

**IsSHA256** reports whether `s` is 64 hexadecimal digits (a SHA-256 digest).

<details><summary><b>Example</b></summary>

```go
fmt.Println(
    xstrings.IsSHA256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"),
)
fmt.Println(xstrings.IsSHA256("deadbeef"))
```

Output:

```text
true
false
```

</details>

<a name="IsSlug"></a>

## func [IsSlug](<https://github.com/gechr/x/blob/main/strings/slug.go#L8>)

```go
func IsSlug(s string) bool
```

**IsSlug** reports whether `s` is a valid slug: a non-empty, URL-friendly identifier of lowercase alphanumerics and '-', starting and ending with an alphanumeric (e.g. `my-service`). Underscores are not permitted; `-` is the only allowed separator, and it may not appear consecutively. Every valid slug is therefore a fixed point of slugification. An empty string is not a slug.

<a name="IsSlugLenient"></a>

## func [IsSlugLenient](<https://github.com/gechr/x/blob/main/strings/slug.go#L34>)

```go
func IsSlugLenient(s string) bool
```

**IsSlugLenient** reports whether `s` is a valid lenient slug: a non-empty identifier of lowercase alphanumerics, '-', and '\_', starting and ending with an alphanumeric (e.g. `my-service`, `my_service`, `a--b__c`). Unlike [IsSlug](<#IsSlug>), underscores are permitted and separators may appear consecutively or mixed; only leading and trailing separators are rejected. An empty string is not a slug.

<a name="LessNatural"></a>

## func [LessNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L41>)

```go
func LessNatural(a, b string) bool
```

**LessNatural** reports whether `a` sorts before `b` in natural order, as decided by [CompareNatural](<#CompareNatural>). It reads cleanly at call sites that want a boolean rather than a three-way result, such as sort predicates and conditionals.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.LessNatural("v2", "v10"))
fmt.Println(xstrings.LessNatural("v10", "v2"))
```

Output:

```text
true
false
```

</details>

<a name="PadCenter"></a>

## func [PadCenter](<https://github.com/gechr/x/blob/main/strings/pad.go#L31>)

```go
func PadCenter(s string, width int) string
```

**PadCenter** pads `s` with spaces on both sides to `width` runes, centring it. An odd rune of padding goes on the right. Strings already `width` runes or longer are returned unchanged.

```go
PadCenter("hi", 5) // " hi  "
```

<details><summary><b>Example</b></summary>

**PadCenter** places the odd rune of padding on the right.

```go
fmt.Printf("%q\n", xstrings.PadCenter("hi", 5))
```

Output:

```text
" hi  "
```

</details>

<a name="PadLeft"></a>

## func [PadLeft](<https://github.com/gechr/x/blob/main/strings/pad.go#L14>)

```go
func PadLeft(s string, width int) string
```

**PadLeft** pads `s` with spaces on the left to `width` runes, right-aligning it. Strings already `width` runes or longer are returned unchanged. Width is counted in runes; for display-width-aware handling of ANSI text use the [ansi](<../ansi/README.md>) package.

```go
PadLeft("hi", 5) // "   hi"
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.PadLeft("hi", 5))
```

Output:

```text
"   hi"
```

</details>

<a name="PadRight"></a>

## func [PadRight](<https://github.com/gechr/x/blob/main/strings/pad.go#L22>)

```go
func PadRight(s string, width int) string
```

**PadRight** pads `s` with spaces on the right to `width` runes, left-aligning it. Strings already `width` runes or longer are returned unchanged.

```go
PadRight("hi", 5) // "hi   "
```

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.PadRight("hi", 5))
```

Output:

```text
"hi   "
```

</details>

<a name="SplitAny"></a>

## func [SplitAny](<https://github.com/gechr/x/blob/main/strings/any.go#L14>)

```go
func SplitAny(s, chars string) []string
```

**SplitAny** splits `s` around each occurrence of any Unicode code point in `chars`, following the cutset convention of [strings.IndexAny](<https://pkg.go.dev/strings#IndexAny>). Empty segments between adjacent separators are preserved, matching [strings.Split](<https://pkg.go.dev/strings#Split>) semantics. If `chars` is empty, [SplitAny](<#SplitAny>) returns a single-element slice containing `s`.

<details><summary><b>Example</b></summary>

Empty segments between adjacent separators are preserved.

```go
fmt.Printf("%q\n", xstrings.SplitAny("a,b;;c", ",;"))
```

Output:

```text
["a" "b" "" "c"]
```

</details>

<a name="SplitBy"></a>

## func [SplitBy](<https://github.com/gechr/x/blob/main/strings/split.go#L7>)

```go
func SplitBy(s, sep string) []string
```

**SplitBy** splits `s` by `sep`, trims whitespace from each part, and drops empty values.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.SplitBy(" a | b || c ", "|"))
```

Output:

```text
["a" "b" "c"]
```

</details>

<a name="SplitCSV"></a>

## func [SplitCSV](<https://github.com/gechr/x/blob/main/strings/csv.go#L10>)

```go
func SplitCSV(s string) []string
```

**SplitCSV** splits `s` on commas, trims whitespace, and drops empty values.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.SplitCSV(" a, b ,, c "))
```

Output:

```text
["a" "b" "c"]
```

</details>

<a name="SplitLines"></a>

## func [SplitLines](<https://github.com/gechr/x/blob/main/strings/lines.go#L6>)

```go
func SplitLines(s string) []string
```

**SplitLines** splits `s` into non-empty trimmed lines.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.SplitLines("foo\n\n  bar \n"))
```

Output:

```text
["foo" "bar"]
```

</details>

<a name="SplitLinesRaw"></a>

## func [SplitLinesRaw](<https://github.com/gechr/x/blob/main/strings/lines.go#L13>)

```go
func SplitLinesRaw(s string) []string
```

**SplitLinesRaw** splits `s` into lines losslessly, normalizing CRLF to LF: every line is kept verbatim - empty lines and the trailing empty element included - so the result joins back with `"\n"` without losing content or line numbers.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%q\n", xstrings.SplitLinesRaw("foo\r\nbar\n"))
```

Output:

```text
["foo" "bar" ""]
```

</details>

<a name="TrimPrefixes"></a>

## func [TrimPrefixes](<https://github.com/gechr/x/blob/main/strings/trim.go#L8>)

```go
func TrimPrefixes(s string, prefixes ...string) string
```

**TrimPrefixes** returns `s` with the first matching prefix in `prefixes` removed. At most one prefix is removed; if none match, `s` is returned unchanged.

<a name="TrimSuffixes"></a>

## func [TrimSuffixes](<https://github.com/gechr/x/blob/main/strings/trim.go#L20>)

```go
func TrimSuffixes(s string, suffixes ...string) string
```

**TrimSuffixes** returns `s` with the first matching suffix in `suffixes` removed. At most one suffix is removed; if none match, `s` is returned unchanged.

<a name="Truncate"></a>

## func [Truncate](<https://github.com/gechr/x/blob/main/strings/truncate.go#L45>)

```go
func Truncate(s string, n int, marker string) string
```

**Truncate** is an alias for [TruncateRight](<#TruncateRight>), the most common form: it keeps the head and trims the tail.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.Truncate("hello world", 8, "…"))
```

Output:

```text
hello w…
```

</details>

<a name="TruncateLeft"></a>

## func [TruncateLeft](<https://github.com/gechr/x/blob/main/strings/truncate.go#L23>)

```go
func TruncateLeft(s string, n int, marker string) string
```

**TruncateLeft** shortens `s` to at most `n` runes (including `marker`) by removing characters from the left, prepending `marker` when truncation occurs. The tail is kept.

```go
TruncateLeft("hello world", 8, "…") // "…o world"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.TruncateLeft("hello world", 8, "…"))
```

Output:

```text
…o world
```

</details>

<a name="TruncateMiddle"></a>

## func [TruncateMiddle](<https://github.com/gechr/x/blob/main/strings/truncate.go#L35>)

```go
func TruncateMiddle(s string, n int, marker string) string
```

**TruncateMiddle** shortens `s` to at most `n` runes (including `marker`) by removing characters from the middle, inserting `marker` between the kept head and tail so both ends stay visible. This suits hashes and paths, where the start and end are the recognisable parts.

```go
TruncateMiddle("0123456789abcdef", 7, "…") // "012…def"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.TruncateMiddle("0123456789abcdef", 7, "…"))
```

Output:

```text
012…def
```

</details>

<a name="TruncateRight"></a>

## func [TruncateRight](<https://github.com/gechr/x/blob/main/strings/truncate.go#L12>)

```go
func TruncateRight(s string, n int, marker string) string
```

**TruncateRight** shortens `s` to at most `n` runes (including `marker`) by removing characters from the right, appending `marker` when truncation occurs. The head is kept. For display-width-aware truncation of ANSI text use [ansi.Truncate](<../ansi/README.md#Truncate>).

```go
TruncateRight("hello world", 8, "…") // "hello w…"
TruncateRight("hi", 8, "…")          // "hi"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.TruncateRight("hello world", 8, "…"))
fmt.Println(xstrings.TruncateRight("hi", 8, "…"))
```

Output:

```text
hello w…
hi
```

</details>

<a name="Unwrap"></a>

## func [Unwrap](<https://github.com/gechr/x/blob/main/strings/unwrap.go#L9>)

```go
func Unwrap(s, prefix, suffix string) (string, bool)
```

**Unwrap** returns `s` with the leading `prefix` and trailing `suffix` removed and reports whether both were present. Unlike a [strings.TrimPrefix](<https://pkg.go.dev/strings#TrimPrefix>) + [strings.TrimSuffix](<https://pkg.go.dev/strings#TrimSuffix>) chain, nothing is removed unless `s` starts with `prefix` AND ends with `suffix`, so a one-sided match is returned unchanged.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xstrings.Unwrap(`"quoted"`, `"`, `"`))
fmt.Println(xstrings.Unwrap(`"one-sided`, `"`, `"`))
```

Output:

```text
quoted true
"one-sided false
```

</details>
