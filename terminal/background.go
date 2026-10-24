package terminal

import (
	"math"
	"os"
	"sync"
)

const (
	// darkLuminanceThreshold is the WCAG relative luminance below which white
	// text contrasts more with the background than black text does. It is
	// sqrt(1.05 * 0.05) - 0.05, where the two contrast ratios are equal.
	darkLuminanceThreshold = 0.179128784748
	colorChannelMaximum    = 255

	// The sRGB transfer function and luminance weights from the WCAG
	// definition of relative luminance.
	srgbLinearThreshold  = 0.04045
	srgbLinearSlope      = 12.92
	srgbGammaOffset      = 0.055
	srgbGammaExponent    = 2.4
	redLuminanceWeight   = 0.2126
	greenLuminanceWeight = 0.7152
	blueLuminanceWeight  = 0.0722
)

var background = sync.OnceValue(func() backgroundResult {
	dark, ok := detectBackground(terminalFile(os.Stdout, os.Stderr, os.Stdin))
	return backgroundResult{dark: dark, ok: ok}
})

type backgroundResult struct {
	dark bool
	ok   bool
}

// IsDark reports (dark, ok) for the controlling terminal. It performs terminal
// I/O on the first call, returning as soon as the terminal has answered and
// waiting no longer than half a second for one that never does. The result,
// including no response, is cached for the process. `ok` is false if no
// standard stream is a terminal open for reading and writing, or the terminal
// does not respond, in which case the first result is meaningless.
func IsDark() (bool, bool) {
	result := background()
	return result.dark, result.ok
}

// IsLight reports (light, ok) for the controlling terminal. Like [IsDark], it
// performs terminal I/O on the first call and caches the result for the process.
// `ok` is false if no standard stream is a terminal open for reading and
// writing, or the terminal does not respond, in which case the first result is
// meaningless.
func IsLight() (bool, bool) {
	dark, ok := IsDark()
	return ok && !dark, ok
}

// terminalFile returns the first of `files` that can carry a background query,
// or nil if none can. The background is a property of the terminal itself, so
// any stream pointing at it answers the query - probing lets a redirected
// stdout fall through to stderr or stdin.
func terminalFile(files ...*os.File) *os.File {
	for _, f := range files {
		if canQuery(f) {
			return f
		}
	}
	return nil
}

// detectBackground queries `f` for its background color, returning whether it is
// dark and whether the query succeeded.
func detectBackground(f *os.File) (bool, bool) {
	if !canQuery(f) {
		return false, false
	}

	red, green, blue, ok := queryBackground(f)
	if !ok {
		return false, false
	}

	return isDarkColor(red, green, blue), true
}

// isDarkColor reports whether white text contrasts more with the given
// background than black text does.
func isDarkColor(red, green, blue uint8) bool {
	return relativeLuminance(red, green, blue) < darkLuminanceThreshold
}

// relativeLuminance returns the WCAG relative luminance of an sRGB color, from
// 0 (black) to 1 (white).
func relativeLuminance(red, green, blue uint8) float64 {
	return redLuminanceWeight*linearize(red) +
		greenLuminanceWeight*linearize(green) +
		blueLuminanceWeight*linearize(blue)
}

// linearize converts an 8-bit sRGB channel to linear light.
func linearize(channel uint8) float64 {
	value := float64(channel) / colorChannelMaximum
	if value <= srgbLinearThreshold {
		return value / srgbLinearSlope
	}
	return math.Pow((value+srgbGammaOffset)/(1+srgbGammaOffset), srgbGammaExponent)
}
