# set

```go
import "github.com/gechr/x/set"
```

Package `set` provides a generic set backed by a map.

## Index

- [func Sorted\[T cmp.Ordered\](s Set\[T\]) \[\]T](<#Sorted>)
- [func SortedNatural\[T ~string\](s Set\[T\]) \[\]T](<#SortedNatural>)
- [type Set](<#Set>)
  - [func Collect\[T comparable\](seq iter.Seq\[T\]) Set\[T\]](<#Collect>)
  - [func New\[T comparable\](items ...T) Set\[T\]](<#New>)
  - [func (s Set\[T\]) Add(items ...T)](<#Set.Add>)
  - [func (s Set\[T\]) All() iter.Seq\[T\]](<#Set.All>)
  - [func (s Set\[T\]) Clone() Set\[T\]](<#Set.Clone>)
  - [func (s Set\[T\]) Contains(item T) bool](<#Set.Contains>)
  - [func (s Set\[T\]) Delete(items ...T)](<#Set.Delete>)
  - [func (s Set\[T\]) Difference(others ...Set\[T\]) Set\[T\]](<#Set.Difference>)
  - [func (s Set\[T\]) Equal(other Set\[T\]) bool](<#Set.Equal>)
  - [func (s Set\[T\]) Intersect(others ...Set\[T\]) Set\[T\]](<#Set.Intersect>)
  - [func (s Set\[T\]) Len() int](<#Set.Len>)
  - [func (s Set\[T\]) Slice() \[\]T](<#Set.Slice>)
  - [func (s Set\[T\]) SubsetOf(other Set\[T\]) bool](<#Set.SubsetOf>)
  - [func (s Set\[T\]) Union(others ...Set\[T\]) Set\[T\]](<#Set.Union>)
- [type SortedSet](<#SortedSet>)
  - [func CollectSorted\[T cmp.Ordered\](seq iter.Seq\[T\]) SortedSet\[T\]](<#CollectSorted>)
  - [func NewSorted\[T cmp.Ordered\](items ...T) SortedSet\[T\]](<#NewSorted>)
  - [func (s \*SortedSet\[T\]) Add(items ...T)](<#SortedSet.Add>)
  - [func (s SortedSet\[T\]) All() iter.Seq\[T\]](<#SortedSet.All>)
  - [func (s SortedSet\[T\]) Clone() SortedSet\[T\]](<#SortedSet.Clone>)
  - [func (s SortedSet\[T\]) Contains(item T) bool](<#SortedSet.Contains>)
  - [func (s \*SortedSet\[T\]) Delete(items ...T)](<#SortedSet.Delete>)
  - [func (s SortedSet\[T\]) Difference(others ...SortedSet\[T\]) SortedSet\[T\]](<#SortedSet.Difference>)
  - [func (s SortedSet\[T\]) Equal(other SortedSet\[T\]) bool](<#SortedSet.Equal>)
  - [func (s SortedSet\[T\]) Intersect(others ...SortedSet\[T\]) SortedSet\[T\]](<#SortedSet.Intersect>)
  - [func (s SortedSet\[T\]) Len() int](<#SortedSet.Len>)
  - [func (s SortedSet\[T\]) Slice() \[\]T](<#SortedSet.Slice>)
  - [func (s SortedSet\[T\]) SubsetOf(other SortedSet\[T\]) bool](<#SortedSet.SubsetOf>)
  - [func (s SortedSet\[T\]) Union(others ...SortedSet\[T\]) SortedSet\[T\]](<#SortedSet.Union>)

<a name="Sorted"></a>

## func [Sorted](<https://github.com/gechr/x/blob/main/set/sorted.go#L14>)

```go
func Sorted[T cmp.Ordered](s Set[T]) []T
```

**Sorted** returns the items of `s` as a slice in ascending order.

Sorted is a function rather than a [Set](<#Set>) method because it requires `T` to be ordered, not just comparable.

<details><summary><b>Example</b></summary>

**Sorted** turns a Set's indeterminate iteration order into a stable one.

```go
s := set.New("banana", "apple", "cherry")
fmt.Println(set.Sorted(s))
```

Output:

```text
[apple banana cherry]
```

</details>

<a name="SortedNatural"></a>

## func [SortedNatural](<https://github.com/gechr/x/blob/main/set/sorted.go#L26>)

```go
func SortedNatural[T ~string](s Set[T]) []T
```

**SortedNatural** returns the items of `s` as a slice in natural order, so embedded numbers compare by value (`item2` before `item10`) rather than lexically. See [strings.CompareNatural](<../strings/README.md#CompareNatural>).

SortedNatural is a function rather than a [Set](<#Set>) method because it requires `T` to be string-like, not just comparable.

<details><summary><b>Example</b></summary>

**SortedNatural** orders embedded numbers by value, so "item2" sorts before "item10".

```go
s := set.New("item10", "item2", "item1")
fmt.Println(set.Sorted(s))
fmt.Println(set.SortedNatural(s))
```

Output:

```text
[item1 item10 item2]
[item1 item2 item10]
```

</details>

<a name="Set"></a>

## type [Set](<https://github.com/gechr/x/blob/main/set/set.go#L13>)

**Set** is a set of comparable items backed by a map. Pointer types and structs containing pointer fields are compared using shallow equality. The zero value is nil: read operations work, but [Set.Add](<#Set.Add>) panics - use [New](<#New>) or [Collect](<#Collect>) to create a usable [Set](<#Set>).

```go
type Set[T comparable] map[T]struct{}
```

<a name="Collect"></a>

### func [Collect](<https://github.com/gechr/x/blob/main/set/set.go#L23>)

```go
func Collect[T comparable](seq iter.Seq[T]) Set[T]
```

**Collect** returns a [Set](<#Set>) containing the values of `seq`.

<details><summary><b>Example</b></summary>

**Collect** builds a Set from any iter.Seq, such as slices.Values.

```go
s := set.Collect(slices.Values([]int{3, 1, 3, 2}))
fmt.Println(set.Sorted(s))
```

Output:

```text
[1 2 3]
```

</details>

<a name="New"></a>

### func [New](<https://github.com/gechr/x/blob/main/set/set.go#L16>)

```go
func New[T comparable](items ...T) Set[T]
```

**New** returns a [Set](<#Set>) containing `items`.

<details><summary><b>Example</b></summary>

```go
s := set.New("a", "b", "a")
fmt.Println(s.Len())
fmt.Println(s.Contains("b"))
fmt.Println(s.Contains("c"))
```

Output:

```text
2
true
false
```

</details>

<a name="Set.Add"></a>

### func (Set\[T\]) [Add](<https://github.com/gechr/x/blob/main/set/set.go#L32>)

```go
func (s Set[T]) Add(items ...T)
```

**Add** adds `items` to `s`.

<a name="Set.All"></a>

### func (Set\[T\]) [All](<https://github.com/gechr/x/blob/main/set/set.go#L139>)

```go
func (s Set[T]) All() iter.Seq[T]
```

**All** returns an iterator over the items of `s`, in indeterminate order.

<a name="Set.Clone"></a>

### func (Set\[T\]) [Clone](<https://github.com/gechr/x/blob/main/set/set.go#L121>)

```go
func (s Set[T]) Clone() Set[T]
```

**Clone** returns a copy of `s`.

<a name="Set.Contains"></a>

### func (Set\[T\]) [Contains](<https://github.com/gechr/x/blob/main/set/set.go#L46>)

```go
func (s Set[T]) Contains(item T) bool
```

**Contains** returns whether `item` is present in `s`.

<a name="Set.Delete"></a>

### func (Set\[T\]) [Delete](<https://github.com/gechr/x/blob/main/set/set.go#L39>)

```go
func (s Set[T]) Delete(items ...T)
```

**Delete** removes `items` from `s`.

<a name="Set.Difference"></a>

### func (Set\[T\]) [Difference](<https://github.com/gechr/x/blob/main/set/set.go#L110>)

```go
func (s Set[T]) Difference(others ...Set[T]) Set[T]
```

**Difference** returns a new [Set](<#Set>) containing the items of `s` not present in any of `others`.

<details><summary><b>Example</b></summary>

```go
a := set.New(1, 2, 3)
b := set.New(2, 4)
fmt.Println(set.Sorted(a.Difference(b)))
```

Output:

```text
[1 3]
```

</details>

<a name="Set.Equal"></a>

### func (Set\[T\]) [Equal](<https://github.com/gechr/x/blob/main/set/set.go#L57>)

```go
func (s Set[T]) Equal(other Set[T]) bool
```

**Equal** returns whether `s` and `other` contain the same items.

<a name="Set.Intersect"></a>

### func (Set\[T\]) [Intersect](<https://github.com/gechr/x/blob/main/set/set.go#L98>)

```go
func (s Set[T]) Intersect(others ...Set[T]) Set[T]
```

**Intersect** returns a new [Set](<#Set>) containing the items of `s` present in every one of `others`.

<details><summary><b>Example</b></summary>

```go
a := set.New(1, 2, 3)
b := set.New(2, 3, 4)
fmt.Println(set.Sorted(a.Intersect(b)))
```

Output:

```text
[2 3]
```

</details>

<a name="Set.Len"></a>

### func (Set\[T\]) [Len](<https://github.com/gechr/x/blob/main/set/set.go#L52>)

```go
func (s Set[T]) Len() int
```

**Len** returns the number of items in `s`.

<a name="Set.Slice"></a>

### func (Set\[T\]) [Slice](<https://github.com/gechr/x/blob/main/set/set.go#L130>)

```go
func (s Set[T]) Slice() []T
```

**Slice** returns the items of `s` as a slice, in indeterminate order.

<a name="Set.SubsetOf"></a>

### func (Set\[T\]) [SubsetOf](<https://github.com/gechr/x/blob/main/set/set.go#L70>)

```go
func (s Set[T]) SubsetOf(other Set[T]) bool
```

**SubsetOf** returns whether every item in `s` is present in `other`.

<details><summary><b>Example</b></summary>

```go
a := set.New("x", "y")
b := set.New("x", "y", "z")
fmt.Println(a.SubsetOf(b))
fmt.Println(b.SubsetOf(a))
```

Output:

```text
true
false
```

</details>

<a name="Set.Union"></a>

### func (Set\[T\]) [Union](<https://github.com/gechr/x/blob/main/set/set.go#L83>)

```go
func (s Set[T]) Union(others ...Set[T]) Set[T]
```

**Union** returns a new [Set](<#Set>) containing the items of `s` and all `others`.

<details><summary><b>Example</b></summary>

```go
a := set.New(1, 2)
b := set.New(2, 3)
fmt.Println(set.Sorted(a.Union(b)))
```

Output:

```text
[1 2 3]
```

</details>

<a name="SortedSet"></a>

## type [SortedSet](<https://github.com/gechr/x/blob/main/set/sortedset.go#L17-L19>)

**SortedSet** is a set of ordered items, kept in ascending sorted order at all times: [SortedSet.Add](<#SortedSet.Add>) preserves sorted order, and combining sets ([SortedSet.Union](<#SortedSet.Union>)/[SortedSet.Intersect](<#SortedSet.Intersect>)/[SortedSet.Difference](<#SortedSet.Difference>)) always yields a sorted result. Unlike [Set](<#Set>), [SortedSet.Slice](<#SortedSet.Slice>) and [SortedSet.All](<#SortedSet.All>) iterate in deterministic ascending order rather than indeterminate map order.

The zero value is an empty, usable [SortedSet](<#SortedSet>).

```go
type SortedSet[T cmp.Ordered] struct {
    // contains filtered or unexported fields
}
```

<a name="CollectSorted"></a>

### func [CollectSorted](<https://github.com/gechr/x/blob/main/set/sortedset.go#L22>)

```go
func CollectSorted[T cmp.Ordered](seq iter.Seq[T]) SortedSet[T]
```

**CollectSorted** returns a [SortedSet](<#SortedSet>) containing the values of `seq`.

<details><summary><b>Example</b></summary>

**CollectSorted** builds a SortedSet from any iter.Seq, such as slices.Values.

```go
s := set.CollectSorted(slices.Values([]int{3, 1, 3, 2}))
fmt.Println(s.Slice())
```

Output:

```text
[1 2 3]
```

</details>

<a name="NewSorted"></a>

### func [NewSorted](<https://github.com/gechr/x/blob/main/set/sortedset.go#L28>)

```go
func NewSorted[T cmp.Ordered](items ...T) SortedSet[T]
```

**NewSorted** returns a [SortedSet](<#SortedSet>) containing `items`, sorted ascending with duplicates removed.

<details><summary><b>Example</b></summary>

SortedSet keeps its items in ascending order at all times, so iteration is deterministic.

```go
s := set.NewSorted(3, 1, 2, 1)
s.Add(0)
s.Delete(2)
fmt.Println(s.Slice())
```

Output:

```text
[0 1 3]
```

</details>

<a name="SortedSet.Add"></a>

### func (\*SortedSet\[T\]) [Add](<https://github.com/gechr/x/blob/main/set/sortedset.go#L33>)

```go
func (s *SortedSet[T]) Add(items ...T)
```

**Add** adds `items` to `s`, preserving sorted order and ignoring duplicates.

<a name="SortedSet.All"></a>

### func (SortedSet\[T\]) [All](<https://github.com/gechr/x/blob/main/set/sortedset.go#L41>)

```go
func (s SortedSet[T]) All() iter.Seq[T]
```

**All** returns an iterator over the items of `s`, in ascending order.

<details><summary><b>Example</b></summary>

```go
s := set.NewSorted("banana", "apple", "cherry")
for item := range s.All() {
    fmt.Println(item)
}
```

Output:

```text
apple
banana
cherry
```

</details>

<a name="SortedSet.Clone"></a>

### func (SortedSet\[T\]) [Clone](<https://github.com/gechr/x/blob/main/set/sortedset.go#L46>)

```go
func (s SortedSet[T]) Clone() SortedSet[T]
```

**Clone** returns a copy of `s`.

<a name="SortedSet.Contains"></a>

### func (SortedSet\[T\]) [Contains](<https://github.com/gechr/x/blob/main/set/sortedset.go#L51>)

```go
func (s SortedSet[T]) Contains(item T) bool
```

**Contains** returns whether `item` is present in `s`.

<a name="SortedSet.Delete"></a>

### func (\*SortedSet\[T\]) [Delete](<https://github.com/gechr/x/blob/main/set/sortedset.go#L57>)

```go
func (s *SortedSet[T]) Delete(items ...T)
```

**Delete** removes `items` from `s`.

<a name="SortedSet.Difference"></a>

### func (SortedSet\[T\]) [Difference](<https://github.com/gechr/x/blob/main/set/sortedset.go#L66>)

```go
func (s SortedSet[T]) Difference(others ...SortedSet[T]) SortedSet[T]
```

**Difference** returns a new [SortedSet](<#SortedSet>) containing the items of `s` not present in any of `others`.

<a name="SortedSet.Equal"></a>

### func (SortedSet\[T\]) [Equal](<https://github.com/gechr/x/blob/main/set/sortedset.go#L75>)

```go
func (s SortedSet[T]) Equal(other SortedSet[T]) bool
```

**Equal** returns whether `s` and `other` contain the same items.

<a name="SortedSet.Intersect"></a>

### func (SortedSet\[T\]) [Intersect](<https://github.com/gechr/x/blob/main/set/sortedset.go#L83>)

```go
func (s SortedSet[T]) Intersect(others ...SortedSet[T]) SortedSet[T]
```

**Intersect** returns a new [SortedSet](<#SortedSet>) containing the items of `s` present in every one of `others`.

<a name="SortedSet.Len"></a>

### func (SortedSet\[T\]) [Len](<https://github.com/gechr/x/blob/main/set/sortedset.go#L92>)

```go
func (s SortedSet[T]) Len() int
```

**Len** returns the number of items in `s`.

<a name="SortedSet.Slice"></a>

### func (SortedSet\[T\]) [Slice](<https://github.com/gechr/x/blob/main/set/sortedset.go#L97>)

```go
func (s SortedSet[T]) Slice() []T
```

**Slice** returns the items of `s` as a slice, in ascending order.

<a name="SortedSet.SubsetOf"></a>

### func (SortedSet\[T\]) [SubsetOf](<https://github.com/gechr/x/blob/main/set/sortedset.go#L102>)

```go
func (s SortedSet[T]) SubsetOf(other SortedSet[T]) bool
```

**SubsetOf** returns whether every item in `s` is present in `other`.

<a name="SortedSet.Union"></a>

### func (SortedSet\[T\]) [Union](<https://github.com/gechr/x/blob/main/set/sortedset.go#L120>)

```go
func (s SortedSet[T]) Union(others ...SortedSet[T]) SortedSet[T]
```

**Union** returns a new [SortedSet](<#SortedSet>) containing the items of `s` and all `others`.
