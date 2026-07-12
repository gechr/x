# encoding

```go
import "github.com/gechr/x/encoding"
```

Package `encoding` provides helpers for structured document formats such as JSON and YAML: building field paths for diagnostics and lookup.

## Index

- [type Path](<#Path>)
  - [func NewPath(name string, moreNames ...string) \*Path](<#NewPath>)
  - [func (p \*Path) Child(name string, moreNames ...string) \*Path](<#Path.Child>)
  - [func (p \*Path) Index(index int) \*Path](<#Path.Index>)
  - [func (p \*Path) Key(key string) \*Path](<#Path.Key>)
  - [func (p \*Path) Lookup(doc any) (any, bool)](<#Path.Lookup>)
  - [func (p \*Path) LookupAll(doc any) \[\]any](<#Path.LookupAll>)
  - [func (p \*Path) Render(opts ...RenderOption) string](<#Path.Render>)
  - [func (p \*Path) String() string](<#Path.String>)
  - [func (p \*Path) Wildcard() \*Path](<#Path.Wildcard>)
- [type RenderOption](<#RenderOption>)
  - [func WithRoot(marker ...rune) RenderOption](<#WithRoot>)

<a name="Path"></a>

## type [Path](<https://github.com/gechr/x/blob/main/encoding/path.go#L32-L37>)

**Path** is an immutable field path into a structured document (JSON, YAML, ...), built segment by segment for diagnostics and lookup:

```go
NewPath("items").Index(0).Child("foo", "bar").Wildcard() // items[0].foo.bar[*]
```

Each method allocates a new Path referencing its receiver as parent, so multiple children can safely branch off a shared prefix - handy in recursive document walkers.

A nil \*Path is the empty path: it renders as "" and addresses the document itself in lookups. All methods are nil-safe.

```go
type Path struct {
    // contains filtered or unexported fields
}
```

<a name="NewPath"></a>

### func [NewPath](<https://github.com/gechr/x/blob/main/encoding/path.go#L41>)

```go
func NewPath(name string, moreNames ...string) *Path
```

**NewPath** returns a [Path](<#Path>) rooted at `name`, extended with `moreNames` as nested children.

<details><summary><b>Example</b></summary>

```go
p := xencoding.NewPath("items").Index(0).Child("foo", "bar").Wildcard()
fmt.Println(p)
```

Output:

```text
items[0].foo.bar[*]
```

</details>

<a name="Path.Child"></a>

### func (\*Path) [Child](<https://github.com/gechr/x/blob/main/encoding/path.go#L50>)

```go
func (p *Path) Child(name string, moreNames ...string) *Path
```

**Child** returns `p` extended with `name` (and `moreNames`) as nested field segments. Names render in dot notation (".name"), or bracket-quoted (`["a.b"]`) when they contain characters other than letters, digits, '\_', and '-'.

<details><summary><b>Example</b></summary>

Names that cannot appear in dot notation are bracket-quoted automatically.

```go
p := xencoding.NewPath("metadata", "labels").Child("kubernetes.io/hostname")
fmt.Println(p)
```

Output:

```text
metadata.labels["kubernetes.io/hostname"]
```

</details>

<a name="Path.Index"></a>

### func (\*Path) [Index](<https://github.com/gechr/x/blob/main/encoding/path.go#L59>)

```go
func (p *Path) Index(index int) *Path
```

**Index** returns `p` extended with an array index segment, rendered as "\[3\]".

<a name="Path.Key"></a>

### func (\*Path) [Key](<https://github.com/gechr/x/blob/main/encoding/path.go#L66>)

```go
func (p *Path) Key(key string) *Path
```

**Key** returns `p` extended with an explicit key segment, always rendered bracket-quoted (`["name"]`) even when `key` would be valid in dot notation. Lookup treats it exactly like [Path.Child](<#Path.Child>).

<a name="Path.Lookup"></a>

### func (\*Path) [Lookup](<https://github.com/gechr/x/blob/main/encoding/lookup.go#L18>)

```go
func (p *Path) Lookup(doc any) (any, bool)
```

**Lookup** resolves `p` against `doc`, a document decoded into generic Go values (map\[string\]any, map\[any\]any, \[\]any), and returns the value it addresses. It reports false when a segment is missing, an index is out of range, or the path contains a wildcard segment (use [Path.LookupAll](<#Path.LookupAll>)).

<details><summary><b>Example</b></summary>

```go
doc := map[string]any{"spec": map[string]any{"replicas": 3}}
v, ok := xencoding.NewPath("spec", "replicas").Lookup(doc)
fmt.Println(v, ok)
```

Output:

```text
3 true
```

</details>

<a name="Path.LookupAll"></a>

### func (\*Path) [LookupAll](<https://github.com/gechr/x/blob/main/encoding/lookup.go#L37>)

```go
func (p *Path) LookupAll(doc any) []any
```

**LookupAll** resolves `p` against `doc`, fanning out at wildcard segments: each [Path.Wildcard](<#Path.Wildcard>) matches every element of an array in order, or every value of a map in natural key order (see [strings.CompareNatural](<../strings/README.md#CompareNatural>), with ties broken lexically and then by the key's Go type). Values missing a later segment are skipped. Without wildcards the result has at most one value. Returns nil when nothing matches.

<details><summary><b>Example</b></summary>

```go
doc := map[string]any{"items": []any{
    map[string]any{"name": "a"},
    map[string]any{"name": "b"},
}}
names := xencoding.NewPath("items").Wildcard().Child("name").LookupAll(doc)
fmt.Println(names)
```

Output:

```text
[a b]
```

</details>

<a name="Path.Render"></a>

### func (\*Path) [Render](<https://github.com/gechr/x/blob/main/encoding/path.go#L82>)

```go
func (p *Path) Render(opts ...RenderOption) string
```

**Render** returns the path in dot/bracket notation, e.g. "items\[0\].foo.bar\[\*\]". Names that cannot appear in dot notation are bracket-quoted: `spec["a.b"]`. Pass [WithRoot](<#WithRoot>) to prefix a JSONPath-style root marker: "$.items\[0\]". The output is a human-readable diagnostic notation, not strict RFC 9535 JSONPath (bare names are broader, and quoting follows Go syntax).

<details><summary><b>Example</b></summary>

```go
p := xencoding.NewPath("items").Index(0)
fmt.Println(p.Render())
fmt.Println(p.Render(xencoding.WithRoot()))
fmt.Println(p.Render(xencoding.WithRoot('@')))
```

Output:

```text
items[0]
$.items[0]
@.items[0]
```

</details>

<a name="Path.String"></a>

### func (\*Path) [String](<https://github.com/gechr/x/blob/main/encoding/path.go#L98>)

```go
func (p *Path) String() string
```

**String** implements [fmt.Stringer](<https://pkg.go.dev/fmt#Stringer>); it is [Path.Render](<#Path.Render>) with default options.

<a name="Path.Wildcard"></a>

### func (\*Path) [Wildcard](<https://github.com/gechr/x/blob/main/encoding/path.go#L73>)

```go
func (p *Path) Wildcard() *Path
```

**Wildcard** returns `p` extended with a "\[\*\]" segment matching every element of an array or every value of a map. [Path.LookupAll](<#Path.LookupAll>) fans out at wildcard segments; [Path.Lookup](<#Path.Lookup>) cannot resolve them.

<a name="RenderOption"></a>

## type [RenderOption](<https://github.com/gechr/x/blob/main/encoding/options.go#L4>)

**RenderOption** configures [Path.Render](<#Path.Render>).

```go
type RenderOption func(*renderConfig)
```

<a name="WithRoot"></a>

### func [WithRoot](<https://github.com/gechr/x/blob/main/encoding/options.go#L13>)

```go
func WithRoot(marker ...rune) RenderOption
```

**WithRoot** prefixes the rendered path with a root marker: '$' by default ("$.items\[0\]"), or `marker` if given (WithRoot('@') renders "@.items\[0\]").
