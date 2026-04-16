package ansi

import (
	"os"

	"github.com/gechr/x/terminal"
)

// ANSI produces ANSI-aware output, falling back to plain text
// when the output is not a terminal.
type ANSI struct {
	terminal          bool
	hyperlinkFallback HyperlinkFallback
}

// New creates an ANSI with the given options.
func New(opts ...Option) *ANSI {
	w := &ANSI{}
	for _, o := range opts {
		o(w)
	}
	return w
}

// Never creates an ANSI with ANSI output unconditionally disabled.
func Never() *ANSI {
	return &ANSI{}
}

// Force creates an ANSI with ANSI output unconditionally enabled.
func Force() *ANSI {
	return &ANSI{terminal: true}
}

// Auto creates an ANSI that auto-detects whether the output is a terminal.
// All provided files must be terminals for ANSI output to be enabled.
// Defaults to os.Stdout if no files are provided.
func Auto(files ...*os.File) *ANSI {
	if len(files) == 0 {
		files = []*os.File{os.Stdout}
	}
	for _, f := range files {
		if f == nil || !terminal.Is(f) {
			return Never()
		}
	}
	return Force()
}

// Terminal reports whether the output target is a terminal.
func (w *ANSI) Terminal() bool { return w.terminal }
