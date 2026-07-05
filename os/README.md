# os

```go
import "github.com/gechr/x/os"
```

Package `os` provides OS helpers: file probes, safe writes, copy, and line I/O.

## Index

- [func AtomicWrite(path string, data \[\]byte, perm os.FileMode) error](<#AtomicWrite>)
- [func CopyFile(src, dst string) error](<#CopyFile>)
- [func EnsureDir(dir string, perm os.FileMode) error](<#EnsureDir>)
- [func EnsureFile(path string, perm os.FileMode) error](<#EnsureFile>)
- [func Exists(path string) (bool, error)](<#Exists>)
- [func IsDir(path string) (bool, error)](<#IsDir>)
- [func IsFile(path string) (bool, error)](<#IsFile>)
- [func IsSymlink(path string) (bool, error)](<#IsSymlink>)
- [func IsWritableDir(dir string) bool](<#IsWritableDir>)
- [func ReadLines(path string) (\[\]string, error)](<#ReadLines>)
- [func SameFile(a, b string) (bool, error)](<#SameFile>)
- [func Trash(path string) error](<#Trash>)
- [func WriteLines(path string, lines \[\]string, perm os.FileMode) error](<#WriteLines>)

<a name="AtomicWrite"></a>

## func [AtomicWrite](<https://github.com/gechr/x/blob/main/os/write.go#L12>)

```go
func AtomicWrite(path string, data []byte, perm os.FileMode) error
```

**AtomicWrite** writes `data` to `path` via a temp-file-and-rename in the same directory. The temp file is removed on any failure.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "config.txt")
if err := xos.AtomicWrite(path, []byte("hello\n"), 0o600); err != nil {
    fmt.Println(err)
    return
}

data, _ := os.ReadFile(path)
fmt.Printf("%s", data)
```

Output:

```text
hello
```

</details>

<a name="CopyFile"></a>

## func [CopyFile](<https://github.com/gechr/x/blob/main/os/write.go#L88>)

```go
func CopyFile(src, dst string) error
```

**CopyFile** copies `src` to `dst`, preserving `src`'s mode bits. `dst` is fsynced before close. When `src` and `dst` are the same file (including via hard link) [CopyFile](<#CopyFile>) is a no-op.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

src := filepath.Join(dir, "src.txt")
dst := filepath.Join(dir, "dst.txt")
_ = os.WriteFile(src, []byte("hello\n"), 0o600)

if err := xos.CopyFile(src, dst); err != nil {
    fmt.Println(err)
    return
}

data, _ := os.ReadFile(dst)
fmt.Printf("%s", data)
```

Output:

```text
hello
```

</details>

<a name="EnsureDir"></a>

## func [EnsureDir](<https://github.com/gechr/x/blob/main/os/write.go#L50>)

```go
func EnsureDir(dir string, perm os.FileMode) error
```

**EnsureDir** creates `dir` and any missing parents, and guarantees `dir` itself has mode `perm` even if it already existed with a different mode or the umask interfered at creation time. Pre-existing parents are left untouched.

<details><summary><b>Example</b></summary>

**EnsureDir** creates missing parent directories, like mkdir -p.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

nested := filepath.Join(dir, "a", "b", "c")
if err := xos.EnsureDir(nested, 0o755); err != nil {
    fmt.Println(err)
    return
}

isDir, _ := xos.IsDir(nested)
fmt.Println(isDir)
```

Output:

```text
true
```

</details>

<a name="EnsureFile"></a>

## func [EnsureFile](<https://github.com/gechr/x/blob/main/os/write.go#L69>)

```go
func EnsureFile(path string, perm os.FileMode) error
```

**EnsureFile** creates `path` as an empty file with mode `perm` if it does not exist, creating any missing parent directories. An existing file's contents, mode, and timestamps are left untouched.

<details><summary><b>Example</b></summary>

**EnsureFile** creates the file and any missing parent directories.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "a", "b", "config.txt")
if err := xos.EnsureFile(path, 0o600); err != nil {
    fmt.Println(err)
    return
}

isFile, _ := xos.IsFile(path)
fmt.Println(isFile)
```

Output:

```text
true
```

</details>

<a name="Exists"></a>

## func [Exists](<https://github.com/gechr/x/blob/main/os/stat.go#L26>)

```go
func Exists(path string) (bool, error)
```

**Exists** reports whether `path` exists.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)

exists, _ := xos.Exists(path)
missing, _ := xos.Exists(filepath.Join(dir, "missing.txt"))
fmt.Println(exists)
fmt.Println(missing)
```

Output:

```text
true
false
```

</details>

<a name="IsDir"></a>

## func [IsDir](<https://github.com/gechr/x/blob/main/os/stat.go#L38>)

```go
func IsDir(path string) (bool, error)
```

**IsDir** reports whether `path` is a directory.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)

isDir, _ := xos.IsDir(dir)
notDir, _ := xos.IsDir(path)
fmt.Println(isDir)
fmt.Println(notDir)
```

Output:

```text
true
false
```

</details>

<a name="IsFile"></a>

## func [IsFile](<https://github.com/gechr/x/blob/main/os/stat.go#L32>)

```go
func IsFile(path string) (bool, error)
```

**IsFile** reports whether `path` is a regular file.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)

file, _ := xos.IsFile(path)
notFile, _ := xos.IsFile(dir)
fmt.Println(file)
fmt.Println(notFile)
```

Output:

```text
true
false
```

</details>

<a name="IsSymlink"></a>

## func [IsSymlink](<https://github.com/gechr/x/blob/main/os/stat.go#L44>)

```go
func IsSymlink(path string) (bool, error)
```

**IsSymlink** reports whether `path` is a symbolic link.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)
link := filepath.Join(dir, "link.txt")
_ = os.Symlink(path, link)

isLink, _ := xos.IsSymlink(link)
notLink, _ := xos.IsSymlink(path)
fmt.Println(isLink)
fmt.Println(notLink)
```

Output:

```text
true
false
```

</details>

<a name="IsWritableDir"></a>

## func [IsWritableDir](<https://github.com/gechr/x/blob/main/os/stat.go#L59>)

```go
func IsWritableDir(dir string) bool
```

**IsWritableDir** reports whether `dir` exists and the current process can create files in it. Uses a probe file rather than permission-bit inspection so that ACLs and immutable mounts are handled correctly.

<details><summary><b>Example</b></summary>

A missing path is not a writable directory.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

fmt.Println(xos.IsWritableDir(dir))
fmt.Println(xos.IsWritableDir(filepath.Join(dir, "missing")))
```

Output:

```text
true
false
```

</details>

<a name="ReadLines"></a>

## func [ReadLines](<https://github.com/gechr/x/blob/main/os/lines.go#L10>)

```go
func ReadLines(path string) ([]string, error)
```

**ReadLines** reads `path` and returns its non-empty, trimmed lines.

<details><summary><b>Example</b></summary>

**ReadLines** drops blank lines and trims surrounding whitespace.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "lines.txt")
_ = os.WriteFile(path, []byte("  alpha  \n\n\tbeta\n\ngamma\n"), 0o600)

lines, _ := xos.ReadLines(path)
for _, line := range lines {
    fmt.Println(line)
}
```

Output:

```text
alpha
beta
gamma
```

</details>

<a name="SameFile"></a>

## func [SameFile](<https://github.com/gechr/x/blob/main/os/file.go#L13>)

```go
func SameFile(a, b string) (bool, error)
```

**SameFile** reports whether `a` and `b` identify the same file. Missing leaf paths are compared after resolving their parent directories, and existing files are compared with [os.SameFile](<#SameFile>) to detect hard links.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)
link := filepath.Join(dir, "link.txt")
_ = os.Link(path, link)
other := filepath.Join(dir, "other.txt")
_ = os.WriteFile(other, []byte("hello"), 0o600)

same, _ := xos.SameFile(path, link)
different, _ := xos.SameFile(path, other)
fmt.Println(same)
fmt.Println(different)
```

Output:

```text
true
false
```

</details>

<a name="Trash"></a>

## func [Trash](<https://github.com/gechr/x/blob/main/os/trash.go#L24>)

```go
func Trash(path string) error
```

**Trash** asks the operating system to move `path` to its trash (or recycle bin) rather than removing it permanently like [os.Remove](<https://pkg.go.dev/os#Remove>), so it can typically be recovered. The `path` is resolved to an absolute path first, so a relative path trashes the intended file regardless of the working directory.

The mechanism is platform-specific: the system trash tool on macOS (so the Finder's "Put Back" works), the FreeDesktop.org trash specification on Linux and other Unix systems, and the shell file operation that targets the Recycle Bin on Windows. Recoverability is the OS's to honor, not a guarantee: an environment with the Recycle Bin disabled, for instance, may delete outright.

Where the platform cannot trash, it returns an error wrapping [errors.ErrUnsupported](<https://pkg.go.dev/errors#ErrUnsupported>), so a caller can detect the case and decide what to do (e.g. fall back to [os.Remove](<https://pkg.go.dev/os#Remove>)). This covers a macOS older than 15 (which lacks the system trash tool) and a Unix file with no usable same-device trash.

<a name="WriteLines"></a>

## func [WriteLines](<https://github.com/gechr/x/blob/main/os/lines.go#L27>)

```go
func WriteLines(path string, lines []string, perm os.FileMode) error
```

**WriteLines** atomically writes `lines` to `path`, one per line, with a trailing newline.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "lines.txt")
if err := xos.WriteLines(path, []string{"alpha", "beta"}, 0o600); err != nil {
    fmt.Println(err)
    return
}

data, _ := os.ReadFile(path)
fmt.Printf("%q\n", data)
```

Output:

```text
"alpha\nbeta\n"
```

</details>
