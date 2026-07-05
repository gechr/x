# ptr

```go
import "github.com/gechr/x/ptr"
```

Package `ptr` provides pointer helpers.

## Index

- [func Deref\[T any\](p \*T) T](<#Deref>)

<a name="Deref"></a>

## func [Deref](<https://github.com/gechr/x/blob/main/ptr/ptr.go#L5>)

```go
func Deref[T any](p *T) T
```

**Deref** returns the value `p` points to, or the zero value when `p` is nil.

<details><summary><b>Example</b></summary>

```go
s := "hello"
fmt.Println(ptr.Deref(&s))
```

Output:

```text
hello
```

</details>

<details><summary><b>Example (Nil)</b></summary>

A nil pointer dereferences to the zero value instead of panicking.

```go
var s *string
var n *int
fmt.Printf("%q\n", ptr.Deref(s))
fmt.Println(ptr.Deref(n))
```

Output:

```text
""
0
```

</details>

<details><summary><b>Example (OptionalField)</b></summary>

**Deref** is handy for optional struct fields modelled as pointers.

```go
type Config struct {
    Retries *int
}

retries := 3
fmt.Println(ptr.Deref(Config{Retries: &retries}.Retries))
fmt.Println(ptr.Deref(Config{}.Retries))
```

Output:

```text
3
0
```

</details>
