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

<a name="Clamp01"></a>

## func [Clamp01](<https://github.com/gechr/x/blob/main/math/clamp.go#L16>)

```go
func Clamp01(v float64) float64
```

**Clamp01** restricts `v` to the \[0, 1\] range. NaN clamps to 0.
