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

**Group** collects the pairs of `seq` into a map of slices, grouping values by key in encounter order.

<details><summary><b>Example</b></summary>

```go
words := []string{"apple", "banana", "avocado", "blueberry", "cherry"}
byInitial := xmaps.Group(func(yield func(byte, string) bool) {
    for _, w := range words {
        if !yield(w[0], w) {
            return
        }
    }
})
for initial, group := range xmaps.Sorted(byInitial) {
    fmt.Printf("%c: %v\n", initial, group)
}
```

Output:

```text
a: [apple avocado]
b: [banana blueberry]
c: [cherry]
```

</details>

<a name="GroupFunc"></a>

## func [GroupFunc](<https://github.com/gechr/x/blob/main/maps/group.go#L18>)

```go
func GroupFunc[K comparable, V any](seq iter.Seq[V], key func(V) K) map[K][]V
```

**GroupFunc** collects the values of `seq` into a map of slices, grouping values in encounter order by the key returned by `key`.

<details><summary><b>Example</b></summary>

```go
words := []string{"go", "rust", "zig", "java", "c"}
byLength := xmaps.GroupFunc(slices.Values(words), func(w string) int {
    return len(w)
})
for length, group := range xmaps.Sorted(byLength) {
    fmt.Println(length, group)
}
```

Output:

```text
1 [c]
2 [go]
3 [zig]
4 [rust java]
```

</details>

<a name="Invert"></a>

## func [Invert](<https://github.com/gechr/x/blob/main/maps/invert.go#L6>)

```go
func Invert[M ~map[K]V, K, V comparable](m M) map[V]K
```

**Invert** returns a new map with the keys and values of `m` swapped. If multiple keys map to the same value, exactly one of them survives as the value in the result, chosen arbitrarily due to map iteration order.

<details><summary><b>Example</b></summary>

```go
codes := map[string]int{"a": 1, "b": 2, "c": 3}
letters := xmaps.Invert(codes)
for code, letter := range xmaps.Sorted(letters) {
    fmt.Println(code, letter)
}
```

Output:

```text
1 a
2 b
3 c
```

</details>

<a name="KeysSlice"></a>

## func [KeysSlice](<https://github.com/gechr/x/blob/main/maps/keys.go#L4>)

```go
func KeysSlice[M ~map[K]V, K comparable, V any](m M) []K
```

**KeysSlice** returns the keys of `m` as a slice, in indeterminate order.

<details><summary><b>Example</b></summary>

```go
m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
keys := xmaps.KeysSlice(m)
slices.Sort(keys)
fmt.Println(strings.Join(keys, ", "))
```

Output:

```text
alpha, beta, charlie
```

</details>

<a name="Sorted"></a>

## func [Sorted](<https://github.com/gechr/x/blob/main/maps/sorted.go#L11>)

```go
func Sorted[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq2[K, V]
```

**Sorted** returns an iterator over the entries of `m` in ascending key order.

<details><summary><b>Example</b></summary>

```go
m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
for k, v := range xmaps.Sorted(m) {
    fmt.Println(k, v)
}
```

Output:

```text
alpha 1
beta 2
charlie 3
```

</details>

<a name="SortedFunc"></a>

## func [SortedFunc](<https://github.com/gechr/x/blob/main/maps/sorted.go#L23>)

```go
func SortedFunc[M ~map[K]V, K comparable, V any](m M, compare func(x, y K) int) iter.Seq2[K, V]
```

**SortedFunc** returns an iterator over the entries of `m` in the key order determined by `compare`, which follows the [cmp.Compare](<https://pkg.go.dev/cmp#Compare>) convention.

<details><summary><b>Example</b></summary>

**SortedFunc** accepts any comparison following the [cmp.Compare](<https://pkg.go.dev/cmp#Compare>) convention, such as a descending key order.

```go
m := map[int]string{1: "one", 2: "two", 3: "three"}
descending := func(x, y int) int { return cmp.Compare(y, x) }
for k, v := range xmaps.SortedFunc(m, descending) {
    fmt.Println(k, v)
}
```

Output:

```text
3 three
2 two
1 one
```

</details>

<a name="ValuesSlice"></a>

## func [ValuesSlice](<https://github.com/gechr/x/blob/main/maps/keys.go#L13>)

```go
func ValuesSlice[M ~map[K]V, K comparable, V any](m M) []V
```

**ValuesSlice** returns the values of `m` as a slice, in indeterminate order.

<details><summary><b>Example</b></summary>

```go
m := map[string]int{"charlie": 3, "alpha": 1, "beta": 2}
values := xmaps.ValuesSlice(m)
slices.Sort(values)
fmt.Println(values)
```

Output:

```text
[1 2 3]
```

</details>
