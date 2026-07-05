# ptr

```go
import "github.com/gechr/x/ptr"
```

Package ptr provides pointer helpers.

## Index

- [func Deref\[T any\](p \*T) T](<#Deref>)

<a name="Deref"></a>

## func [Deref](<https://github.com/gechr/x/blob/main/ptr/ptr.go#L5>)

```go
func Deref[T any](p *T) T
```

**Deref** returns the value `p` points to, or the zero value when `p` is nil.
