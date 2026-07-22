# filepath

```go
import "github.com/gechr/x/filepath"
```

Package `filepath` provides path helpers: symlink resolution and containment checks.

## Index

- [func Expand(path string) string](<#Expand>)
- [func IsWithin(base string, targets ...string) bool](<#IsWithin>)
- [func Merge(paths \[\]string, opts ...MergeOption) \[\]string](<#Merge>)
- [func Resolve(path string) (string, error)](<#Resolve>)
- [func ResolveLenient(path string) (string, error)](<#ResolveLenient>)
- [func SplitPath(list string) \[\]string](<#SplitPath>)
- [type MergeOption](<#MergeOption>)
  - [func WithResolveSymlinks() MergeOption](<#WithResolveSymlinks>)

<a name="Expand"></a>

## func [Expand](<https://github.com/gechr/x/blob/main/filepath/path.go#L14>)

```go
func Expand(path string) string
```

**Expand** expands a leading ~ to the user's home directory and resolves environment variables via [os.ExpandEnv](<https://pkg.go.dev/os#ExpandEnv>). It is purely lexical: the result is not checked for existence or resolved against the filesystem (use [Resolve](<#Resolve>) or [ResolveLenient](<#ResolveLenient>) for that).

<details><summary><b>Example</b></summary>

**Expand** also expands a leading ~ to the user's home directory.

```go
_ = os.Setenv("PROJECT_ROOT", "/srv/app")
fmt.Println(xfilepath.Expand("$PROJECT_ROOT/config.toml"))
```

Output:

```text
/srv/app/config.toml
```

</details>

<a name="IsWithin"></a>

## func [IsWithin](<https://github.com/gechr/x/blob/main/filepath/path.go#L96>)

```go
func IsWithin(base string, targets ...string) bool
```

**IsWithin** reports whether all target paths are equal to or contained within `base`. Returns false when no `targets` are provided.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xfilepath.IsWithin("src", "src/foo.go"))
fmt.Println(xfilepath.IsWithin("src", "src"))
fmt.Println(xfilepath.IsWithin("src", "lib/foo.go"))
fmt.Println(xfilepath.IsWithin("src"))
```

Output:

```text
true
true
false
false
```

</details>

<details><summary><b>Example (MultipleTargets)</b></summary>

**IsWithin** only reports `true` when every target is contained within the base.

```go
fmt.Println(xfilepath.IsWithin(".", "src/foo.go", "lib/bar.go"))
fmt.Println(xfilepath.IsWithin("src", "src/foo.go", "lib/bar.go"))
```

Output:

```text
true
false
```

</details>

<a name="Merge"></a>

## func [Merge](<https://github.com/gechr/x/blob/main/filepath/path.go#L160>)

```go
func Merge(paths []string, opts ...MergeOption) []string
```

**Merge** reduces `paths` to the minimal set covering the same locations: comparing them as cleaned absolute paths, it drops any that duplicate or are nested within another, so a later walk visits each file once. Survivors keep their original form and first-seen order; a path whose absolute form cannot be computed is compared by its cleaned form.

The comparison is lexical by default; pass [WithResolveSymlinks](<#WithResolveSymlinks>) to compare resolved physical locations instead.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xfilepath.Merge([]string{".", "./sub"}))
fmt.Println(xfilepath.Merge([]string{"a/b", "a"}))
fmt.Println(xfilepath.Merge([]string{"a", "b"}))
```

Output:

```text
[.]
[a]
[a b]
```

</details>

<details><summary><b>Example (Duplicates)</b></summary>

Exact duplicates are merged; the first occurrence survives in its original spelling.

```go
fmt.Println(xfilepath.Merge([]string{"a", "./a", "a/"}))
```

Output:

```text
[a]
```

</details>

<a name="Resolve"></a>

## func [Resolve](<https://github.com/gechr/x/blob/main/filepath/path.go#L56>)

```go
func Resolve(path string) (string, error)
```

**Resolve** recursively follows every symlink along `path` and returns the fully resolved absolute path. On any error (missing component, cycle, permission) the input path is returned alongside the error so callers can choose whether to handle it or fall back.

<details><summary><b>Example</b></summary>

**Resolve** follows symlinks, so a link and its target resolve to the same path.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

target := filepath.Join(dir, "target.txt")
_ = os.WriteFile(target, []byte("hello"), 0o600)
link := filepath.Join(dir, "link.txt")
_ = os.Symlink(target, link)

resolvedLink, _ := xfilepath.Resolve(link)
resolvedTarget, _ := xfilepath.Resolve(target)
fmt.Println(resolvedLink == resolvedTarget)
```

Output:

```text
true
```

</details>

<a name="ResolveLenient"></a>

## func [ResolveLenient](<https://github.com/gechr/x/blob/main/filepath/path.go#L72>)

```go
func ResolveLenient(path string) (string, error)
```

**ResolveLenient** returns an absolute path with symlinks resolved where possible. If `path` itself cannot be resolved, it resolves the parent directory and rejoins the original base name. If neither can be resolved, it returns the absolute path.

<details><summary><b>Example</b></summary>

**ResolveLenient** succeeds where Resolve fails: a missing file resolves via its parent directory, keeping the original base name.

```go
dir, _ := os.MkdirTemp("", "example")
defer func() { _ = os.RemoveAll(dir) }()

missing := filepath.Join(dir, "missing.txt")
_, err := xfilepath.Resolve(missing)
fmt.Println(err != nil)

resolved, _ := xfilepath.ResolveLenient(missing)
fmt.Println(filepath.Base(resolved))
```

Output:

```text
true
missing.txt
```

</details>

<a name="SplitPath"></a>

## func [SplitPath](<https://github.com/gechr/x/blob/main/filepath/path.go#L41>)

```go
func SplitPath(list string) []string
```

**SplitPath** splits a PATH-style list (such as $PATH or $GOPATH) on the OS-specific list separator ([os.PathListSeparator](<https://pkg.go.dev/os#PathListSeparator>)), dropping the empty entries produced by leading, trailing, or doubled separators - an empty entry otherwise resolves to the current directory when joined. On Windows it honours the same quoting rules as [filepath.SplitList](<https://pkg.go.dev/path/filepath#SplitList>).

<details><summary><b>Example</b></summary>

**SplitPath** splits a PATH-style list and drops the empty entry left by the trailing separator.

```go
sep := string(os.PathListSeparator)
fmt.Println(xfilepath.SplitPath("/usr/bin" + sep + "/bin" + sep))
```

Output:

```text
[/usr/bin /bin]
```

</details>

<a name="MergeOption"></a>

## type [MergeOption](<https://github.com/gechr/x/blob/main/filepath/path.go#L131>)

**MergeOption** configures [Merge](<#Merge>).

```go
type MergeOption func(*mergeConfig)
```

<a name="WithResolveSymlinks"></a>

### func [WithResolveSymlinks](<https://github.com/gechr/x/blob/main/filepath/path.go#L141>)

```go
func WithResolveSymlinks() MergeOption
```

**WithResolveSymlinks** makes [Merge](<#Merge>) compare paths by their resolved physical location (via [ResolveLenient](<#ResolveLenient>)) rather than lexically, so two spellings that reach the same target through a symlink are merged. It touches the filesystem; without it [Merge](<#Merge>) is pure and lexical.
