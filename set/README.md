# set

```go
import "github.com/gechr/x/set"
```

Package set provides a generic set backed by a map.

## Index

- [func Sorted\[T cmp.Ordered\]\(s Set\[T\]\) \[\]T](<#Sorted>)
- [func SortedNatural\[T \~string\]\(s Set\[T\]\) \[\]T](<#SortedNatural>)
- [type Set](<#Set>)
  - [func Collect\[T comparable\]\(seq iter.Seq\[T\]\) Set\[T\]](<#Collect>)
  - [func New\[T comparable\]\(items ...T\) Set\[T\]](<#New>)
  - [func \(s Set\[T\]\) Add\(items ...T\)](<#Set[T].Add>)
  - [func \(s Set\[T\]\) All\(\) iter.Seq\[T\]](<#Set[T].All>)
  - [func \(s Set\[T\]\) Clone\(\) Set\[T\]](<#Set[T].Clone>)
  - [func \(s Set\[T\]\) Contains\(item T\) bool](<#Set[T].Contains>)
  - [func \(s Set\[T\]\) Delete\(items ...T\)](<#Set[T].Delete>)
  - [func \(s Set\[T\]\) Difference\(others ...Set\[T\]\) Set\[T\]](<#Set[T].Difference>)
  - [func \(s Set\[T\]\) Equal\(other Set\[T\]\) bool](<#Set[T].Equal>)
  - [func \(s Set\[T\]\) Intersect\(others ...Set\[T\]\) Set\[T\]](<#Set[T].Intersect>)
  - [func \(s Set\[T\]\) Len\(\) int](<#Set[T].Len>)
  - [func \(s Set\[T\]\) Slice\(\) \[\]T](<#Set[T].Slice>)
  - [func \(s Set\[T\]\) SubsetOf\(other Set\[T\]\) bool](<#Set[T].SubsetOf>)
  - [func \(s Set\[T\]\) Union\(others ...Set\[T\]\) Set\[T\]](<#Set[T].Union>)
- [type SortedSet](<#SortedSet>)
  - [func CollectSorted\[T cmp.Ordered\]\(seq iter.Seq\[T\]\) SortedSet\[T\]](<#CollectSorted>)
  - [func NewSorted\[T cmp.Ordered\]\(items ...T\) SortedSet\[T\]](<#NewSorted>)
  - [func \(s \*SortedSet\[T\]\) Add\(items ...T\)](<#SortedSet[T].Add>)
  - [func \(s SortedSet\[T\]\) All\(\) iter.Seq\[T\]](<#SortedSet[T].All>)
  - [func \(s SortedSet\[T\]\) Clone\(\) SortedSet\[T\]](<#SortedSet[T].Clone>)
  - [func \(s SortedSet\[T\]\) Contains\(item T\) bool](<#SortedSet[T].Contains>)
  - [func \(s \*SortedSet\[T\]\) Delete\(items ...T\)](<#SortedSet[T].Delete>)
  - [func \(s SortedSet\[T\]\) Difference\(others ...SortedSet\[T\]\) SortedSet\[T\]](<#SortedSet[T].Difference>)
  - [func \(s SortedSet\[T\]\) Equal\(other SortedSet\[T\]\) bool](<#SortedSet[T].Equal>)
  - [func \(s SortedSet\[T\]\) Intersect\(others ...SortedSet\[T\]\) SortedSet\[T\]](<#SortedSet[T].Intersect>)
  - [func \(s SortedSet\[T\]\) Len\(\) int](<#SortedSet[T].Len>)
  - [func \(s SortedSet\[T\]\) Slice\(\) \[\]T](<#SortedSet[T].Slice>)
  - [func \(s SortedSet\[T\]\) SubsetOf\(other SortedSet\[T\]\) bool](<#SortedSet[T].SubsetOf>)
  - [func \(s SortedSet\[T\]\) Union\(others ...SortedSet\[T\]\) SortedSet\[T\]](<#SortedSet[T].Union>)

<a name="Sorted"></a>

## func [Sorted](<https://github.com/gechr/x/blob/main/set/sorted.go#L14>)

```go
func Sorted[T cmp.Ordered](s Set[T]) []T
```

Sorted returns the items of s as a slice in ascending order.

Sorted is a function rather than a [Set](<#Set>) method because it requires T to be ordered, not just comparable.

<a name="SortedNatural"></a>

## func [SortedNatural](<https://github.com/gechr/x/blob/main/set/sorted.go#L26>)

```go
func SortedNatural[T ~string](s Set[T]) []T
```

SortedNatural returns the items of s as a slice in natural order, so embedded numbers compare by value \("item2" before "item10"\) rather than lexically. See \[xstrings.CompareNatural\].

SortedNatural is a function rather than a [Set](<#Set>) method because it requires T to be string\-like, not just comparable.

<a name="Set"></a>

## type [Set](<https://github.com/gechr/x/blob/main/set/set.go#L13>)

Set is a set of comparable items backed by a map. Pointer types and structs containing pointer fields are compared using shallow equality. The zero value is nil: read operations work, but Add panics \- use [New](<#New>) or [Collect](<#Collect>) to create a usable Set.

```go
type Set[T comparable] map[T]struct{}
```

<a name="Collect"></a>

### func [Collect](<https://github.com/gechr/x/blob/main/set/set.go#L23>)

```go
func Collect[T comparable](seq iter.Seq[T]) Set[T]
```

Collect returns a Set containing the values of seq.

<a name="New"></a>

### func [New](<https://github.com/gechr/x/blob/main/set/set.go#L16>)

```go
func New[T comparable](items ...T) Set[T]
```

New returns a Set containing items.

<a name="Set[T].Add"></a>

### func \(Set\[T\]\) [Add](<https://github.com/gechr/x/blob/main/set/set.go#L32>)

```go
func (s Set[T]) Add(items ...T)
```

Add adds items to s.

<a name="Set[T].All"></a>

### func \(Set\[T\]\) [All](<https://github.com/gechr/x/blob/main/set/set.go#L139>)

```go
func (s Set[T]) All() iter.Seq[T]
```

All returns an iterator over the items of s, in indeterminate order.

<a name="Set[T].Clone"></a>

### func \(Set\[T\]\) [Clone](<https://github.com/gechr/x/blob/main/set/set.go#L121>)

```go
func (s Set[T]) Clone() Set[T]
```

Clone returns a copy of s.

<a name="Set[T].Contains"></a>

### func \(Set\[T\]\) [Contains](<https://github.com/gechr/x/blob/main/set/set.go#L46>)

```go
func (s Set[T]) Contains(item T) bool
```

Contains returns whether item is present in s.

<a name="Set[T].Delete"></a>

### func \(Set\[T\]\) [Delete](<https://github.com/gechr/x/blob/main/set/set.go#L39>)

```go
func (s Set[T]) Delete(items ...T)
```

Delete removes items from s.

<a name="Set[T].Difference"></a>

### func \(Set\[T\]\) [Difference](<https://github.com/gechr/x/blob/main/set/set.go#L110>)

```go
func (s Set[T]) Difference(others ...Set[T]) Set[T]
```

Difference returns a new Set containing the items of s not present in any of others.

<a name="Set[T].Equal"></a>

### func \(Set\[T\]\) [Equal](<https://github.com/gechr/x/blob/main/set/set.go#L57>)

```go
func (s Set[T]) Equal(other Set[T]) bool
```

Equal returns whether s and other contain the same items.

<a name="Set[T].Intersect"></a>

### func \(Set\[T\]\) [Intersect](<https://github.com/gechr/x/blob/main/set/set.go#L98>)

```go
func (s Set[T]) Intersect(others ...Set[T]) Set[T]
```

Intersect returns a new Set containing the items of s present in every one of others.

<a name="Set[T].Len"></a>

### func \(Set\[T\]\) [Len](<https://github.com/gechr/x/blob/main/set/set.go#L52>)

```go
func (s Set[T]) Len() int
```

Len returns the number of items in s.

<a name="Set[T].Slice"></a>

### func \(Set\[T\]\) [Slice](<https://github.com/gechr/x/blob/main/set/set.go#L130>)

```go
func (s Set[T]) Slice() []T
```

Slice returns the items of s as a slice, in indeterminate order.

<a name="Set[T].SubsetOf"></a>

### func \(Set\[T\]\) [SubsetOf](<https://github.com/gechr/x/blob/main/set/set.go#L70>)

```go
func (s Set[T]) SubsetOf(other Set[T]) bool
```

SubsetOf returns whether every item in s is present in other.

<a name="Set[T].Union"></a>

### func \(Set\[T\]\) [Union](<https://github.com/gechr/x/blob/main/set/set.go#L83>)

```go
func (s Set[T]) Union(others ...Set[T]) Set[T]
```

Union returns a new Set containing the items of s and all others.

<a name="SortedSet"></a>

## type [SortedSet](<https://github.com/gechr/x/blob/main/set/sortedset.go#L16-L18>)

SortedSet is a set of ordered items, kept in ascending sorted order at all times: Add inserts in sorted position, and combining sets \(Union/Intersect/Difference\) always yields a sorted result. Unlike [Set](<#Set>), Slice and All iterate in deterministic ascending order rather than indeterminate map order.

The zero value is an empty, usable SortedSet.

```go
type SortedSet[T cmp.Ordered] struct {
    // contains filtered or unexported fields
}
```

<a name="CollectSorted"></a>

### func [CollectSorted](<https://github.com/gechr/x/blob/main/set/sortedset.go#L29>)

```go
func CollectSorted[T cmp.Ordered](seq iter.Seq[T]) SortedSet[T]
```

CollectSorted returns a SortedSet containing the values of seq.

<a name="NewSorted"></a>

### func [NewSorted](<https://github.com/gechr/x/blob/main/set/sortedset.go#L22>)

```go
func NewSorted[T cmp.Ordered](items ...T) SortedSet[T]
```

NewSorted returns a SortedSet containing items, sorted ascending with duplicates removed.

<a name="SortedSet[T].Add"></a>

### func \(\*SortedSet\[T\]\) [Add](<https://github.com/gechr/x/blob/main/set/sortedset.go#L39>)

```go
func (s *SortedSet[T]) Add(items ...T)
```

Add adds items to s, inserting each in sorted position and ignoring duplicates.

<a name="SortedSet[T].All"></a>

### func \(SortedSet\[T\]\) [All](<https://github.com/gechr/x/blob/main/set/sortedset.go#L125>)

```go
func (s SortedSet[T]) All() iter.Seq[T]
```

All returns an iterator over the items of s, in ascending order.

<a name="SortedSet[T].Clone"></a>

### func \(SortedSet\[T\]\) [Clone](<https://github.com/gechr/x/blob/main/set/sortedset.go#L115>)

```go
func (s SortedSet[T]) Clone() SortedSet[T]
```

Clone returns a copy of s.

<a name="SortedSet[T].Contains"></a>

### func \(SortedSet\[T\]\) [Contains](<https://github.com/gechr/x/blob/main/set/sortedset.go#L58>)

```go
func (s SortedSet[T]) Contains(item T) bool
```

Contains returns whether item is present in s.

<a name="SortedSet[T].Delete"></a>

### func \(\*SortedSet\[T\]\) [Delete](<https://github.com/gechr/x/blob/main/set/sortedset.go#L49>)

```go
func (s *SortedSet[T]) Delete(items ...T)
```

Delete removes items from s.

<a name="SortedSet[T].Difference"></a>

### func \(SortedSet\[T\]\) [Difference](<https://github.com/gechr/x/blob/main/set/sortedset.go#L91>)

```go
func (s SortedSet[T]) Difference(others ...SortedSet[T]) SortedSet[T]
```

Difference returns a new SortedSet containing the items of s not present in any of others.

<a name="SortedSet[T].Equal"></a>

### func \(SortedSet\[T\]\) [Equal](<https://github.com/gechr/x/blob/main/set/sortedset.go#L69>)

```go
func (s SortedSet[T]) Equal(other SortedSet[T]) bool
```

Equal returns whether s and other contain the same items.

<a name="SortedSet[T].Intersect"></a>

### func \(SortedSet\[T\]\) [Intersect](<https://github.com/gechr/x/blob/main/set/sortedset.go#L85>)

```go
func (s SortedSet[T]) Intersect(others ...SortedSet[T]) SortedSet[T]
```

Intersect returns a new SortedSet containing the items of s present in every one of others.

<a name="SortedSet[T].Len"></a>

### func \(SortedSet\[T\]\) [Len](<https://github.com/gechr/x/blob/main/set/sortedset.go#L64>)

```go
func (s SortedSet[T]) Len() int
```

Len returns the number of items in s.

<a name="SortedSet[T].Slice"></a>

### func \(SortedSet\[T\]\) [Slice](<https://github.com/gechr/x/blob/main/set/sortedset.go#L120>)

```go
func (s SortedSet[T]) Slice() []T
```

Slice returns the items of s as a slice, in ascending order.

<a name="SortedSet[T].SubsetOf"></a>

### func \(SortedSet\[T\]\) [SubsetOf](<https://github.com/gechr/x/blob/main/set/sortedset.go#L74>)

```go
func (s SortedSet[T]) SubsetOf(other SortedSet[T]) bool
```

SubsetOf returns whether every item in s is present in other.

<a name="SortedSet[T].Union"></a>

### func \(SortedSet\[T\]\) [Union](<https://github.com/gechr/x/blob/main/set/sortedset.go#L79>)

```go
func (s SortedSet[T]) Union(others ...SortedSet[T]) SortedSet[T]
```

Union returns a new SortedSet containing the items of s and all others.
