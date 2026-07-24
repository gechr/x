package palette

import (
	"image/color"
	"sync"
)

// Assigner hands out colors from a [Palette] so distinct keys receive distinct
// colors. Each key is placed at its hash-derived color - the same slot
// [Palette.Color] would choose - and if that color is already taken, the
// Assigner probes forward to the next free color. Assignments are remembered,
// so a key always resolves to the same color for the Assigner's lifetime.
//
// This blends the two extremes: like [Palette.Color] a key tends to land on its
// stable hash color across runs, but unlike it distinct keys never share a color
// until the palette is exhausted (after which colors repeat). A key only moves
// off its hash color when an earlier-seen key collided into that slot, so with
// few keys relative to the palette size, most keys stay stable across runs.
//
// An Assigner is safe for concurrent use.
type Assigner struct {
	palette Palette

	mu   sync.Mutex
	seen map[string]int
	used map[int]bool
}

// NewAssigner returns an Assigner that draws from the given colors. When no
// colors are given, it defaults to [Auto], selecting a palette that matches the
// terminal background and true-color support. Pass an explicit palette by
// spreading it, e.g. NewAssigner(TrueColorDark()...).
func NewAssigner(colors ...color.Color) *Assigner {
	pal := Palette(colors)
	if len(pal) == 0 {
		pal = Auto()
	}
	return &Assigner{
		palette: pal,
		seen:    make(map[string]int),
		used:    make(map[int]bool),
	}
}

// Palette returns the underlying palette the Assigner draws from.
func (a *Assigner) Palette() Palette {
	return a.palette
}

// Assign returns the color for key: its hash-derived color when free, otherwise
// the next free color. The choice is remembered, so a key always resolves to the
// same color thereafter. It returns nil when the palette is empty.
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

	// Start at the key's stable hash slot, then linear-probe to the next free
	// color while any remain, so distinct keys stay distinct until the palette
	// is exhausted.
	i := index(key, n)
	for len(a.used) < n && a.used[i] {
		i = (i + 1) % n
	}

	a.seen[key] = i
	a.used[i] = true
	return a.palette[i]
}
