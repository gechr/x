package palette_test

import (
	"image/color"
	"math"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/gechr/x/palette"
	"github.com/stretchr/testify/require"
)

func TestColorStable(t *testing.T) {
	p := palette.DefaultDark()

	first := p.Color("alpha")
	second := p.Color("alpha")

	require.NotNil(t, first)
	require.Equal(t, first, second)
}

func TestColorAllInputsResolve(t *testing.T) {
	p := palette.DefaultDark()
	members := map[color.Color]bool{}
	for _, c := range p {
		members[c] = true
	}

	// Hash values vary widely (including ones whose high bit is set); every
	// input must map to a color from the palette and never panic.
	for _, text := range []string{
		"alpha", "beta", "gamma", "delta", "epsilon", "zeta",
		"eta", "theta", "iota", "kappa", "lambda", "mu",
	} {
		c := p.Color(text)
		require.NotNil(t, c, text)
		require.True(t, members[c], text)
	}
}

func TestColorEmptyPalette(t *testing.T) {
	require.Nil(t, palette.Palette(nil).Color("anything"))
	require.Nil(t, palette.Palette{}.Color("anything"))
}

func TestColorSingleton(t *testing.T) {
	only := lipgloss.Color("#112233")
	p := palette.Palette{only}

	require.Equal(t, only, p.Color("anything"))
}

func TestTrueColorPalettesAreDistinct(t *testing.T) {
	for name, p := range map[string]palette.Palette{
		"dark":  palette.TrueColorDark(),
		"light": palette.TrueColorLight(),
	} {
		require.Len(t, p, 36, name)

		seen := map[color.Color]bool{}
		for _, c := range p {
			require.False(t, seen[c], name, c)
			seen[c] = true
		}
	}
}

func TestTrueColorPalettesContrastWithBackground(t *testing.T) {
	for name, tc := range map[string]struct {
		palette    palette.Palette
		background color.Color
	}{
		"dark":  {palette.TrueColorDark(), color.Black},
		"light": {palette.TrueColorLight(), color.White},
	} {
		t.Run(name, func(t *testing.T) {
			for _, foreground := range tc.palette {
				require.GreaterOrEqual(t, contrastRatio(foreground, tc.background), 4.5)
			}
		})
	}
}

func TestTrueColorAdaptsToBackground(t *testing.T) {
	dark := palette.TrueColorDark()
	light := palette.TrueColorLight()

	require.Len(t, light, len(dark))
	require.NotEqual(t, []color.Color(dark), []color.Color(light))
}

func TestDefaultLightIsDarkerThanDefaultDark(t *testing.T) {
	dark := palette.DefaultDark()
	light := palette.DefaultLight()

	require.Len(t, light, len(dark))
	require.Less(t, luminance(light[0]), luminance(dark[0]))
}

func TestConstructorsReturnIsolatedPalettes(t *testing.T) {
	// Each constructor must hand back a palette the caller owns; mutating it
	// must not leak into the next call's result.
	for name, ctor := range map[string]func() palette.Palette{
		"DefaultDark":    palette.DefaultDark,
		"DefaultLight":   palette.DefaultLight,
		"TrueColorDark":  palette.TrueColorDark,
		"TrueColorLight": palette.TrueColorLight,
	} {
		t.Run(name, func(t *testing.T) {
			first := ctor()
			require.NotEmpty(t, first)
			first[0] = lipgloss.Color("#000000")

			require.NotEqual(t, first[0], ctor()[0])
		})
	}
}

func TestAutoReturnsNonEmptyPalette(t *testing.T) {
	require.NotEmpty(t, palette.Auto())
}

func TestAutoWithTrueColorForcesGlasbey(t *testing.T) {
	// Detection can't yield true color in a test harness (no TTY, no
	// COLORTERM), so the override is what proves the wiring.
	require.Len(t, palette.Auto(palette.WithTrueColor()), 36)
}

func luminance(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	return r + g + b
}

func contrastRatio(a, b color.Color) float64 {
	lighter, darker := relativeLuminance(a), relativeLuminance(b)
	if lighter < darker {
		lighter, darker = darker, lighter
	}
	return (lighter + 0.05) / (darker + 0.05)
}

func relativeLuminance(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	channel := func(v uint32) float64 {
		srgb := float64(v) / 0xffff
		if srgb <= 0.04045 {
			return srgb / 12.92
		}
		return math.Pow((srgb+0.055)/1.055, 2.4)
	}
	return 0.2126*channel(r) + 0.7152*channel(g) + 0.0722*channel(b)
}
