# os

```go
import "github.com/gechr/x/os"
```

Package os provides OS helpers: file probes, safe writes, copy, and line I/O.

## Index

- [func AtomicWrite\(path string, data \[\]byte, perm stdos.FileMode\) error](<#AtomicWrite>)
- [func CopyFile\(src, dst string\) error](<#CopyFile>)
- [func EnsureDir\(dir string, perm stdos.FileMode\) error](<#EnsureDir>)
- [func Exists\(path string\) \(bool, error\)](<#Exists>)
- [func IsDir\(path string\) \(bool, error\)](<#IsDir>)
- [func IsFile\(path string\) \(bool, error\)](<#IsFile>)
- [func IsSymlink\(path string\) \(bool, error\)](<#IsSymlink>)
- [func IsWritableDir\(dir string\) bool](<#IsWritableDir>)
- [func ReadLines\(path string\) \(\[\]string, error\)](<#ReadLines>)
- [func SameFile\(a, b string\) \(bool, error\)](<#SameFile>)
- [func Trash\(path string\) error](<#Trash>)
- [func WriteLines\(path string, lines \[\]string, perm stdos.FileMode\) error](<#WriteLines>)

<a name="AtomicWrite"></a>

## func [AtomicWrite](<https://github.com/gechr/x/blob/main/os/write.go#L12>)

```go
func AtomicWrite(path string, data []byte, perm stdos.FileMode) error
```

AtomicWrite writes data to path via a temp\-file\-and\-rename in the same directory. The temp file is removed on any failure.

<a name="CopyFile"></a>

## func [CopyFile](<https://github.com/gechr/x/blob/main/os/write.go#L55>)

```go
func CopyFile(src, dst string) error
```

CopyFile copies src to dst, preserving src's mode bits. dst is fsynced before close. When src and dst are the same file \(including via hard link\) CopyFile is a no\-op.

<a name="EnsureDir"></a>

## func [EnsureDir](<https://github.com/gechr/x/blob/main/os/write.go#L48>)

```go
func EnsureDir(dir string, perm stdos.FileMode) error
```

EnsureDir creates dir and any parents with the given permissions.

<a name="Exists"></a>

## func [Exists](<https://github.com/gechr/x/blob/main/os/stat.go#L25>)

```go
func Exists(path string) (bool, error)
```

Exists reports whether path exists.

<a name="IsDir"></a>

## func [IsDir](<https://github.com/gechr/x/blob/main/os/stat.go#L37>)

```go
func IsDir(path string) (bool, error)
```

IsDir reports whether path is a directory.

<a name="IsFile"></a>

## func [IsFile](<https://github.com/gechr/x/blob/main/os/stat.go#L31>)

```go
func IsFile(path string) (bool, error)
```

IsFile reports whether path is a regular file.

<a name="IsSymlink"></a>

## func [IsSymlink](<https://github.com/gechr/x/blob/main/os/stat.go#L43>)

```go
func IsSymlink(path string) (bool, error)
```

IsSymlink reports whether path is a symbolic link.

<a name="IsWritableDir"></a>

## func [IsWritableDir](<https://github.com/gechr/x/blob/main/os/stat.go#L58>)

```go
func IsWritableDir(dir string) bool
```

IsWritableDir reports whether dir exists and the current process can create files in it. Uses a probe file rather than permission\-bit inspection so that ACLs and immutable mounts are handled correctly.

<a name="ReadLines"></a>

## func [ReadLines](<https://github.com/gechr/x/blob/main/os/lines.go#L10>)

```go
func ReadLines(path string) ([]string, error)
```

ReadLines reads path and returns its non\-empty, trimmed lines.

<a name="SameFile"></a>

## func [SameFile](<https://github.com/gechr/x/blob/main/os/file.go#L13>)

```go
func SameFile(a, b string) (bool, error)
```

SameFile reports whether a and b identify the same file. Missing leaf paths are compared after resolving their parent directories, and existing files are compared with os.SameFile to detect hard links.

<a name="Trash"></a>

## func [Trash](<https://github.com/gechr/x/blob/main/os/trash.go#L24>)

```go
func Trash(path string) error
```

Trash asks the operating system to move path to its trash \(or recycle bin\) rather than removing it permanently like os.Remove, so it can typically be recovered. The path is resolved to an absolute path first, so a relative path trashes the intended file regardless of the working directory.

The mechanism is platform\-specific: the system trash tool on macOS \(so the Finder's "Put Back" works\), the FreeDesktop.org trash specification on Linux and other Unix systems, and the shell file operation that targets the Recycle Bin on Windows. Recoverability is the OS's to honor, not a guarantee: an environment with the Recycle Bin disabled, for instance, may delete outright.

Where the platform cannot trash, it returns an error wrapping [errors.ErrUnsupported](<https://pkg.go.dev/errors/#ErrUnsupported>), so a caller can detect the case and decide what to do \(e.g. fall back to os.Remove\). This covers a macOS older than 15 \(which lacks the system trash tool\) and a Unix file with no usable same\-device trash.

<a name="WriteLines"></a>

## func [WriteLines](<https://github.com/gechr/x/blob/main/os/lines.go#L27>)

```go
func WriteLines(path string, lines []string, perm stdos.FileMode) error
```

WriteLines atomically writes lines to path, one per line, with a trailing newline.
