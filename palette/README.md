# palette

```go
import "github.com/gechr/x/palette"
```

Package `palette` provides stable, hash-based color selection from ordered color palettes, plus curated palettes tuned for light and dark terminal backgrounds.

The mapping from string to color is deterministic: the same text always resolves to the same color for a given palette, across processes and platforms. This makes it suitable for colorizing entities such as environments, hostnames, or identifiers so they stay visually consistent.

## Index

- [type Assigner](<#Assigner>)
  - [func NewAssigner(colors ...color.Color) \*Assigner](<#NewAssigner>)
  - [func (a \*Assigner) Assign(key string) color.Color](<#Assigner.Assign>)
  - [func (a \*Assigner) Palette() Palette](<#Assigner.Palette>)
- [type Option](<#Option>)
  - [func WithTrueColor() Option](<#WithTrueColor>)
- [type Palette](<#Palette>)
  - [func Auto(opts ...Option) Palette](<#Auto>)
  - [func DefaultDark() Palette](<#DefaultDark>)
  - [func DefaultLight() Palette](<#DefaultLight>)
  - [func TrueColorDark() Palette](<#TrueColorDark>)
  - [func TrueColorLight() Palette](<#TrueColorLight>)
  - [func (p Palette) Color(text string) color.Color](<#Palette.Color>)

<a name="Assigner"></a>

## type [Assigner](<https://github.com/gechr/x/blob/main/palette/assigner.go#L21-L27>)

**Assigner** hands out colors from a [Palette](<#Palette>) so distinct keys receive distinct colors. Each key is placed at its hash-derived color - the same slot [Palette.Color](<#Palette.Color>) would choose - and if that color is already taken, the Assigner probes forward to the next free color. Assignments are remembered, so a key always resolves to the same color for the Assigner's lifetime.

This blends the two extremes: like [Palette.Color](<#Palette.Color>) a key tends to land on its stable hash color across runs, but unlike it distinct keys never share a color until the palette is exhausted (after which colors repeat). A key only moves off its hash color when an earlier-seen key collided into that slot, so with few keys relative to the palette size, most keys stay stable across runs.

An Assigner is safe for concurrent use.

```go
type Assigner struct {
    // contains filtered or unexported fields
}
```

<a name="NewAssigner"></a>

### func [NewAssigner](<https://github.com/gechr/x/blob/main/palette/assigner.go#L33>)

```go
func NewAssigner(colors ...color.Color) *Assigner
```

**NewAssigner** returns an Assigner that draws from the given colors. When no colors are given, it defaults to [Auto](<#Auto>), selecting a palette that matches the terminal background and true-color support. Pass an explicit palette by spreading it, e.g. NewAssigner(TrueColorDark()...).

<details><summary><b>Example</b></summary>

```go
// An Assigner gives distinct keys distinct colors (no duplicates until the
// palette is exhausted) while keeping each key stable. With no colors it
// defaults to Auto for the current terminal.
a := palette.NewAssigner()

fmt.Println(a.Assign("web") == a.Assign("web"))
```

Output:

```text
true
```

</details>

<a name="Assigner.Assign"></a>

### func (\*Assigner) [Assign](<https://github.com/gechr/x/blob/main/palette/assigner.go#L53>)

```go
func (a *Assigner) Assign(key string) color.Color
```

**Assign** returns the color for key: its hash-derived color when free, otherwise the next free color. The choice is remembered, so a key always resolves to the same color thereafter. It returns nil when the palette is empty.

<a name="Assigner.Palette"></a>

### func (\*Assigner) [Palette](<https://github.com/gechr/x/blob/main/palette/assigner.go#L46>)

```go
func (a *Assigner) Palette() Palette
```

**Palette** returns the underlying palette the Assigner draws from.

<a name="Option"></a>

## type [Option](<https://github.com/gechr/x/blob/main/palette/options.go#L9>)

**Option** configures [Auto](<#Auto>).

```go
type Option func(*config)
```

<a name="WithTrueColor"></a>

### func [WithTrueColor](<https://github.com/gechr/x/blob/main/palette/options.go#L14>)

```go
func WithTrueColor() Option
```

**WithTrueColor** forces [Auto](<#Auto>) to return the 256-color Glasbey palette, overriding its automatic terminal-capability detection. The caller is responsible for ensuring the terminal supports true color.

<a name="Palette"></a>

## type [Palette](<https://github.com/gechr/x/blob/main/palette/palette.go#L23>)

**Palette** is an ordered list of colors from which a stable color is chosen for a given string.

```go
type Palette []color.Color
```

<a name="Auto"></a>

### func [Auto](<https://github.com/gechr/x/blob/main/palette/palette.go#L47>)

```go
func Auto(opts ...Option) Palette
```

**Auto** returns a palette matching the terminal, defaulting to the dark palette when background detection is unavailable. By default it selects the Glasbey palette when the terminal supports true color and the ANSI-256 palette otherwise; pass [WithTrueColor](<#WithTrueColor>) to force the Glasbey palette.

<details><summary><b>Example</b></summary>

```go
// Auto detects the terminal background and true-color support, then colors
// entities so identical strings share a stable color.
p := palette.Auto()

for _, entity := range []string{"alpha", "beta", "alpha"} {
    _ = p.Color(entity)
}

// Force the 256-color Glasbey palette regardless of detection.
p = palette.Auto(palette.WithTrueColor())
fmt.Println(len(p))
```

Output:

```text
256
```

</details>

<a name="DefaultDark"></a>

### func [DefaultDark](<https://github.com/gechr/x/blob/main/palette/palette.go#L82>)

```go
func DefaultDark() Palette
```

**DefaultDark** returns the default ANSI-256 palette tuned for dark backgrounds.

<a name="DefaultLight"></a>

### func [DefaultLight](<https://github.com/gechr/x/blob/main/palette/palette.go#L119>)

```go
func DefaultLight() Palette
```

**DefaultLight** returns the default palette tuned for light backgrounds: the [DefaultDark](<#DefaultDark>) palette darkened for contrast against a light background.

<a name="TrueColorDark"></a>

### func [TrueColorDark](<https://github.com/gechr/x/blob/main/palette/glasbey.go#L11>)

```go
func TrueColorDark() Palette
```

**TrueColorDark** returns the 256-color Glasbey palette curated for dark backgrounds - colorcet's glasbey\_light ([https://github.com/holoviz/colorcet](<https://github.com/holoviz/colorcet>)), generated once by perceptual max-min optimization so every prefix stays visually distinct. Use it when entities outnumber the default ANSI-256 palette; the caller is responsible for ensuring the terminal supports true color. Each call returns a fresh palette the caller owns.

<a name="TrueColorLight"></a>

### func [TrueColorLight](<https://github.com/gechr/x/blob/main/palette/glasbey.go#L72>)

```go
func TrueColorLight() Palette
```

**TrueColorLight** returns the 256-color Glasbey palette curated for light backgrounds - colorcet's glasbey\_dark. See [TrueColorDark](<#TrueColorDark>).

<a name="Palette.Color"></a>

### func (Palette) [Color](<https://github.com/gechr/x/blob/main/palette/palette.go#L27>)

```go
func (p Palette) Color(text string) color.Color
```

**Color** returns a stable color for text, or nil when the palette is empty. The same text always yields the same color for a given palette.

<details><summary><b>Example</b></summary>

```go
p := palette.DefaultDark()

// The same entity always resolves to the same color.
fmt.Println(p.Color("alpha") == p.Color("alpha"))
```

Output:

```text
true
```

</details>

<details><summary><b>Example (Empty)</b></summary>

```go
// An empty palette resolves every input to nil, leaving text uncolored.
fmt.Println(palette.Palette(nil).Color("alpha"))
```

Output:

```text
<nil>
```

</details>
