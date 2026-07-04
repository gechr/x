# maps

```go
import "github.com/gechr/x/maps"
```

Package maps provides map helpers: sorted iteration, grouping, and inversion.

## Index

- [func Group\[K comparable, V any\](seq iter.Seq2\[K, V\]) map\[K\]\[\]V](<#Group>)
- [func GroupFunc\[K comparable, V any\](seq iter.Seq\[V\], key func(V) K) map\[K\]\[\]V](<#GroupFunc>)
- [func Invert\[M ~map\[K\]V, K, V comparable\](m M) map\[V\]K](<#Invert>)
- [func KeysSlice\[M ~map\[K\]V, K comparable, V any\](m M) \[\]K](<#KeysSlice>)
- [func Sorted\[M ~map\[K\]V, K cmp.Ordered, V any\](m M) iter.Seq2\[K, V\]](<#Sorted>)
- [func SortedFunc\[M ~map\[K\]V, K comparable, V any\](m M, compare func(x, y K) int) iter.Seq2\[K, V\]](<#SortedFunc>)
- [func ValuesSlice\[M ~map\[K\]V, K comparable, V any\](m M) \[\]V](<#ValuesSlice>)

<a name="Group"></a>

## func [Group](<https://github.com/gechr/x/blob/main/maps/group.go#L8>)

```go
func Group[K comparable, V any](seq iter.Seq2[K, V]) map[K][]V
```

**Group** collects the pairs of seq into a map of slices, grouping values by key in encounter order.

<a name="GroupFunc"></a>

## func [GroupFunc](<https://github.com/gechr/x/blob/main/maps/group.go#L18>)

```go
func GroupFunc[K comparable, V any](seq iter.Seq[V], key func(V) K) map[K][]V
```

**GroupFunc** collects the values of seq into a map of slices, grouping values in encounter order by the key returned by key.

<a name="Invert"></a>

## func [Invert](<https://github.com/gechr/x/blob/main/maps/invert.go#L6>)

```go
func Invert[M ~map[K]V, K, V comparable](m M) map[V]K
```

**Invert** returns a new map with the keys and values of m swapped. If multiple keys map to the same value, exactly one of them survives as the value in the result, chosen arbitrarily due to map iteration order.

<a name="KeysSlice"></a>

## func [KeysSlice](<https://github.com/gechr/x/blob/main/maps/keys.go#L4>)

```go
func KeysSlice[M ~map[K]V, K comparable, V any](m M) []K
```

**KeysSlice** returns the keys of m as a slice, in indeterminate order.

<a name="Sorted"></a>

## func [Sorted](<https://github.com/gechr/x/blob/main/maps/sorted.go#L11>)

```go
func Sorted[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V]
```

**Sorted** returns an iterator over the entries of m in ascending key order.

<a name="SortedFunc"></a>

## func [SortedFunc](<https://github.com/gechr/x/blob/main/maps/sorted.go#L23>)

```go
func SortedFunc[M ~map[K]V, K comparable, V any](m M, compare func(x, y K) int) iter.Seq2[K, V]
```

**SortedFunc** returns an iterator over the entries of m in the key order determined by compare, which follows the [cmp.Compare](<https://pkg.go.dev/cmp#Compare>) convention.

<a name="ValuesSlice"></a>

## func [ValuesSlice](<https://github.com/gechr/x/blob/main/maps/keys.go#L13>)

```go
func ValuesSlice[M ~map[K]V, K comparable, V any](m M) []V
```

**ValuesSlice** returns the values of m as a slice, in indeterminate order.
