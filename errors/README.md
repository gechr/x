# errors

```go
import "github.com/gechr/x/errors"
```

Package `errors` provides helpers for working with errors.

## Index

- [func IsAny(err error, targets ...error) bool](<#IsAny>)

<a name="IsAny"></a>

## func [IsAny](<https://github.com/gechr/x/blob/main/errors/errors.go#L8>)

```go
func IsAny(err error, targets ...error) bool
```

**IsAny** reports whether any error in `targets` matches `err`'s error tree. Each target is compared using [errors.Is](<https://pkg.go.dev/errors#Is>).

<details><summary><b>Example</b></summary>

```go
errNotFound := errors.New("not found")
errUnavailable := errors.New("unavailable")
err := fmt.Errorf("lookup: %w", errNotFound)

fmt.Println(xerrors.IsAny(err, errNotFound, errUnavailable))
```

Output:

```text
true
```

</details>
