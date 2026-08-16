package palette

import "image/color"

// MinReservedDistance exposes the [Palette.Avoiding] threshold to tests.
const MinReservedDistance = minReservedDistance

// AvoidingAsANSI256 exposes the ANSI-256-aware filter [Auto] uses on its
// non-true-color path to tests.
func AvoidingAsANSI256(p Palette, reserved ...color.Color) Palette {
	return p.avoiding(ansi256, reserved...)
}

// ANSI256 exposes the render-time quantization to tests.
func ANSI256(c color.Color) color.Color {
	return ansi256(c)
}

// WithoutTrueColor is the inverse of [WithTrueColor], sending [Auto] down its
// ANSI-256 path whatever the terminal supports. Tests need it because
// detection is cached for the process, so the environment cannot pin it.
func WithoutTrueColor() Option {
	return func(c *config) {
		off := false
		c.trueColor = &off
	}
}
