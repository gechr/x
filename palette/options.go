package palette

// config holds the resolved settings for [Auto].
type config struct {
	trueColor *bool // nil = auto-detect terminal support
}

// Option configures [Auto].
type Option func(*config)

// WithTrueColor forces [Auto] to return the 256-color Glasbey palette,
// overriding its automatic terminal-capability detection. The caller is
// responsible for ensuring the terminal supports true color.
func WithTrueColor() Option {
	return func(c *config) {
		on := true
		c.trueColor = &on
	}
}
