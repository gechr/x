package palette

import (
	"image/color"
	"sync"
)

// Assigner hands out colors from a [Palette] so distinct keys receive distinct
// colors in palette order, repeating only after the palette is exhausted.
// Assignments are remembered, so a key always resolves to the same color for
// the Assigner's lifetime. Assign keys in a stable order, such as sorted order,
// to reproduce the same mapping across runs.
//
// An Assigner is safe for concurrent use.
type Assigner struct {
	palette Palette

	mu   sync.Mutex
	seen map[string]int
	next int
}

// NewAssigner returns an Assigner that draws from the given colors. When no
// colors are given, it defaults to [Auto], selecting a palette that matches the
// terminal background and true-color support. Pass an explicit palette by
// spreading it, e.g. NewAssigner(TrueColorDark()...). Spreading an empty
// palette also triggers the [Auto] fallback; use [Palette.Assigner] for a
// palette that must stay empty, such as the result of [Palette.Avoiding].
func NewAssigner(colors ...color.Color) *Assigner {
	pal := Palette(colors)
	if len(pal) == 0 {
		pal = Auto()
	}
	return pal.Assigner()
}

// Assigner returns an Assigner that draws from exactly this palette, even
// when it is empty — unlike [NewAssigner], which treats no colors as a
// request for [Auto]. An empty palette assigns nil to every key.
func (p Palette) Assigner() *Assigner {
	return &Assigner{
		palette: p,
		seen:    make(map[string]int),
	}
}

// Palette returns the underlying palette the Assigner draws from.
func (a *Assigner) Palette() Palette {
	return a.palette
}

// Assign returns the next palette color for a new key. The choice is remembered,
// so a key always resolves to the same color thereafter. Colors repeat in
// palette order after the palette is exhausted. It returns nil when the palette
// is empty.
func (a *Assigner) Assign(key string) color.Color {
	n := len(a.palette)
	if n == 0 {
		return nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if i, ok := a.seen[key]; ok {
		return a.palette[i]
	}

	i := a.next % n
	a.next++
	a.seen[key] = i
	return a.palette[i]
}
