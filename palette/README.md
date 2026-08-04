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
  - [func WithDark(dark bool) Option](<#WithDark>)
  - [func WithReserved(reserved ...color.Color) Option](<#WithReserved>)
  - [func WithTrueColor() Option](<#WithTrueColor>)
- [type Palette](<#Palette>)
  - [func Auto(opts ...Option) Palette](<#Auto>)
  - [func DefaultDark() Palette](<#DefaultDark>)
  - [func DefaultLight() Palette](<#DefaultLight>)
  - [func TrueColorDark() Palette](<#TrueColorDark>)
  - [func TrueColorLight() Palette](<#TrueColorLight>)
  - [func (p Palette) Assigner() \*Assigner](<#Palette.Assigner>)
  - [func (p Palette) Avoiding(reserved ...color.Color) Palette](<#Palette.Avoiding>)
  - [func (p Palette) Color(text string) color.Color](<#Palette.Color>)
- [type Semantic](<#Semantic>)
  - [func SemanticDark() Semantic](<#SemanticDark>)
  - [func SemanticLight() Semantic](<#SemanticLight>)
  - [func (s Semantic) Colors() \[\]color.Color](<#Semantic.Colors>)

<a name="Assigner"></a>

## type [Assigner](<https://github.com/gechr/x/blob/main/palette/assigner.go#L15-L21>)

**Assigner** hands out colors from a [Palette](<#Palette>) so distinct keys receive distinct colors in palette order, repeating only after the palette is exhausted. Assignments are remembered, so a key always resolves to the same color for the Assigner's lifetime. Assign keys in a stable order, such as sorted order, to reproduce the same mapping across runs.

An Assigner is safe for concurrent use.

```go
type Assigner struct {
    // contains filtered or unexported fields
}
```

<a name="NewAssigner"></a>

### func [NewAssigner](<https://github.com/gechr/x/blob/main/palette/assigner.go#L29>)

```go
func NewAssigner(colors ...color.Color) *Assigner
```

**NewAssigner** returns an Assigner that draws from the given colors. When no colors are given, it defaults to [Auto](<#Auto>), selecting a palette that matches the terminal background and true-color support. Pass an explicit palette by spreading it, e.g. NewAssigner(TrueColorDark()...). Spreading an empty palette also triggers the [Auto](<#Auto>) fallback; use [Palette.Assigner](<#Palette.Assigner>) for a palette that must stay empty, such as the result of [Palette.Avoiding](<#Palette.Avoiding>).

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

### func (\*Assigner) [Assign](<https://github.com/gechr/x/blob/main/palette/assigner.go#L56>)

```go
func (a *Assigner) Assign(key string) color.Color
```

**Assign** returns the next palette color for a new key. The choice is remembered, so a key always resolves to the same color thereafter. Colors repeat in palette order after the palette is exhausted. It returns nil when the palette is empty.

<a name="Assigner.Palette"></a>

### func (\*Assigner) [Palette](<https://github.com/gechr/x/blob/main/palette/assigner.go#L48>)

```go
func (a *Assigner) Palette() Palette
```

**Palette** returns the underlying palette the Assigner draws from.

<a name="Option"></a>

## type [Option](<https://github.com/gechr/x/blob/main/palette/options.go#L13>)

**Option** configures [Auto](<#Auto>).

```go
type Option func(*config)
```

<a name="WithDark"></a>

### func [WithDark](<https://github.com/gechr/x/blob/main/palette/options.go#L28>)

```go
func WithDark(dark bool) Option
```

**WithDark** tells [Auto](<#Auto>) whether the terminal background is dark, skipping its background detection entirely. Use it when the background was already detected, for example once at startup.

<a name="WithReserved"></a>

### func [WithReserved](<https://github.com/gechr/x/blob/main/palette/options.go#L40>)

```go
func WithReserved(reserved ...color.Color) Option
```

**WithReserved** keeps the returned palette perceptually clear of the reserved colors, as if the selected palette were passed to [Palette.Avoiding](<#Palette.Avoiding>): any palette color that could be mistaken for a reserved one is removed. Use it when entity colors render alongside colors that carry meaning of their own, such as a theme's semantic red or a [Semantic](<#Semantic>) set. On the non-true-color path, colors are measured as ANSI-256 renders them.

<a name="WithTrueColor"></a>

### func [WithTrueColor](<https://github.com/gechr/x/blob/main/palette/options.go#L18>)

```go
func WithTrueColor() Option
```

**WithTrueColor** forces [Auto](<#Auto>) to return the true-color palette, overriding its automatic terminal-capability detection. The caller is responsible for ensuring the terminal supports true color.

<a name="Palette"></a>

## type [Palette](<https://github.com/gechr/x/blob/main/palette/palette.go#L23>)

**Palette** is an ordered list of colors from which a stable color is chosen for a given string.

```go
type Palette []color.Color
```

<a name="Auto"></a>

### func [Auto](<https://github.com/gechr/x/blob/main/palette/palette.go#L52>)

```go
func Auto(opts ...Option) Palette
```

**Auto** returns a palette matching the terminal, defaulting to the dark palette when background detection is unavailable. By default it selects the true-color palette when the terminal supports true color and the ANSI-256 palette otherwise; pass [WithTrueColor](<#WithTrueColor>) to force the true-color palette.

Background detection queries the terminal, waiting up to 10 milliseconds for a response on the first call; the result is cached for the process. Pass [WithDark](<#WithDark>) to skip detection entirely, for example when the background was already detected at startup.

<details><summary><b>Example</b></summary>

```go
// Auto detects the terminal background and true-color support, then colors
// entities so identical strings share a stable color.
p := palette.Auto()

for _, entity := range []string{"alpha", "beta", "alpha"} {
    _ = p.Color(entity)
}

// Force the true-color palette regardless of detection.
p = palette.Auto(palette.WithTrueColor())
fmt.Println(len(p))
```

Output:

```text
32
```

</details>

<a name="DefaultDark"></a>

### func [DefaultDark](<https://github.com/gechr/x/blob/main/palette/palette.go#L104>)

```go
func DefaultDark() Palette
```

**DefaultDark** returns the default ANSI-256 palette tuned for dark backgrounds.

<a name="DefaultLight"></a>

### func [DefaultLight](<https://github.com/gechr/x/blob/main/palette/palette.go#L141>)

```go
func DefaultLight() Palette
```

**DefaultLight** returns the default palette tuned for light backgrounds: the [DefaultDark](<#DefaultDark>) palette darkened for contrast against a light background.

<a name="TrueColorDark"></a>

### func [TrueColorDark](<https://github.com/gechr/x/blob/main/palette/truecolor.go#L7>)

```go
func TrueColorDark() Palette
```

**TrueColorDark** returns the 24-bit palette tuned for dark backgrounds. Every color clears 4.5:1 contrast against #1e1e1e, and colors are ordered so that the first N are the most separable N. Each call returns a fresh palette the caller owns.

<a name="TrueColorLight"></a>

### func [TrueColorLight](<https://github.com/gechr/x/blob/main/palette/truecolor.go#L21>)

```go
func TrueColorLight() Palette
```

**TrueColorLight** returns the 24-bit palette tuned for light backgrounds, ordered like [TrueColorDark](<#TrueColorDark>) but measured against #fafafa.

<a name="Palette.Assigner"></a>

### func (Palette) [Assigner](<https://github.com/gechr/x/blob/main/palette/assigner.go#L40>)

```go
func (p Palette) Assigner() *Assigner
```

**Assigner** returns an Assigner that draws from exactly this palette, even when it is empty — unlike [NewAssigner](<#NewAssigner>), which treats no colors as a request for [Auto](<#Auto>). An empty palette assigns nil to every key.

<a name="Palette.Avoiding"></a>

### func (Palette) [Avoiding](<https://github.com/gechr/x/blob/main/palette/avoid.go#L37>)

```go
func (p Palette) Avoiding(reserved ...color.Color) Palette
```

**Avoiding** returns a new palette without the colors that sit perceptually close to any reserved color, so entity colors cannot be mistaken for a reserved one — a theme's semantic red, say. Closeness is measured with the CIEDE2000 color-difference formula. The remaining colors keep their original order, so a most-separable-first palette stays that way.

Distances are measured on the colors' own RGBA values, so the guarantee is for opaque colors rendered in true color. Rendering in a reduced color profile quantizes colors and can narrow their distances; [Auto](<#Auto>) accounts for this on its non-true-color path by measuring colors as ANSI-256 renders them. Nil and fully transparent reserved colors are ignored; palette colors that cannot be measured are kept.

The mapping from string to color differs between the filtered and unfiltered palettes, since [Palette.Color](<#Palette.Color>) hashes over the palette length. A reserved set that blankets the palette can leave it empty, in which case [Palette.Color](<#Palette.Color>) returns nil; use [Palette.Assigner](<#Palette.Assigner>) rather than [NewAssigner](<#NewAssigner>) to keep such a palette empty.

<details><summary><b>Example</b></summary>

```go
// A theme that renders dangerous entities in its own semantic red can keep
// entity colors perceptually clear of it, so the red stays unmistakable.
themeRed := lipgloss.Color("#f38ba8")
p := palette.TrueColorDark().Avoiding(themeRed)

fmt.Println(len(p))
```

Output:

```text
27
```

</details>

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

<a name="Semantic"></a>

## type [Semantic](<https://github.com/gechr/x/blob/main/palette/semantic.go#L17-L22>)

**Semantic** holds measured colors for conventional message meanings. Like the entity palettes, every color clears 4.5:1 contrast against its target background, and the colors are mutually distinguishable. Which strings warrant a semantic color is the caller's policy; the set only guarantees the colors themselves are legible and distinct.

To keep entity colors perceptually clear of the set, reserve it: pass the [Semantic.Colors](<#Semantic.Colors>) slice to [WithReserved](<#WithReserved>) or [Palette.Avoiding](<#Palette.Avoiding>).

```go
type Semantic struct {
    Danger  color.Color
    Warning color.Color
    Success color.Color
    Info    color.Color
}
```

<a name="SemanticDark"></a>

### func [SemanticDark](<https://github.com/gechr/x/blob/main/palette/semantic.go#L26>)

```go
func SemanticDark() Semantic
```

**SemanticDark** returns the semantic set tuned for dark backgrounds, measured against #1e1e1e like [TrueColorDark](<#TrueColorDark>).

<details><summary><b>Example</b></summary>

```go
// A measured semantic set: every color clears 4.5:1 against a dark
// background, and reserving it keeps entity colors clear of the set.
sem := palette.SemanticDark()
p := palette.Auto(
    palette.WithTrueColor(),
    palette.WithDark(true),
    palette.WithReserved(sem.Colors()...),
)

fmt.Println(len(p) < len(palette.TrueColorDark()))
```

Output:

```text
true
```

</details>

<a name="SemanticLight"></a>

### func [SemanticLight](<https://github.com/gechr/x/blob/main/palette/semantic.go#L37>)

```go
func SemanticLight() Semantic
```

**SemanticLight** returns the semantic set tuned for light backgrounds, measured against #fafafa like [TrueColorLight](<#TrueColorLight>).

<a name="Semantic.Colors"></a>

### func (Semantic) [Colors](<https://github.com/gechr/x/blob/main/palette/semantic.go#L48>)

```go
func (s Semantic) Colors() []color.Color
```

**Colors** returns the set as a slice, in Danger, Warning, Success, Info order — convenient for [Palette.Avoiding](<#Palette.Avoiding>) and [WithReserved](<#WithReserved>).
