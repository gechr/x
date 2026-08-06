package palette

import (
	"image/color"

	"github.com/charmbracelet/colorprofile"
	xslices "github.com/gechr/x/slices"
	colorful "github.com/lucasb-eyer/go-colorful"
)

// minReservedDistance is the minimum CIEDE2000 distance a palette color must
// keep from every reserved color to survive [Palette.Avoiding], in
// go-colorful's scale where a just-noticeable difference is roughly 0.01
// (conventional ΔE2000 divided by 100). 0.16 removes look-alikes that could
// be mistaken for a reserved color at a glance, while a four-color reserved
// set still leaves over half of the true-color palettes standing.
const minReservedDistance = 0.16

// Avoiding returns a new palette without the colors that sit perceptually
// close to any reserved color, so entity colors cannot be mistaken for a
// reserved one - a theme's semantic red, say. Closeness is measured with the
// CIEDE2000 color-difference formula. The remaining colors keep their
// original order, so a most-separable-first palette stays that way.
//
// Distances are measured on the colors' own RGBA values, so the guarantee is
// for opaque colors rendered in true color. Rendering in a reduced color
// profile quantizes colors and can narrow their distances; [Auto] accounts
// for this on its non-true-color path by measuring colors as ANSI-256
// renders them. Nil and fully transparent reserved colors are ignored;
// palette colors that cannot be measured are kept.
//
// The mapping from string to color differs between the filtered and
// unfiltered palettes, since [Palette.Color] hashes over the palette length.
// A reserved set that blankets the palette can leave it empty, in which case
// [Palette.Color] returns nil; use [Palette.Assigner] rather than
// [NewAssigner] to keep such a palette empty.
func (p Palette) Avoiding(reserved ...color.Color) Palette {
	return p.avoiding(nil, reserved...)
}

// avoiding filters like [Palette.Avoiding], measuring every color through
// `transform` first when it is non-nil, so distances reflect how the colors
// will actually render.
func (p Palette) avoiding(
	transform func(color.Color) color.Color, reserved ...color.Color,
) Palette {
	avoid := make([]colorful.Color, 0, len(reserved))
	for _, r := range reserved {
		if c, ok := measure(r, transform); ok {
			avoid = append(avoid, c)
		}
	}

	return xslices.Reject(p, func(pc color.Color) bool {
		c, ok := measure(pc, transform)
		if !ok {
			return false
		}
		return xslices.AnyFunc(avoid, func(r colorful.Color) bool {
			return c.DistanceCIEDE2000(r) < minReservedDistance
		})
	})
}

// ansi256 quantizes a color to the ANSI-256 palette, approximating how a
// terminal without true-color support will render it.
func ansi256(c color.Color) color.Color {
	return colorprofile.ANSI256.Convert(c)
}

// measure converts a color for distance measurement, applying `transform`
// first when it is non-nil. It reports failure for nil and fully transparent
// colors, which have no measurable hue.
func measure(
	c color.Color, transform func(color.Color) color.Color,
) (colorful.Color, bool) {
	if c == nil {
		return colorful.Color{}, false
	}
	if transform != nil {
		if c = transform(c); c == nil {
			return colorful.Color{}, false
		}
	}
	return colorful.MakeColor(c)
}
