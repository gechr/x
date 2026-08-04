package palette_test

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/gechr/x/palette"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/require"
)

func TestAvoidingRemovesLookalikes(t *testing.T) {
	// The motivating case: a theme's semantic red rendered beside entity
	// colors. Every surviving entity color must be unmistakable for it.
	themeRed := lipgloss.Color("#f38ba8")
	p := palette.TrueColorDark().Avoiding(themeRed)

	require.NotEmpty(t, p)
	require.Less(t, len(p), len(palette.TrueColorDark()))

	for _, c := range p {
		require.GreaterOrEqual(t, distance(t, themeRed, c), palette.MinReservedDistance)
	}
	for _, lookalike := range []string{"#ffacc8", "#f368d3", "#e6577d", "#f98a7e"} {
		require.NotContains(t, p, lipgloss.Color(lookalike))
	}
}

func TestAvoidingPreservesOrder(t *testing.T) {
	original := palette.TrueColorDark()
	filtered := original.Avoiding(lipgloss.Color("#f38ba8"))

	// The survivors must appear in their original most-separable-first order.
	members := map[color.Color]bool{}
	for _, c := range filtered {
		members[c] = true
	}

	var expected palette.Palette
	for _, c := range original {
		if members[c] {
			expected = append(expected, c)
		}
	}
	require.Equal(t, expected, filtered)
}

func TestAvoidingNoReservedReturnsIsolatedCopy(t *testing.T) {
	original := palette.TrueColorDark()
	filtered := original.Avoiding()

	require.Equal(t, original, filtered)

	filtered[0] = lipgloss.Color("#000000")
	require.Equal(t, palette.TrueColorDark()[0], original[0])
}

func TestAvoidingEverythingYieldsEmpty(t *testing.T) {
	p := palette.TrueColorDark()

	require.Empty(t, p.Avoiding(p...))
	require.Nil(t, p.Avoiding(p...).Color("anything"))
}

func TestAvoidingKeepsUnmeasurableColors(t *testing.T) {
	transparent := color.RGBA{}
	p := palette.Palette{transparent, nil}.Avoiding(lipgloss.Color("#f38ba8"))

	require.Equal(t, palette.Palette{transparent, nil}, p)
}

func TestAvoidingIgnoresUnmeasurableReserved(t *testing.T) {
	original := palette.TrueColorDark()

	require.Equal(t, original, original.Avoiding(color.RGBA{}))
	require.Equal(t, original, original.Avoiding(nil))
	require.Equal(t, original, original.Avoiding(palette.Semantic{}.Colors()...))
}

func TestAvoidingAsANSI256MeasuresAsRendered(t *testing.T) {
	// Rendering without true color quantizes both sides to ANSI-256, which
	// can pull colors below the exclusion distance even though their
	// true-color values clear it. The render-aware filter must exclude those;
	// the plain filter provably does not (proving the ANSI path is load-
	// bearing, not redundant).
	for name, tc := range map[string]struct {
		entity   palette.Palette
		reserved []color.Color
	}{
		"dark":  {palette.DefaultDark(), palette.SemanticDark().Colors()},
		"light": {palette.DefaultLight(), palette.SemanticLight().Colors()},
	} {
		t.Run(name, func(t *testing.T) {
			violations := func(p palette.Palette) int {
				count := 0
				for _, e := range p {
					for _, r := range tc.reserved {
						rendered := distance(t, palette.ANSI256(e), palette.ANSI256(r))
						if rendered < palette.MinReservedDistance {
							count++
						}
					}
				}
				return count
			}

			aware := palette.AvoidingAsANSI256(tc.entity, tc.reserved...)
			require.NotEmpty(t, aware)
			require.Zero(t, violations(aware))

			plain := tc.entity.Avoiding(tc.reserved...)
			require.Positive(t, violations(plain))
		})
	}
}

// distance returns the CIEDE2000 distance between two colors in go-colorful's
// scale, matching the metric used by Palette.Avoiding.
func distance(t *testing.T, a, b color.Color) float64 {
	t.Helper()

	ca, ok := colorful.MakeColor(a)
	require.True(t, ok)
	cb, ok := colorful.MakeColor(b)
	require.True(t, ok)

	return ca.DistanceCIEDE2000(cb)
}
