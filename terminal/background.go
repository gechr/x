package terminal

import (
	"os"
	"sync"
)

// darkLightnessThreshold is the HSL lightness below which a background is
// considered dark. 0.5 splits the range evenly between black and white.
const (
	darkLightnessThreshold = 0.5
	colorChannelMaximum    = 255
	lightnessExtremaCount  = 2
)

var background = sync.OnceValue(func() backgroundResult {
	dark, ok := detectBackground(terminalFile())
	return backgroundResult{dark: dark, ok: ok}
})

type backgroundResult struct {
	dark bool
	ok   bool
}

// IsDark reports (dark, ok) for the controlling terminal. It performs terminal
// I/O on the first call, waiting up to 10 milliseconds for a background-color
// response. The result, including no response, is cached for the process. `ok`
// is false if no standard stream is a terminal or the terminal does not
// respond, in which case the first result is meaningless.
func IsDark() (bool, bool) {
	result := background()
	return result.dark, result.ok
}

// IsLight reports (light, ok) for the controlling terminal. Like [IsDark], it
// performs terminal I/O on the first call and caches the result for the process.
// `ok` is false if no standard stream is a terminal or the terminal does not
// respond, in which case the first result is meaningless.
func IsLight() (bool, bool) {
	dark, ok := IsDark()
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

	red, green, blue, ok := queryBackground(f)
	if !ok {
		return false, false
	}

	lightness := float64(int(max(red, green, blue))+int(min(red, green, blue))) /
		(lightnessExtremaCount * colorChannelMaximum)
	return lightness < darkLightnessThreshold, true
}
