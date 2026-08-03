# os

```go
import "github.com/gechr/x/os"
```

Package `os` provides OS helpers: file probes, safe writes, copy, line I/O, and platform/architecture detection.

## Index

- [Constants](<#constants>)
- [func AtomicWrite(path string, data \[\]byte, perm os.FileMode) error](<#AtomicWrite>)
- [func CopyFile(src, dst string) error](<#CopyFile>)
- [func EnsureDir(dir string, perm os.FileMode) error](<#EnsureDir>)
- [func EnsureFile(path string, perm os.FileMode) error](<#EnsureFile>)
- [func Exists(path string) (bool, error)](<#Exists>)
- [func IsAndroid() bool](<#IsAndroid>)
- [func IsBSD() bool](<#IsBSD>)
- [func IsDarwin() bool](<#IsDarwin>)
- [func IsDir(path string) (bool, error)](<#IsDir>)
- [func IsExecutable(path string) (bool, error)](<#IsExecutable>)
- [func IsFile(path string) (bool, error)](<#IsFile>)
- [func IsIOS() bool](<#IsIOS>)
- [func IsLinux() bool](<#IsLinux>)
- [func IsSymlink(path string) (bool, error)](<#IsSymlink>)
- [func IsUnix() bool](<#IsUnix>)
- [func IsWasm() bool](<#IsWasm>)
- [func IsWindows() bool](<#IsWindows>)
- [func IsWritableDir(dir string) bool](<#IsWritableDir>)
- [func ReadLines(path string) (\[\]string, error)](<#ReadLines>)
- [func RemoveIfExists(path string) error](<#RemoveIfExists>)
- [func SameFile(a, b string) (bool, error)](<#SameFile>)
- [func Trash(path string) error](<#Trash>)
- [func WriteLines(path string, lines \[\]string, perm os.FileMode) error](<#WriteLines>)

## Constants

<a name="Arch386"></a>Arch constants are the recognized [runtime.GOARCH](<https://pkg.go.dev/runtime#GOARCH>) values. Go exposes GOARCH only as a string, so these name the tokens to avoid scattering string literals across build-time comparisons.

```go
const (
    Arch386      = "386"
    ArchAMD64    = "amd64"
    ArchARM      = "arm"
    ArchARM64    = "arm64"
    ArchLoong64  = "loong64"
    ArchMIPS     = "mips"
    ArchMIPS64   = "mips64"
    ArchMIPS64LE = "mips64le"
    ArchMIPSLE   = "mipsle"
    ArchPPC64    = "ppc64"
    ArchPPC64LE  = "ppc64le"
    ArchRISCV64  = "riscv64"
    ArchS390X    = "s390x"
    ArchWASM     = "wasm"
)
```

<a name="PlatformAIX"></a>Platform constants are the recognized [runtime.GOOS](<https://pkg.go.dev/runtime#GOOS>) values. Go exposes GOOS only as a string, so these name the tokens to avoid scattering string literals across build-time comparisons.

```go
const (
    PlatformAIX       = "aix"
    PlatformAndroid   = "android"
    PlatformDarwin    = "darwin"
    PlatformDragonfly = "dragonfly"
    PlatformFreeBSD   = "freebsd"
    PlatformIllumos   = "illumos"
    PlatformIOS       = "ios"
    PlatformJS        = "js"
    PlatformLinux     = "linux"
    PlatformNetBSD    = "netbsd"
    PlatformOpenBSD   = "openbsd"
    PlatformPlan9     = "plan9"
    PlatformSolaris   = "solaris"
    PlatformWASIP1    = "wasip1"
    PlatformWindows   = "windows"
)
```

<a name="AtomicWrite"></a>

## func [AtomicWrite](<https://github.com/gechr/x/blob/main/os/write.go#L27>)

```go
func AtomicWrite(path string, data []byte, perm os.FileMode) error
```

**AtomicWrite** writes `data` to `path` via a temp-file-and-rename in the same directory. The temp file is removed on any failure.

The replacement is durable, not merely atomic: the temp file's contents are synced before the rename and `path`'s directory is synced after it, so the new directory entry survives a crash rather than being lost to unflushed metadata. Directory syncing is Unix-only and a no-op elsewhere (see [IsUnix](<#IsUnix>)). A failing directory sync is reported as an error even though the rename has already landed - `data` is at `path`, only its durability is unproven - and `path` is deliberately left as written rather than rolled back.

A symlink at `path` is replaced by the new file, not written through to its target, because [os.Rename](<https://pkg.go.dev/os#Rename>) acts on the name. That is what keeps the write atomic, since a link's target may sit on another filesystem where no rename is possible. Callers wanting to write through a link must resolve it first (see [filepath.Resolve](<../filepath/README.md#Resolve>)) and accept the target's directory as the one being modified.

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

## func [CopyFile](<https://github.com/gechr/x/blob/main/os/write.go#L110>)

```go
func CopyFile(src, dst string) error
```

**CopyFile** copies `src` to `dst`, preserving `src`'s mode bits. `dst` is fsynced before close and its directory afterwards, so a newly created `dst` is durable and not just written - the same reasoning as [AtomicWrite](<#AtomicWrite>), and likewise Unix-only. When `src` and `dst` are the same file (including via hard link) [CopyFile](<#CopyFile>) is a no-op.

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

## func [EnsureDir](<https://github.com/gechr/x/blob/main/os/write.go#L70>)

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

## func [EnsureFile](<https://github.com/gechr/x/blob/main/os/write.go#L89>)

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

## func [Exists](<https://github.com/gechr/x/blob/main/os/stat.go#L28>)

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

<a name="IsAndroid"></a>

## func [IsAndroid](<https://github.com/gechr/x/blob/main/os/platform.go#L44>)

```go
func IsAndroid() bool
```

**IsAndroid** reports whether the program is running on Android. Android is its own GOOS but is also Unix-like, so it additionally satisfies [IsUnix](<#IsUnix>).

<a name="IsBSD"></a>

## func [IsBSD](<https://github.com/gechr/x/blob/main/os/platform.go#L60>)

```go
func IsBSD() bool
```

**IsBSD** reports whether the program is running on a BSD-family OS: FreeBSD, NetBSD, OpenBSD, or DragonFly BSD. Go has no `bsd` build constraint, so this is a fixed GOOS set. macOS is deliberately excluded even though Darwin is BSD-derived; use [IsDarwin](<#IsDarwin>) for it. Every OS reported here is also Unix-like, so it satisfies [IsUnix](<#IsUnix>).

<a name="IsDarwin"></a>

## func [IsDarwin](<https://github.com/gechr/x/blob/main/os/platform.go#L33>)

```go
func IsDarwin() bool
```

**IsDarwin** reports whether the program is running on macOS. It matches Go's `darwin` GOOS token; iOS is a separate GOOS (see [IsIOS](<#IsIOS>)) and reports false.

<a name="IsDir"></a>

## func [IsDir](<https://github.com/gechr/x/blob/main/os/stat.go#L40>)

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

<a name="IsExecutable"></a>

## func [IsExecutable](<https://github.com/gechr/x/blob/main/os/stat.go#L66>)

```go
func IsExecutable(path string) (bool, error)
```

**IsExecutable** reports whether `path`, with every symlink resolved, is a regular file that the current process can run as a binary. It answers the practical question rather than merely inspecting the permission bits: on Unix via `access(2)` with `X_OK` (so the owner/group/other bit that actually applies to this process is used), and on Windows via the resolved file's extension appearing in `%PATHEXT%` (Windows has no execute bit). A non-existent path reports false; a directory is traversable, not runnable, so it also reports false.

<a name="IsFile"></a>

## func [IsFile](<https://github.com/gechr/x/blob/main/os/stat.go#L34>)

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

<a name="IsIOS"></a>

## func [IsIOS](<https://github.com/gechr/x/blob/main/os/platform.go#L51>)

```go
func IsIOS() bool
```

**IsIOS** reports whether the program is running on iOS. iOS is its own GOOS - distinct from macOS (see [IsDarwin](<#IsDarwin>)) - but is also Unix-like, so it additionally satisfies [IsUnix](<#IsUnix>).

<a name="IsLinux"></a>

## func [IsLinux](<https://github.com/gechr/x/blob/main/os/platform.go#L38>)

```go
func IsLinux() bool
```

**IsLinux** reports whether the program is running on Linux.

<a name="IsSymlink"></a>

## func [IsSymlink](<https://github.com/gechr/x/blob/main/os/stat.go#L46>)

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

<a name="IsUnix"></a>

## func [IsUnix](<https://github.com/gechr/x/blob/main/os/platform.go#L73>)

```go
func IsUnix() bool
```

**IsUnix** reports whether the program is running on a Unix-like OS. It mirrors Go's `unix` build constraint - which spans Linux, macOS, the BSDs, and mobile GOOSes, among others - rather than any single GOOS, so [IsLinux](<#IsLinux>), [IsDarwin](<#IsDarwin>), [IsAndroid](<#IsAndroid>), [IsIOS](<#IsIOS>), and [IsBSD](<#IsBSD>) all imply IsUnix.

<a name="IsWasm"></a>

## func [IsWasm](<https://github.com/gechr/x/blob/main/os/arch.go#L28>)

```go
func IsWasm() bool
```

**IsWasm** reports whether the program was compiled to WebAssembly. Unlike the OS predicates it checks the architecture ([runtime.GOARCH](<https://pkg.go.dev/runtime#GOARCH>)), not the OS, since WebAssembly runs under either the `js` or `wasip1` GOOS.

<a name="IsWindows"></a>

## func [IsWindows](<https://github.com/gechr/x/blob/main/os/platform.go#L27>)

```go
func IsWindows() bool
```

**IsWindows** reports whether the program is running on Windows.

<a name="IsWritableDir"></a>

## func [IsWritableDir](<https://github.com/gechr/x/blob/main/os/stat.go#L84>)

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

## func [ReadLines](<https://github.com/gechr/x/blob/main/os/lines.go#L12>)

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

<a name="RemoveIfExists"></a>

## func [RemoveIfExists](<https://github.com/gechr/x/blob/main/os/file.go#L20>)

```go
func RemoveIfExists(path string) error
```

**RemoveIfExists** removes `path`. It succeeds without error if `path` does not exist.

<details><summary><b>Example</b></summary>

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

path := filepath.Join(dir, "file.txt")
_ = os.WriteFile(path, []byte("hello"), 0o600)

fmt.Println(xos.RemoveIfExists(path))
fmt.Println(xos.RemoveIfExists(path))
```

Output:

```text
<nil>
<nil>
```

</details>

<a name="SameFile"></a>

## func [SameFile](<https://github.com/gechr/x/blob/main/os/file.go#L14>)

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

## func [Trash](<https://github.com/gechr/x/blob/main/os/trash.go#L30>)

```go
func Trash(path string) error
```

**Trash** asks the operating system to move `path` to its trash (or recycle bin) rather than removing it permanently like [os.Remove](<https://pkg.go.dev/os#Remove>), so it can typically be recovered. The `path` is resolved to an absolute path first, so a relative path trashes the intended file regardless of the working directory.

The mechanism is platform-specific: the system trash tool on macOS (so the Finder's "Put Back" works), the FreeDesktop.org trash specification on Linux and other Unix systems, and the shell file operation that targets the Recycle Bin on Windows. Recoverability is the OS's to honor, not a guarantee: an environment with the Recycle Bin disabled, for instance, may delete outright.

Where the platform cannot trash, it returns an error wrapping [errors.ErrUnsupported](<https://pkg.go.dev/errors#ErrUnsupported>), so a caller can detect the case and decide what to do (e.g. fall back to [os.Remove](<https://pkg.go.dev/os#Remove>)). This covers a macOS older than 15 (which lacks the system trash tool) and a Unix file with no usable same-device trash.

A `path` that is already gone yields an error wrapping [os.ErrNotExist](<https://pkg.go.dev/os#ErrNotExist>), both when it is missing up front and when it vanishes mid-trash (e.g. a concurrent process trashing the same file wins the race). A caller for which the file's absence is the intended end state can treat that with [errors.Is](<https://pkg.go.dev/errors#Is>) as success.

<a name="WriteLines"></a>

## func [WriteLines](<https://github.com/gechr/x/blob/main/os/lines.go#L22>)

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
