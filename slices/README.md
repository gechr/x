# slices

```go
import "github.com/gechr/x/slices"
```

Package `slices` provides slice helpers.

## Index

- [func AllFunc\[S ~\[\]E, E any\](items S, match func(E) bool) bool](<#AllFunc>)
- [func AnyFunc\[S ~\[\]E, E any\](items S, match func(E) bool) bool](<#AnyFunc>)
- [func ContainedByAll\[S ~\[\]E, E comparable\](target E, lists ...S) bool](<#ContainedByAll>)
- [func ContainedByAny\[S ~\[\]E, E comparable\](target E, lists ...S) bool](<#ContainedByAny>)
- [func ContainsAll\[S ~\[\]E, E comparable\](items S, targets ...E) bool](<#ContainsAll>)
- [func ContainsAny\[S ~\[\]E, E comparable\](items S, targets ...E) bool](<#ContainsAny>)
- [func ContainsFold\[S ~\[\]E, E ~string\](items S, target E) bool](<#ContainsFold>)
- [func Count\[S ~\[\]E, E comparable\](items S, target E) int](<#Count>)
- [func CountFunc\[S ~\[\]E, E any\](items S, match func(E) bool) int](<#CountFunc>)
- [func Difference\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Difference>)
- [func Filter\[S ~\[\]E, E any\](items S, keep func(E) bool) S](<#Filter>)
- [func Format(format string, args ...any) \[\]string](<#Format>)
- [func Intersect\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Intersect>)
- [func LastIndex\[S ~\[\]E, E comparable\](items S, target E) int](<#LastIndex>)
- [func LastIndexFunc\[S ~\[\]E, E any\](items S, match func(E) bool) int](<#LastIndexFunc>)
- [func Map\[S ~\[\]E, E, R any\](items S, fn func(E) R) \[\]R](<#Map>)
- [func Partition\[S ~\[\]E, E any\](items S, match func(E) bool) (S, S)](<#Partition>)
- [func Reject\[S ~\[\]E, E any\](items S, drop func(E) bool) S](<#Reject>)
- [func SortNatural\[S ~\[\]E, E ~string\](s S)](<#SortNatural>)
- [func Surround\[S ~\[\]E, E ~string\](items S, prefix, suffix E) \[\]E](<#Surround>)
- [func Trim\[S ~\[\]E, E comparable\](items, cutset S) S](<#Trim>)
- [func TrimLeft\[S ~\[\]E, E comparable\](items, cutset S) S](<#TrimLeft>)
- [func TrimRight\[S ~\[\]E, E comparable\](items, cutset S) S](<#TrimRight>)
- [func Union\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Union>)
- [func Unique\[S ~\[\]E, E comparable\](items S) S](<#Unique>)
- [func UniqueFold\[S ~\[\]E, E ~string\](items S) S](<#UniqueFold>)
- [func UniqueFunc\[S ~\[\]E, E any, K comparable\](items S, key func(E) K) S](<#UniqueFunc>)

<a name="AllFunc"></a>

## func [AllFunc](<https://github.com/gechr/x/blob/main/slices/match.go#L7>)

```go
func AllFunc[S ~[]E, E any](items S, match func(E) bool) bool
```

**AllFunc** reports whether every element of `items` satisfies `match`. It returns true when `items` is empty.

<details><summary><b>Example</b></summary>

Every element must satisfy the predicate; an empty slice reports true.

```go
isEven := func(n int) bool { return n%2 == 0 }
fmt.Println(xslices.AllFunc([]int{2, 4, 6}, isEven))
fmt.Println(xslices.AllFunc([]int{2, 3, 4}, isEven))
fmt.Println(xslices.AllFunc([]int{}, isEven))
```

Output:

```text
true
false
true
```

</details>

<a name="AnyFunc"></a>

## func [AnyFunc](<https://github.com/gechr/x/blob/main/slices/match.go#L15>)

```go
func AnyFunc[S ~[]E, E any](items S, match func(E) bool) bool
```

**AnyFunc** reports whether any element of `items` satisfies `match`. It returns false when `items` is empty.

<details><summary><b>Example</b></summary>

A single matching element suffices; an empty slice reports false.

```go
isEven := func(n int) bool { return n%2 == 0 }
fmt.Println(xslices.AnyFunc([]int{1, 2, 3}, isEven))
fmt.Println(xslices.AnyFunc([]int{1, 3, 5}, isEven))
fmt.Println(xslices.AnyFunc([]int{}, isEven))
```

Output:

```text
true
false
false
```

</details>

<a name="ContainedByAll"></a>

## func [ContainedByAll](<https://github.com/gechr/x/blob/main/slices/contains.go#L11>)

```go
func ContainedByAll[S ~[]E, E comparable](target E, lists ...S) bool
```

**ContainedByAll** reports whether `target` occurs in every one of `lists`. It returns true when no `lists` are given.

<details><summary><b>Example</b></summary>

Every slice must contain the target; no slices reports true.

```go
fmt.Println(xslices.ContainedByAll("a", []string{"a", "b"}, []string{"c", "a"}))
fmt.Println(xslices.ContainedByAll("b", []string{"a", "b"}, []string{"c", "a"}))
fmt.Println(xslices.ContainedByAll[[]string]("a"))
```

Output:

```text
true
false
true
```

</details>

<a name="ContainedByAny"></a>

## func [ContainedByAny](<https://github.com/gechr/x/blob/main/slices/contains.go#L18>)

```go
func ContainedByAny[S ~[]E, E comparable](target E, lists ...S) bool
```

**ContainedByAny** reports whether `target` occurs in any one of `lists`.

<details><summary><b>Example</b></summary>

A single slice containing the target suffices; no slices reports false.

```go
fmt.Println(xslices.ContainedByAny("b", []string{"a", "b"}, []string{"c", "d"}))
fmt.Println(xslices.ContainedByAny("z", []string{"a", "b"}, []string{"c", "d"}))
fmt.Println(xslices.ContainedByAny[[]string]("a"))
```

Output:

```text
true
false
false
```

</details>

<a name="ContainsAll"></a>

## func [ContainsAll](<https://github.com/gechr/x/blob/main/slices/contains.go#L26>)

```go
func ContainsAll[S ~[]E, E comparable](items S, targets ...E) bool
```

**ContainsAll** reports whether `items` contains every one of `targets`. It returns true when no `targets` are given.

<details><summary><b>Example</b></summary>

Every target must occur in the slice; no targets reports true.

```go
fmt.Println(xslices.ContainsAll([]string{"a", "b", "c"}, "a", "c"))
fmt.Println(xslices.ContainsAll([]string{"a", "b", "c"}, "a", "z"))
fmt.Println(xslices.ContainsAll([]string{"a", "b", "c"}))
```

Output:

```text
true
false
true
```

</details>

<a name="ContainsAny"></a>

## func [ContainsAny](<https://github.com/gechr/x/blob/main/slices/contains.go#L33>)

```go
func ContainsAny[S ~[]E, E comparable](items S, targets ...E) bool
```

**ContainsAny** reports whether `items` contains any one of `targets`.

<details><summary><b>Example</b></summary>

A single matching target suffices; no targets reports false.

```go
fmt.Println(xslices.ContainsAny([]string{"a", "b", "c"}, "b", "z"))
fmt.Println(xslices.ContainsAny([]string{"a", "b", "c"}, "x", "z"))
fmt.Println(xslices.ContainsAny([]string{"a", "b", "c"}))
```

Output:

```text
true
false
false
```

</details>

<a name="ContainsFold"></a>

## func [ContainsFold](<https://github.com/gechr/x/blob/main/slices/contains.go#L41>)

```go
func ContainsFold[S ~[]E, E ~string](items S, target E) bool
```

**ContainsFold** reports whether `items` contains `target` case-insensitively, using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<details><summary><b>Example</b></summary>

```go
tags := []string{"Latest", "Stable"}
fmt.Println(xslices.ContainsFold(tags, "latest"))
fmt.Println(xslices.ContainsFold(tags, "STABLE"))
fmt.Println(xslices.ContainsFold(tags, "beta"))
```

Output:

```text
true
true
false
```

</details>

<a name="Count"></a>

## func [Count](<https://github.com/gechr/x/blob/main/slices/count.go#L4>)

```go
func Count[S ~[]E, E comparable](items S, target E) int
```

**Count** returns the number of elements in `items` equal to `target`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Count([]string{"a", "b", "a", "c", "a"}, "a"))
fmt.Println(xslices.Count([]string{"a", "b"}, "z"))
```

Output:

```text
3
0
```

</details>

<a name="CountFunc"></a>

## func [CountFunc](<https://github.com/gechr/x/blob/main/slices/count.go#L15>)

```go
func CountFunc[S ~[]E, E any](items S, match func(E) bool) int
```

**CountFunc** returns the number of elements in `items` satisfying `match`.

<details><summary><b>Example</b></summary>

```go
isEven := func(n int) bool { return n%2 == 0 }
fmt.Println(xslices.CountFunc([]int{1, 2, 3, 4, 5, 6}, isEven))
```

Output:

```text
3
```

</details>

<a name="Difference"></a>

## func [Difference](<https://github.com/gechr/x/blob/main/slices/sets.go#L7>)

```go
func Difference[S ~[]E, E comparable](items S, others ...S) S
```

**Difference** returns the elements of `items` not present in any of `others`, preserving order and duplicates from `items`.

<details><summary><b>Example</b></summary>

Duplicates and order in the first slice are preserved.

```go
fmt.Println(xslices.Difference([]int{1, 1, 2, 3}, []int{2}))
fmt.Println(xslices.Difference([]int{1, 2, 3}, []int{2}, []int{3}))
```

Output:

```text
[1 1 3]
[1]
```

</details>

<a name="Filter"></a>

## func [Filter](<https://github.com/gechr/x/blob/main/slices/filter.go#L5>)

```go
func Filter[S ~[]E, E any](items S, keep func(E) bool) S
```

**Filter** returns the elements of `items` satisfying `keep`, preserving their original order.

<details><summary><b>Example</b></summary>

```go
items := []int{1, 2, 3, 4, 5, 6}
fmt.Println(xslices.Filter(items, func(n int) bool { return n%2 == 0 }))
```

Output:

```text
[2 4 6]
```

</details>

<a name="Format"></a>

## func [Format](<https://github.com/gechr/x/blob/main/slices/format.go#L15>)

```go
func Format(format string, args ...any) []string
```

**Format** returns the result of applying fmt.Sprintf(format, ...) once per element of the shortest slice in args, substituting each slice argument with its i'th element while repeating non-slice arguments unchanged. If args contains no slices, it returns a single formatted string.

Byte slices (\[\]byte and named types with that underlying type) are treated as scalars rather than iterated, so they format as a single value.

<details><summary><b>Example</b></summary>

Scalar arguments repeat for every element; the shortest slice determines the result length.

```go
names := []string{"Valentina", "Ander", "Olivia", "Sam"}
fmt.Println(xslices.Format("Hello, %s!", names))
fmt.Println(xslices.Format("%s, %s!", "Salutations", names))
```

Output:

```text
[Hello, Valentina! Hello, Ander! Hello, Olivia! Hello, Sam!]
[Salutations, Valentina! Salutations, Ander! Salutations, Olivia! Salutations, Sam!]
```

</details>

<a name="Intersect"></a>

## func [Intersect](<https://github.com/gechr/x/blob/main/slices/sets.go#L23>)

```go
func Intersect[S ~[]E, E comparable](items S, others ...S) S
```

**Intersect** returns the elements of `items` also present in every one of `others`, preserving order and duplicates from `items`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Intersect([]string{"a", "b", "c"}, []string{"c", "b", "d"}))
fmt.Println(xslices.Intersect([]int{1, 2, 3}, []int{2, 3}, []int{2}))
```

Output:

```text
[b c]
[2]
```

</details>

<a name="LastIndex"></a>

## func [LastIndex](<https://github.com/gechr/x/blob/main/slices/index.go#L5>)

```go
func LastIndex[S ~[]E, E comparable](items S, target E) int
```

**LastIndex** returns the index of the last occurrence of `target` in `items`, or -1 if not present.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.LastIndex([]string{"a", "b", "a", "c"}, "a"))
fmt.Println(xslices.LastIndex([]string{"a", "b"}, "z"))
```

Output:

```text
2
-1
```

</details>

<a name="LastIndexFunc"></a>

## func [LastIndexFunc](<https://github.com/gechr/x/blob/main/slices/index.go#L16>)

```go
func LastIndexFunc[S ~[]E, E any](items S, match func(E) bool) int
```

**LastIndexFunc** returns the index of the last element of `items` satisfying `match`, or -1 if none do.

<details><summary><b>Example</b></summary>

```go
isEven := func(n int) bool { return n%2 == 0 }
fmt.Println(xslices.LastIndexFunc([]int{2, 3, 4, 5}, isEven))
fmt.Println(xslices.LastIndexFunc([]int{1, 3, 5}, isEven))
```

Output:

```text
2
-1
```

</details>

<a name="Map"></a>

## func [Map](<https://github.com/gechr/x/blob/main/slices/map.go#L5>)

```go
func Map[S ~[]E, E, R any](items S, fn func(E) R) []R
```

**Map** returns a new slice containing the result of applying `fn` to each element of `items`, preserving order.

<a name="Partition"></a>

## func [Partition](<https://github.com/gechr/x/blob/main/slices/partition.go#L5>)

```go
func Partition[S ~[]E, E any](items S, match func(E) bool) (S, S)
```

**Partition** splits `items` into two slices: elements satisfying `match`, and elements that do not, preserving the original relative order in both.

<details><summary><b>Example</b></summary>

```go
isEven := func(n int) bool { return n%2 == 0 }
even, odd := xslices.Partition([]int{1, 2, 3, 4, 5, 6}, isEven)
fmt.Println(even)
fmt.Println(odd)
```

Output:

```text
[2 4 6]
[1 3 5]
```

</details>

<a name="Reject"></a>

## func [Reject](<https://github.com/gechr/x/blob/main/slices/reject.go#L5>)

```go
func Reject[S ~[]E, E any](items S, drop func(E) bool) S
```

**Reject** returns the elements of `items` not satisfying `drop`, preserving their original order.

<details><summary><b>Example</b></summary>

```go
items := []int{1, 2, 3, 4, 5, 6}
fmt.Println(xslices.Reject(items, func(n int) bool { return n%2 == 0 }))
```

Output:

```text
[1 3 5]
```

</details>

<a name="SortNatural"></a>

## func [SortNatural](<https://github.com/gechr/x/blob/main/slices/sort.go#L12>)

```go
func SortNatural[S ~[]E, E ~string](s S)
```

**SortNatural** sorts a string slice in place in natural order, so embedded numbers compare by value (`item2` before `item10`) rather than lexically. See [strings.CompareNatural](<../strings/README.md#CompareNatural>).

<details><summary><b>Example</b></summary>

Embedded numbers compare by value, so "item2" sorts before "item10".

```go
items := []string{"item10", "item2", "item1", "item20", "item3"}
xslices.SortNatural(items)
fmt.Println(items)
```

Output:

```text
[item1 item2 item3 item10 item20]
```

</details>

<a name="Surround"></a>

## func [Surround](<https://github.com/gechr/x/blob/main/slices/surround.go#L5>)

```go
func Surround[S ~[]E, E ~string](items S, prefix, suffix E) []E
```

**Surround** returns a new slice with `prefix` and `suffix` concatenated onto each element of `items`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Surround([]string{"a", "b", "c"}, `"`, `"`))
```

Output:

```text
["a" "b" "c"]
```

</details>

<a name="Trim"></a>

## func [Trim](<https://github.com/gechr/x/blob/main/slices/trim.go#L8>)

```go
func Trim[S ~[]E, E comparable](items, cutset S) S
```

**Trim** returns `items` with all leading and trailing elements contained in `cutset` removed. The result is a subslice of `items`, sharing its backing array.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Trim([]int{0, 0, 1, 2, 0}, []int{0}))
fmt.Println(xslices.Trim([]string{"a", "b", "c", "b", "a"}, []string{"a", "b"}))
```

Output:

```text
[1 2]
[c]
```

</details>

<a name="TrimLeft"></a>

## func [TrimLeft](<https://github.com/gechr/x/blob/main/slices/trim.go#L14>)

```go
func TrimLeft[S ~[]E, E comparable](items, cutset S) S
```

**TrimLeft** returns `items` with all leading elements contained in `cutset` removed. The result is a subslice of `items`, sharing its backing array.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.TrimLeft([]int{0, 0, 1, 2, 0}, []int{0}))
```

Output:

```text
[1 2 0]
```

</details>

<a name="TrimRight"></a>

## func [TrimRight](<https://github.com/gechr/x/blob/main/slices/trim.go#L28>)

```go
func TrimRight[S ~[]E, E comparable](items, cutset S) S
```

**TrimRight** returns `items` with all trailing elements contained in `cutset` removed. The result is a subslice of `items`, sharing its backing array.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.TrimRight([]int{0, 0, 1, 2, 0}, []int{0}))
```

Output:

```text
[0 0 1 2]
```

</details>

<a name="Union"></a>

## func [Union](<https://github.com/gechr/x/blob/main/slices/sets.go#L46>)

```go
func Union[S ~[]E, E comparable](items S, others ...S) S
```

**Union** returns the elements of `items` followed by the elements of `others`, in first-seen order with duplicates removed.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Union([]int{1, 2}, []int{2, 3}, []int{3, 4, 1}))
```

Output:

```text
[1 2 3 4]
```

</details>

<a name="Unique"></a>

## func [Unique](<https://github.com/gechr/x/blob/main/slices/unique.go#L11>)

```go
func Unique[S ~[]E, E comparable](items S) S
```

**Unique** returns `items` in first-seen order with duplicates removed.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xslices.Unique([]string{"a", "b", "a", "A", "c", "b"}))
```

Output:

```text
[a b A c]
```

</details>

<a name="UniqueFold"></a>

## func [UniqueFold](<https://github.com/gechr/x/blob/main/slices/unique.go#L43>)

```go
func UniqueFold[S ~[]E, E ~string](items S) S
```

**UniqueFold** returns strings in first-seen order with duplicates removed case-insensitively, using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<details><summary><b>Example</b></summary>

The first-seen spelling wins; folding matches strings.EqualFold, so Greek final sigma "ς", medial "σ", and capital "Σ" are duplicates.

```go
fmt.Println(xslices.UniqueFold([]string{"Go", "GO", "go", "Rust"}))
fmt.Println(xslices.UniqueFold([]string{"ς", "σ", "Σ"}))
```

Output:

```text
[Go Rust]
[ς]
```

</details>

<a name="UniqueFunc"></a>

## func [UniqueFunc](<https://github.com/gechr/x/blob/main/slices/unique.go#L26>)

```go
func UniqueFunc[S ~[]E, E any, K comparable](items S, key func(E) K) S
```

**UniqueFunc** returns `items` in first-seen order with duplicates removed, where two items are duplicates when `key` reports the same value for both.

<details><summary><b>Example</b></summary>

The first-seen item wins for each key.

```go
type user struct {
    name string
    age  int
}
users := []user{{"alice", 30}, {"bob", 25}, {"alice", 40}}
fmt.Println(xslices.UniqueFunc(users, func(u user) string { return u.name }))
```

Output:

```text
[{alice 30} {bob 25}]
```

</details>
