package palette

import "image/color"

// config holds the resolved settings for [Auto].
type config struct {
	trueColor *bool // nil = auto-detect terminal support
	dark      *bool // nil = auto-detect terminal background
	reserved  []color.Color
}

// Option configures [Auto].
type Option func(*config)

// WithTrueColor forces [Auto] to return the true-color palette,
// overriding its automatic terminal-capability detection. The caller is
// responsible for ensuring the terminal supports true color.
func WithTrueColor() Option {
	return func(c *config) {
		on := true
		c.trueColor = &on
	}
}

// WithDark tells [Auto] whether the terminal background is dark,
// skipping its background detection entirely. Use it when the background was
// already detected, for example once at startup.
func WithDark(dark bool) Option {
	return func(c *config) {
		c.dark = &dark
	}
}

// WithReserved keeps the returned palette perceptually clear of the reserved
// colors, as if the selected palette were passed to [Palette.Avoiding]: any
// palette color that could be mistaken for a reserved one is removed. Use it
// when entity colors render alongside colors that carry meaning of their own,
// such as a theme's semantic red or a [Semantic] set. On the non-true-color
// path, colors are measured as ANSI-256 renders them.
func WithReserved(reserved ...color.Color) Option {
	return func(c *config) {
		c.reserved = append(c.reserved, reserved...)
	}
}
