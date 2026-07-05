# filepath

```go
import "github.com/gechr/x/filepath"
```

Package filepath provides path helpers: symlink resolution and containment checks.

## Index

- [func Expand(path string) string](<#Expand>)
- [func IsWithin(base string, targets ...string) bool](<#IsWithin>)
- [func Merge(paths \[\]string, opts ...MergeOption) \[\]string](<#Merge>)
- [func Resolve(path string) (string, error)](<#Resolve>)
- [func ResolveLenient(path string) (string, error)](<#ResolveLenient>)
- [type MergeOption](<#MergeOption>)
  - [func WithResolveSymlinks() MergeOption](<#WithResolveSymlinks>)

<a name="Expand"></a>

## func [Expand](<https://github.com/gechr/x/blob/main/filepath/path.go#L14>)

```go
func Expand(path string) string
```

**Expand** expands a leading ~ to the user's home directory and resolves environment variables via [os.ExpandEnv](<https://pkg.go.dev/os#ExpandEnv>). It is purely lexical: the result is not checked for existence or resolved against the filesystem (use [Resolve](<#Resolve>) or [ResolveLenient](<#ResolveLenient>) for that).

<a name="IsWithin"></a>

## func [IsWithin](<https://github.com/gechr/x/blob/main/filepath/path.go#L75>)

```go
func IsWithin(base string, targets ...string) bool
```

**IsWithin** reports whether all target paths are equal to or contained within `base`. Returns false when no `targets` are provided.

Example:

```text
IsWithin("src", "src/foo.go")             // true
IsWithin(".", "src/foo.go", "lib/bar.go") // true
IsWithin("src", "lib/foo.go")             // false
```

<a name="Merge"></a>

## func [Merge](<https://github.com/gechr/x/blob/main/filepath/path.go#L139>)

```go
func Merge(paths []string, opts ...MergeOption) []string
```

**Merge** reduces `paths` to the minimal set covering the same locations: comparing them as cleaned absolute paths, it drops any that duplicate or are nested within another, so a later walk visits each file once. Survivors keep their original form and first-seen order; a path whose absolute form cannot be computed is compared by its cleaned form.

The comparison is lexical by default; pass [WithResolveSymlinks](<#WithResolveSymlinks>) to compare resolved physical locations instead.

Example:

```text
Merge([]string{"a", "a"})     // ["a"]
Merge([]string{".", "./sub"}) // ["."]
Merge([]string{"a/b", "a"})   // ["a"]
Merge([]string{"a", "b"})     // ["a", "b"]
```

<a name="Resolve"></a>

## func [Resolve](<https://github.com/gechr/x/blob/main/filepath/path.go#L35>)

```go
func Resolve(path string) (string, error)
```

**Resolve** recursively follows every symlink along `path` and returns the fully resolved absolute path. On any error (missing component, cycle, permission) the input path is returned alongside the error so callers can choose whether to handle it or fall back.

<a name="ResolveLenient"></a>

## func [ResolveLenient](<https://github.com/gechr/x/blob/main/filepath/path.go#L51>)

```go
func ResolveLenient(path string) (string, error)
```

**ResolveLenient** returns an absolute path with symlinks resolved where possible. If `path` itself cannot be resolved, it resolves the parent directory and rejoins the original base name. If neither can be resolved, it returns the absolute path.

<a name="MergeOption"></a>

## type [MergeOption](<https://github.com/gechr/x/blob/main/filepath/path.go#L110>)

**MergeOption** configures [Merge](<#Merge>).

```go
type MergeOption func(*mergeConfig)
```

<a name="WithResolveSymlinks"></a>

### func [WithResolveSymlinks](<https://github.com/gechr/x/blob/main/filepath/path.go#L120>)

```go
func WithResolveSymlinks() MergeOption
```

**WithResolveSymlinks** makes [Merge](<#Merge>) compare paths by their resolved physical location (via [ResolveLenient](<#ResolveLenient>)) rather than lexically, so two spellings that reach the same target through a symlink are merged. It touches the filesystem; without it [Merge](<#Merge>) is pure and lexical.
