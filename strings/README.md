# strings

```go
import "github.com/gechr/x/strings"
```

Package strings provides string helpers: split, contains, indent/dedent, truncate, and blank checks.

## Index

- [func AppendCSV\(dst \[\]string, raw string\) \[\]string](<#AppendCSV>)
- [func Closest\(target string, candidates \[\]string\) string](<#Closest>)
- [func CompactLines\(s, sep string\) string](<#CompactLines>)
- [func CompareFold\(a, b string\) int](<#CompareFold>)
- [func CompareNatural\(a, b string\) int](<#CompareNatural>)
- [func ContainsAll\(s string, substrings ...string\) bool](<#ContainsAll>)
- [func ContainsAny\(s string, substrings ...string\) bool](<#ContainsAny>)
- [func CountAny\(s, chars string\) int](<#CountAny>)
- [func Dedent\(s string\) string](<#Dedent>)
- [func EqualNatural\(a, b string\) bool](<#EqualNatural>)
- [func Indent\(s, prefix string\) string](<#Indent>)
- [func IsBlank\(s string\) bool](<#IsBlank>)
- [func IsGitCommit\(s string\) bool](<#IsGitCommit>)
- [func IsHex\(s string\) bool](<#IsHex>)
- [func IsHexChar\(c rune\) bool](<#IsHexChar>)
- [func IsSHA256\(s string\) bool](<#IsSHA256>)
- [func LessNatural\(a, b string\) bool](<#LessNatural>)
- [func SplitAny\(s, chars string\) \[\]string](<#SplitAny>)
- [func SplitBy\(s, sep string\) \[\]string](<#SplitBy>)
- [func SplitCSV\(s string\) \[\]string](<#SplitCSV>)
- [func SplitLines\(s string\) \[\]string](<#SplitLines>)
- [func SplitLinesRaw\(s string\) \[\]string](<#SplitLinesRaw>)
- [func Truncate\(s string, n int, marker string\) string](<#Truncate>)
- [func TruncateLeft\(s string, n int, marker string\) string](<#TruncateLeft>)
- [func TruncateMiddle\(s string, n int, marker string\) string](<#TruncateMiddle>)
- [func TruncateRight\(s string, n int, marker string\) string](<#TruncateRight>)
- [func Unwrap\(s, prefix, suffix string\) \(string, bool\)](<#Unwrap>)

<a name="AppendCSV"></a>

## func [AppendCSV](<https://github.com/gechr/x/blob/main/strings/csv.go#L5>)

```go
func AppendCSV(dst []string, raw string) []string
```

AppendCSV splits raw on commas, trims whitespace, drops empty values, and appends the remaining values to dst.

<a name="Closest"></a>

## func [Closest](<https://github.com/gechr/x/blob/main/strings/closest.go#L18>)

```go
func Closest(target string, candidates []string) string
```

Closest returns the candidate nearest to target, suitable for a "did you mean?" suggestion. Distance is the Damerau\-Levenshtein \(optimal string alignment\) edit distance, so an adjacent transposition like "verfiy" counts as one edit, not two \- the common typo plain Levenshtein over\-penalizes. It returns "" when the nearest candidate is further than a third of target's length in edits, so an unrelated word is never suggested. An empty target carries no signal and suggests nothing. Ties resolve to the first candidate.

```text
Closest("verfiy", []string{"verify", "deep"}) // "verify"
Closest("xyzzy", []string{"verify", "deep"})  // ""
```

<a name="CompactLines"></a>

## func [CompactLines](<https://github.com/gechr/x/blob/main/strings/compact.go#L7>)

```go
func CompactLines(s, sep string) string
```

CompactLines trims lines, drops blank lines, removes duplicate lines while preserving first\-seen order, and joins the remaining lines with sep.

<a name="CompareFold"></a>

## func [CompareFold](<https://github.com/gechr/x/blob/main/strings/fold.go#L12>)

```go
func CompareFold(a, b string) int
```

CompareFold compares a and b case\-insensitively, using the same simple case\-folding as strings.EqualFold, and returns \-1, 0, or 1 following the [cmp.Compare](<https://pkg.go.dev/cmp/#Compare>) convention. CompareFold\(a, b\) == 0 iff strings.EqualFold\(a, b\).

<a name="CompareNatural"></a>

## func [CompareNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L12>)

```go
func CompareNatural(a, b string) int
```

CompareNatural orders a and b the way a human reads them, treating each run of digits as a single decimal number so "x2" sorts before "x10". It returns \-1, 0, or \+1 and allocates nothing, handling numbers of any length without overflow.

<a name="ContainsAll"></a>

## func [ContainsAll](<https://github.com/gechr/x/blob/main/strings/contains.go#L6>)

```go
func ContainsAll(s string, substrings ...string) bool
```

ContainsAll reports whether s contains all of the given substrings.

<a name="ContainsAny"></a>

## func [ContainsAny](<https://github.com/gechr/x/blob/main/strings/contains.go#L16>)

```go
func ContainsAny(s string, substrings ...string) bool
```

ContainsAny reports whether s contains any of the given substrings.

<a name="CountAny"></a>

## func [CountAny](<https://github.com/gechr/x/blob/main/strings/any.go#L32>)

```go
func CountAny(s, chars string) int
```

CountAny returns the number of Unicode code points in s that are contained in chars, following the cutset convention of strings.IndexAny.

<a name="Dedent"></a>

## func [Dedent](<https://github.com/gechr/x/blob/main/strings/indent.go#L35>)

```go
func Dedent(s string) string
```

Dedent strips the longest common leading\-whitespace prefix from non\-empty lines. Whitespace\-only lines are normalized to empty \(Python textwrap.dedent\).

```text
Dedent("    foo\n      bar\n    baz") // "foo\n  bar\nbaz"
```

<a name="EqualNatural"></a>

## func [EqualNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L49>)

```go
func EqualNatural(a, b string) bool
```

EqualNatural reports whether a and b compare equal in natural order, as decided by [CompareNatural](<#CompareNatural>). This can differ from a == b, since a numeric run followed by more to compare matches regardless of leading zeros \(for example "a00b00" and "a0b00"\).

<a name="Indent"></a>

## func [Indent](<https://github.com/gechr/x/blob/main/strings/indent.go#L11>)

```go
func Indent(s, prefix string) string
```

Indent prefixes every non\-blank line of s with prefix. Blank and whitespace\-only lines are normalized to empty.

```text
Indent("foo\nbar", "  ")      // "  foo\n  bar"
Indent("foo\n\nbar", "> ")    // "> foo\n\n> bar"
Indent("foo\n   \nbar", "> ") // "> foo\n\n> bar"
```

<a name="IsBlank"></a>

## func [IsBlank](<https://github.com/gechr/x/blob/main/strings/blank.go#L6>)

```go
func IsBlank(s string) bool
```

IsBlank reports whether s is empty or consists only of whitespace.

<a name="IsGitCommit"></a>

## func [IsGitCommit](<https://github.com/gechr/x/blob/main/strings/hex.go#L23>)

```go
func IsGitCommit(s string) bool
```

IsGitCommit reports whether s is 40 hexadecimal digits \(a Git commit hash\).

<a name="IsHex"></a>

## func [IsHex](<https://github.com/gechr/x/blob/main/strings/hex.go#L5>)

```go
func IsHex(s string) bool
```

IsHex reports whether s is non\-empty and consists entirely of hexadecimal digits. An empty string is not hex.

<a name="IsHexChar"></a>

## func [IsHexChar](<https://github.com/gechr/x/blob/main/strings/hex.go#L18>)

```go
func IsHexChar(c rune) bool
```

IsHexChar reports whether c is a valid hexadecimal digit \(0\-9, a\-f, A\-F\).

<a name="IsSHA256"></a>

## func [IsSHA256](<https://github.com/gechr/x/blob/main/strings/hex.go#L28>)

```go
func IsSHA256(s string) bool
```

IsSHA256 reports whether s is 64 hexadecimal digits \(a SHA\-256 digest\).

<a name="LessNatural"></a>

## func [LessNatural](<https://github.com/gechr/x/blob/main/strings/natural.go#L41>)

```go
func LessNatural(a, b string) bool
```

LessNatural reports whether a sorts before b in natural order, as decided by [CompareNatural](<#CompareNatural>). It reads cleanly at call sites that want a boolean rather than a three\-way result, such as sort predicates and conditionals.

<a name="SplitAny"></a>

## func [SplitAny](<https://github.com/gechr/x/blob/main/strings/any.go#L14>)

```go
func SplitAny(s, chars string) []string
```

SplitAny splits s around each occurrence of any Unicode code point in chars, following the cutset convention of strings.IndexAny. Empty segments between adjacent separators are preserved, matching strings.Split semantics. If chars is empty, SplitAny returns a single\-element slice containing s.

<a name="SplitBy"></a>

## func [SplitBy](<https://github.com/gechr/x/blob/main/strings/split.go#L7>)

```go
func SplitBy(s, sep string) []string
```

SplitBy splits s by sep, trims whitespace from each part, and drops empty values.

<a name="SplitCSV"></a>

## func [SplitCSV](<https://github.com/gechr/x/blob/main/strings/csv.go#L10>)

```go
func SplitCSV(s string) []string
```

SplitCSV splits s on commas, trims whitespace, and drops empty values.

<a name="SplitLines"></a>

## func [SplitLines](<https://github.com/gechr/x/blob/main/strings/lines.go#L6>)

```go
func SplitLines(s string) []string
```

SplitLines splits s into non\-empty trimmed lines.

<a name="SplitLinesRaw"></a>

## func [SplitLinesRaw](<https://github.com/gechr/x/blob/main/strings/lines.go#L13>)

```go
func SplitLinesRaw(s string) []string
```

SplitLinesRaw splits s into lines losslessly, normalizing CRLF to LF: every line is kept verbatim \- empty lines and the trailing empty element included \- so the result joins back with "\\n" without losing content or line numbers.

<a name="Truncate"></a>

## func [Truncate](<https://github.com/gechr/x/blob/main/strings/truncate.go#L44>)

```go
func Truncate(s string, n int, marker string) string
```

Truncate is an alias for [TruncateRight](<#TruncateRight>), the most common form: it keeps the head and trims the tail.

<a name="TruncateLeft"></a>

## func [TruncateLeft](<https://github.com/gechr/x/blob/main/strings/truncate.go#L22>)

```go
func TruncateLeft(s string, n int, marker string) string
```

TruncateLeft shortens s to at most n runes \(including marker\) by removing characters from the left, prepending marker when truncation occurs. The tail is kept.

```text
TruncateLeft("hello world", 8, "…") // "…o world"
```

<a name="TruncateMiddle"></a>

## func [TruncateMiddle](<https://github.com/gechr/x/blob/main/strings/truncate.go#L34>)

```go
func TruncateMiddle(s string, n int, marker string) string
```

TruncateMiddle shortens s to at most n runes \(including marker\) by removing characters from the middle, inserting marker between the kept head and tail so both ends stay visible. This suits hashes and paths, where the start and end are the recognisable parts.

```text
TruncateMiddle("0123456789abcdef", 7, "…") // "012…def"
```

<a name="TruncateRight"></a>

## func [TruncateRight](<https://github.com/gechr/x/blob/main/strings/truncate.go#L11>)

```go
func TruncateRight(s string, n int, marker string) string
```

TruncateRight shortens s to at most n runes \(including marker\) by removing characters from the right, appending marker when truncation occurs. The head is kept. For display\-width\-aware truncation of ANSI text use ansi.Truncate.

```text
TruncateRight("hello world", 8, "…") // "hello w…"
TruncateRight("hi", 8, "…")          // "hi"
```

<a name="Unwrap"></a>

## func [Unwrap](<https://github.com/gechr/x/blob/main/strings/unwrap.go#L9>)

```go
func Unwrap(s, prefix, suffix string) (string, bool)
```

Unwrap returns s with the leading prefix and trailing suffix removed and reports whether both were present. Unlike a strings.TrimPrefix \+ strings.TrimSuffix chain, nothing is removed unless s starts with prefix AND ends with suffix, so a one\-sided match is returned unchanged.
