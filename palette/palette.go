// Package palette provides stable, hash-based color selection from ordered
// color palettes, plus curated palettes tuned for light and dark terminal
// backgrounds.
//
// The mapping from string to color is deterministic: the same text always
// resolves to the same color for a given palette, across processes and
// platforms. This makes it suitable for colorizing entities such as
// environments, hostnames, or identifiers so they stay visually consistent.
package palette

import (
	"hash/fnv"
	"image/color"
	"math"

	"charm.land/lipgloss/v2"
	xslices "github.com/gechr/x/slices"
	"github.com/gechr/x/terminal"
)

// Palette is an ordered list of colors from which a stable color is chosen
// for a given string.
type Palette []color.Color

// Color returns a stable color for text, or nil when the palette is empty.
// The same text always yields the same color for a given palette.
func (p Palette) Color(text string) color.Color {
	if len(p) == 0 {
		return nil
	}
	return p[index(text, len(p))]
}

// index returns a stable index in [0, n) for text, deterministic across
// processes and platforms. It is shared by [Palette.Color] and [Assigner] so
// their hashing cannot drift apart.
func index(text string, n int) int {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(text))
	return int(hasher.Sum32()&math.MaxInt32) % n
}

// Auto returns a palette matching the terminal, defaulting to the dark
// palette when background detection is unavailable. By default it selects the
// true-color palette when the terminal supports true color and the ANSI-256
// palette otherwise; pass [WithTrueColor] to force the true-color palette.
//
// Background detection queries the terminal, waiting up to 10 milliseconds
// for a response on the first call; the result is cached for the process.
// Pass [WithDark] to skip detection entirely, for example when the
// background was already detected at startup.
func Auto(opts ...Option) Palette {
	var c config
	for _, opt := range opts {
		opt(&c)
	}

	trueColor := terminal.SupportsTrueColor()
	if c.trueColor != nil {
		trueColor = *c.trueColor
	}

	dark := c.dark != nil && *c.dark
	if c.dark == nil {
		dark = isDark()
	}

	var p Palette
	switch {
	case trueColor && dark:
		p = TrueColorDark()
	case trueColor:
		p = TrueColorLight()
	case dark:
		p = DefaultDark()
	default:
		p = DefaultLight()
	}

	if len(c.reserved) == 0 {
		return p
	}
	if trueColor {
		return p.Avoiding(c.reserved...)
	}
	// Without true color, both palette and reserved colors render quantized
	// to ANSI-256, which can pull them closer together than their true-color
	// values suggest — so measure them as they will render.
	return p.avoiding(ansi256, c.reserved...)
}

// isDark reports whether the terminal has a dark background, defaulting to
// true when detection is unavailable (the conventional dark fallback).
func isDark() bool {
	dark, ok := terminal.IsDark()
	return dark || !ok
}

// darkenPercent is how much [DefaultDark] is darkened to suit a light
// background in [DefaultLight].
const darkenPercent = 0.45

// DefaultDark returns the default ANSI-256 palette tuned for dark backgrounds.
func DefaultDark() Palette {
	return Palette{
		lipgloss.Color("208"), // orange
		lipgloss.Color("51"),  // cyan
		lipgloss.Color("226"), // yellow
		lipgloss.Color("207"), // magenta
		lipgloss.Color("82"),  // green
		lipgloss.Color("75"),  // blue
		lipgloss.Color("214"), // orange (light)
		lipgloss.Color("177"), // purple
		lipgloss.Color("48"),  // spring green
		lipgloss.Color("87"),  // turquoise
		lipgloss.Color("220"), // gold
		lipgloss.Color("141"), // purple (light)
		lipgloss.Color("118"), // green (light)
		lipgloss.Color("50"),  // spring green (light)
		lipgloss.Color("213"), // pink
		lipgloss.Color("111"), // sky blue
		lipgloss.Color("156"), // pale green
		lipgloss.Color("183"), // plum
		lipgloss.Color("229"), // pale yellow
		lipgloss.Color("123"), // pale cyan
		lipgloss.Color("203"), // red
		lipgloss.Color("63"),  // blue
		lipgloss.Color("173"), // brown
		lipgloss.Color("250"), // grey
		lipgloss.Color("37"),  // teal
		lipgloss.Color("57"),  // indigo
		lipgloss.Color("124"), // maroon
		lipgloss.Color("100"), // olive
		lipgloss.Color("209"), // coral
		lipgloss.Color("103"), // slate
	}
}

// DefaultLight returns the default palette tuned for light backgrounds: the
// [DefaultDark] palette darkened for contrast against a light background.
func DefaultLight() Palette {
	return xslices.Map(DefaultDark(), func(c color.Color) color.Color {
		return lipgloss.Darken(c, darkenPercent)
	})
}

// hexColors converts hex strings to a palette.
func hexColors(hexes ...string) Palette {
	return xslices.Map(hexes, lipgloss.Color)
}
