package palette_test

import (
	"image/color"
	"math"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/gechr/x/palette"
	"github.com/stretchr/testify/require"
)

// minContrastRatio is the WCAG AA bar for normal text, the same bar the
// entity palettes are held to.
const minContrastRatio = 4.5

func TestSemanticContrast(t *testing.T) {
	for name, tc := range map[string]struct {
		set        palette.Semantic
		background color.Color
	}{
		"dark":  {palette.SemanticDark(), lipgloss.Color("#1e1e1e")},
		"light": {palette.SemanticLight(), lipgloss.Color("#fafafa")},
	} {
		t.Run(name, func(t *testing.T) {
			for _, c := range tc.set.Colors() {
				require.NotNil(t, c)
				require.GreaterOrEqual(t, contrastRatio(c, tc.background), minContrastRatio)
			}
		})
	}
}

func TestSemanticMutuallyDistinguishable(t *testing.T) {
	for name, set := range map[string]palette.Semantic{
		"dark":  palette.SemanticDark(),
		"light": palette.SemanticLight(),
	} {
		t.Run(name, func(t *testing.T) {
			colors := set.Colors()
			for i, a := range colors {
				for _, b := range colors[i+1:] {
					require.GreaterOrEqual(t, distance(t, a, b), palette.MinReservedDistance)
				}
			}
		})
	}
}

func TestSemanticDisjointFromAvoidingPalette(t *testing.T) {
	// The composition the semantic set is designed for: reserving it keeps
	// every remaining entity color perceptually clear of every semantic color.
	for name, tc := range map[string]struct {
		set    palette.Semantic
		entity palette.Palette
	}{
		"dark":  {palette.SemanticDark(), palette.TrueColorDark()},
		"light": {palette.SemanticLight(), palette.TrueColorLight()},
	} {
		t.Run(name, func(t *testing.T) {
			filtered := tc.entity.Avoiding(tc.set.Colors()...)
			require.NotEmpty(t, filtered)

			for _, entity := range filtered {
				for _, semantic := range tc.set.Colors() {
					require.GreaterOrEqual(
						t, distance(t, entity, semantic), palette.MinReservedDistance,
					)
				}
			}
		})
	}
}

// contrastRatio returns the WCAG contrast ratio between two colors, from
// 1 (identical) to 21 (black on white).
func contrastRatio(a, b color.Color) float64 {
	la, lb := relativeLuminance(a), relativeLuminance(b)
	if la < lb {
		la, lb = lb, la
	}
	return (la + 0.05) / (lb + 0.05)
}

// relativeLuminance returns the WCAG relative luminance of a color.
func relativeLuminance(c color.Color) float64 {
	linear := func(channel uint32) float64 {
		v := float64(channel) / math.MaxUint16
		if v <= 0.04045 {
			return v / 12.92
		}
		return math.Pow((v+0.055)/1.055, 2.4)
	}

	r, g, b, _ := c.RGBA()
	return 0.2126*linear(r) + 0.7152*linear(g) + 0.0722*linear(b)
}
