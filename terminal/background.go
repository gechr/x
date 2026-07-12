package terminal

import (
	"os"

	"github.com/muesli/termenv"
)

// darkLightnessThreshold is the HSL lightness below which a background is
// considered dark. 0.5 splits the range evenly between black and white.
const darkLightnessThreshold = 0.5

// IsDark reports (dark, ok) for the controlling terminal. `ok` is false if no
// standard stream is a terminal or the terminal does not respond to the
// background-color query, in which case the first result is meaningless.
func IsDark() (bool, bool) {
	return detectBackground(terminalFile())
}

// IsLight reports (light, ok) for the controlling terminal. `ok` is false if no
// standard stream is a terminal or the terminal does not respond to the
// background-color query, in which case the first result is meaningless.
func IsLight() (bool, bool) {
	dark, ok := detectBackground(terminalFile())
	return ok && !dark, ok
}

// terminalFile returns the first standard stream connected to a terminal, or
// nil if none is. The background is a property of the terminal itself, so any
// stream pointing at it answers the query - probing lets a redirected stdout
// fall through to stderr or stdin.
func terminalFile() *os.File {
	for _, f := range []*os.File{os.Stdout, os.Stderr, os.Stdin} {
		if Is(f) {
			return f
		}
	}
	return nil
}

// detectBackground queries `f` for its background color, returning whether it is
// dark and whether the query succeeded.
func detectBackground(f *os.File) (bool, bool) {
	if !Is(f) {
		return false, false
	}

	bg := termenv.NewOutput(f, termenv.WithTTY(true)).BackgroundColor()
	if _, isNoColor := bg.(termenv.NoColor); isNoColor {
		return false, false
	}

	_, _, lightness := termenv.ConvertToRGB(bg).Hsl()
	return lightness < darkLightnessThreshold, true
}
