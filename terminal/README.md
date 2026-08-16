# terminal

```go
import "github.com/gechr/x/terminal"
```

Package `terminal` provides terminal detection and size queries.

## Index

- [func Height(f \*os.File) int](<#Height>)
- [func Is(f \*os.File) bool](<#Is>)
- [func IsDark() (bool, bool)](<#IsDark>)
- [func IsLight() (bool, bool)](<#IsLight>)
- [func Size(f \*os.File) (int, int)](<#Size>)
- [func SupportsTrueColor() bool](<#SupportsTrueColor>)
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

<a name="IsDark"></a>

## func [IsDark](<https://github.com/gechr/x/blob/main/terminal/background.go#L32>)

```go
func IsDark() (bool, bool)
```

**IsDark** reports (dark, ok) for the controlling terminal. It performs terminal I/O on the first call, returning as soon as the terminal has answered and waiting no longer than half a second for one that never does. The result, including no response, is cached for the process. `ok` is false if no standard stream is a terminal or the terminal does not respond, in which case the first result is meaningless.

<details><summary><b>Example</b></summary>

**IsDark** reports ok=false when no standard stream is connected to a terminal, since there is no background to query.

```go
dark, ok := terminal.IsDark()
switch {
case !ok:
    fmt.Println("no terminal detected")
case dark:
    fmt.Println("dark")
default:
    fmt.Println("light")
}
```

Output:

```text
no terminal detected
```

</details>

<a name="IsLight"></a>

## func [IsLight](<https://github.com/gechr/x/blob/main/terminal/background.go#L41>)

```go
func IsLight() (bool, bool)
```

**IsLight** reports (light, ok) for the controlling terminal. Like [IsDark](<#IsDark>), it performs terminal I/O on the first call and caches the result for the process. `ok` is false if no standard stream is a terminal or the terminal does not respond, in which case the first result is meaningless.

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

<a name="SupportsTrueColor"></a>

## func [SupportsTrueColor](<https://github.com/gechr/x/blob/main/terminal/color.go#L37>)

```go
func SupportsTrueColor() bool
```

**SupportsTrueColor** reports whether the terminal supports 24-bit "true color" output.

Detection reads COLORTERM, TERM, the terminfo Tc and RGB capabilities and, inside tmux, `tmux info` - tmux does not forward COLORTERM, so only its own capabilities settle the question there. TERM alone is not trusted for capability: the ubiquitous TERM=xterm-256color advertises only 256 colors even on terminals that render 24-bit, which is precisely why COLORTERM exists.

This reports capability, not preference: NO\_COLOR and CLICOLOR are ignored and a redirected stream is still measured against the terminal, so honoring either is left to the caller. The first call detects, and the result is cached for the process.

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
