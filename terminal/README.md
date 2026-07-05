# terminal

```go
import "github.com/gechr/x/terminal"
```

Package terminal provides terminal detection and size queries.

## Index

- [func Height(f \*os.File) int](<#Height>)
- [func Is(f \*os.File) bool](<#Is>)
- [func Size(f \*os.File) (int, int)](<#Size>)
- [func Width(f \*os.File) int](<#Width>)

<a name="Height"></a>

## func [Height](<https://github.com/gechr/x/blob/main/terminal/terminal.go#L28>)

```go
func Height(f *os.File) int
```

**Height** returns the height of the terminal connected to `f`, or 0 if `f` is nil or not a terminal.

<details><summary><b>Example</b></summary>

**Height** returns 0 when the file is nil or not connected to a terminal.

```go
fmt.Println(terminal.Height(nil))
```

Output:

```text
0
```

</details>

<a name="Is"></a>

## func [Is](<https://github.com/gechr/x/blob/main/terminal/terminal.go#L12>)

```go
func Is(f *os.File) bool
```

**Is** returns true if the given file is a terminal. Returns false for nil files.

<details><summary><b>Example</b></summary>

A pipe is not a terminal, and nil files are always reported as non-terminals.

```go
r, w, err := os.Pipe()
if err != nil {
    panic(err)
}
defer r.Close()
defer w.Close()

fmt.Println(terminal.Is(r))
fmt.Println(terminal.Is(w))
fmt.Println(terminal.Is(nil))
```

Output:

```text
false
false
false
```

</details>

<a name="Size"></a>

## func [Size](<https://github.com/gechr/x/blob/main/terminal/terminal.go#L35>)

```go
func Size(f *os.File) (int, int)
```

**Size** returns the (width, height) of the terminal connected to `f` in cells, or (0, 0) if `f` is nil or not a terminal.

<details><summary><b>Example</b></summary>

**Size** returns (0, 0) when the file is nil or not connected to a terminal.

```go
w, h := terminal.Size(nil)
fmt.Println(w, h)
```

Output:

```text
0 0
```

</details>

<a name="Width"></a>

## func [Width](<https://github.com/gechr/x/blob/main/terminal/terminal.go#L21>)

```go
func Width(f *os.File) int
```

**Width** returns the width of the terminal connected to `f`. Returns 0 if `f` is nil or not a terminal.

<details><summary><b>Example</b></summary>

**Width** returns 0 when the file is nil or not connected to a terminal.

```go
fmt.Println(terminal.Width(nil))
```

Output:

```text
0
```

</details>
