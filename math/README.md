# math

```go
import "github.com/gechr/x/math"
```

Package math provides numeric helpers.

## Index

- [func Clamp\[T cmp.Ordered\](v, lo, hi T) T](<#Clamp>)
- [func Clamp01(v float64) float64](<#Clamp01>)

<a name="Clamp"></a>

## func [Clamp](<https://github.com/gechr/x/blob/main/math/clamp.go#L8>)

```go
func Clamp[T cmp.Ordered](v, lo, hi T) T
```

**Clamp** restricts `v` to the \[`lo`, `hi`\] range. NaN clamps to `lo`; infinities clamp to the nearest bound. Unlike min(max(v, lo), hi), NaN does not propagate.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xmath.Clamp(5, 0, 10))
fmt.Println(xmath.Clamp(-3, 0, 10))
fmt.Println(xmath.Clamp(42, 0, 10))
```

Output:

```text
5
0
10
```

</details>

<details><summary><b>Example (Nan)</b></summary>

Unlike min(max(v, lo), hi), NaN does not propagate - it clamps to `lo`.

```go
fmt.Println(xmath.Clamp(math.NaN(), 1.0, 2.0))
fmt.Println(min(max(math.NaN(), 1.0), 2.0))
```

Output:

```text
1
NaN
```

</details>

<details><summary><b>Example (Strings)</b></summary>

**Clamp** works with any ordered type, including strings.

```go
fmt.Println(xmath.Clamp("a", "b", "d"))
fmt.Println(xmath.Clamp("c", "b", "d"))
fmt.Println(xmath.Clamp("z", "b", "d"))
```

Output:

```text
b
c
d
```

</details>

<a name="Clamp01"></a>

## func [Clamp01](<https://github.com/gechr/x/blob/main/math/clamp.go#L16>)

```go
func Clamp01(v float64) float64
```

**Clamp01** restricts `v` to the \[0, 1\] range. NaN clamps to 0.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xmath.Clamp01(0.5))
fmt.Println(xmath.Clamp01(-0.1))
fmt.Println(xmath.Clamp01(1.1))
```

Output:

```text
0.5
0
1
```

</details>
