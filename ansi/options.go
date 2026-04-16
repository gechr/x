package ansi

// Option configures an ANSI.
type Option func(*ANSI)

// WithTerminal sets whether the output target is a terminal.
func WithTerminal(v bool) Option {
	return func(w *ANSI) {
		w.terminal = v
	}
}
