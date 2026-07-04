# slices

```go
import "github.com/gechr/x/slices"
```

Package slices provides slice helpers.

## Index

- [func ContainsFold\[S ~\[\]E, E ~string\](items S, target E) bool](<#ContainsFold>)
- [func Count\[S ~\[\]E, E comparable\](items S, target E) int](<#Count>)
- [func CountFunc\[S ~\[\]E, E any\](items S, match func(E) bool) int](<#CountFunc>)
- [func Difference\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Difference>)
- [func Intersect\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Intersect>)
- [func LastIndex\[S ~\[\]E, E comparable\](items S, target E) int](<#LastIndex>)
- [func LastIndexFunc\[S ~\[\]E, E any\](items S, match func(E) bool) int](<#LastIndexFunc>)
- [func Partition\[S ~\[\]E, E any\](items S, match func(E) bool) (S, S)](<#Partition>)
- [func SortNatural\[S ~\[\]E, E ~string\](s S)](<#SortNatural>)
- [func Trim\[S ~\[\]E, E comparable\](items, cutset S) S](<#Trim>)
- [func TrimLeft\[S ~\[\]E, E comparable\](items, cutset S) S](<#TrimLeft>)
- [func TrimRight\[S ~\[\]E, E comparable\](items, cutset S) S](<#TrimRight>)
- [func Union\[S ~\[\]E, E comparable\](items S, others ...S) S](<#Union>)
- [func Unique\[S ~\[\]E, E comparable\](items S) S](<#Unique>)
- [func UniqueFold\[S ~\[\]E, E ~string\](items S) S](<#UniqueFold>)

<a name="ContainsFold"></a>

## func [ContainsFold](<https://github.com/gechr/x/blob/main/slices/contains.go#L11>)

```go
func ContainsFold[S ~[]E, E ~string](items S, target E) bool
```

**ContainsFold** reports whether items contains target case-insensitively, using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).

<a name="Count"></a>

## func [Count](<https://github.com/gechr/x/blob/main/slices/count.go#L4>)

```go
func Count[S ~[]E, E comparable](items S, target E) int
```

**Count** returns the number of elements in items equal to target.

<a name="CountFunc"></a>

## func [CountFunc](<https://github.com/gechr/x/blob/main/slices/count.go#L15>)

```go
func CountFunc[S ~[]E, E any](items S, match func(E) bool) int
```

**CountFunc** returns the number of elements in items satisfying match.

<a name="Difference"></a>

## func [Difference](<https://github.com/gechr/x/blob/main/slices/sets.go#L5>)

```go
func Difference[S ~[]E, E comparable](items S, others ...S) S
```

**Difference** returns the elements of items not present in any of others, preserving order and duplicates from items.

<a name="Intersect"></a>

## func [Intersect](<https://github.com/gechr/x/blob/main/slices/sets.go#L23>)

```go
func Intersect[S ~[]E, E comparable](items S, others ...S) S
```

**Intersect** returns the elements of items also present in every one of others, preserving order and duplicates from items.

<a name="LastIndex"></a>

## func [LastIndex](<https://github.com/gechr/x/blob/main/slices/index.go#L5>)

```go
func LastIndex[S ~[]E, E comparable](items S, target E) int
```

**LastIndex** returns the index of the last occurrence of target in items, or -1 if not present.

<a name="LastIndexFunc"></a>

## func [LastIndexFunc](<https://github.com/gechr/x/blob/main/slices/index.go#L16>)

```go
func LastIndexFunc[S ~[]E, E any](items S, match func(E) bool) int
```

**LastIndexFunc** returns the index of the last element of items satisfying match, or -1 if none do.

<a name="Partition"></a>

## func [Partition](<https://github.com/gechr/x/blob/main/slices/partition.go#L5>)

```go
func Partition[S ~[]E, E any](items S, match func(E) bool) (S, S)
```

**Partition** splits items into two slices: elements satisfying match, and elements that do not, preserving the original relative order in both.

<a name="SortNatural"></a>

## func [SortNatural](<https://github.com/gechr/x/blob/main/slices/sort.go#L12>)

```go
func SortNatural[S ~[]E, E ~string](s S)
```

**SortNatural** sorts a string slice in place in natural order, so embedded numbers compare by value ("item2" before "item10") rather than lexically. See [xstrings.CompareNatural](<https://pkg.go.dev/github.com/gechr/x/strings#CompareNatural>).

<a name="Trim"></a>

## func [Trim](<https://github.com/gechr/x/blob/main/slices/trim.go#L6>)

```go
func Trim[S ~[]E, E comparable](items, cutset S) S
```

**Trim** returns items with all leading and trailing elements contained in cutset removed. The result is a subslice of items, sharing its backing array.

<a name="TrimLeft"></a>

## func [TrimLeft](<https://github.com/gechr/x/blob/main/slices/trim.go#L12>)

```go
func TrimLeft[S ~[]E, E comparable](items, cutset S) S
```

**TrimLeft** returns items with all leading elements contained in cutset removed. The result is a subslice of items, sharing its backing array.

<a name="TrimRight"></a>

## func [TrimRight](<https://github.com/gechr/x/blob/main/slices/trim.go#L26>)

```go
func TrimRight[S ~[]E, E comparable](items, cutset S) S
```

**TrimRight** returns items with all trailing elements contained in cutset removed. The result is a subslice of items, sharing its backing array.

<a name="Union"></a>

## func [Union](<https://github.com/gechr/x/blob/main/slices/sets.go#L46>)

```go
func Union[S ~[]E, E comparable](items S, others ...S) S
```

**Union** returns the elements of items followed by the elements of others, in first-seen order with duplicates removed.

<a name="Unique"></a>

## func [Unique](<https://github.com/gechr/x/blob/main/slices/unique.go#L9>)

```go
func Unique[S ~[]E, E comparable](items S) S
```

**Unique** returns items in first-seen order with duplicates removed.

<a name="UniqueFold"></a>

## func [UniqueFold](<https://github.com/gechr/x/blob/main/slices/unique.go#L25>)

```go
func UniqueFold[S ~[]E, E ~string](items S) S
```

**UniqueFold** returns strings in first-seen order with duplicates removed case-insensitively, using the same simple case-folding as [strings.EqualFold](<https://pkg.go.dev/strings#EqualFold>).
