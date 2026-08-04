package palette_test

import (
	"image/color"
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
		require.Len(t, p, 32, name)

		seen := map[color.Color]bool{}
		for _, c := range p {
			require.False(t, seen[c], name, c)
			seen[c] = true
		}
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

func TestAutoWithTrueColorForcesTrueColor(t *testing.T) {
	// Detection reads COLORTERM/TERM, so pin them to a non-true-color
	// terminal; the override alone must select the true-color palette.
	t.Setenv("COLORTERM", "")
	t.Setenv("TERM", "xterm-256color")

	require.Len(t, palette.Auto(palette.WithTrueColor()), 32)
}

func TestAutoWithDarkSelectsVariant(t *testing.T) {
	dark := palette.Auto(palette.WithTrueColor(), palette.WithDark(true))
	light := palette.Auto(palette.WithTrueColor(), palette.WithDark(false))

	require.Equal(t, palette.TrueColorDark(), dark)
	require.Equal(t, palette.TrueColorLight(), light)
}

func TestAutoWithReservedAvoidsColors(t *testing.T) {
	themeRed := lipgloss.Color("#f38ba8")
	p := palette.Auto(
		palette.WithTrueColor(),
		palette.WithDark(true),
		palette.WithReserved(themeRed),
	)

	require.Equal(t, palette.TrueColorDark().Avoiding(themeRed), p)
}

func TestAutoWithReservedMeasuresAsRenderedWithoutTrueColor(t *testing.T) {
	// Pin the environment to a non-true-color terminal so Auto takes the
	// ANSI-256 path, which must measure both sides as rendered.
	t.Setenv("COLORTERM", "")
	t.Setenv("TERM", "xterm-256color")

	reserved := palette.SemanticDark().Colors()
	p := palette.Auto(palette.WithDark(true), palette.WithReserved(reserved...))

	require.Equal(t, palette.AvoidingAsANSI256(palette.DefaultDark(), reserved...), p)
}

func luminance(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	return r + g + b
}
