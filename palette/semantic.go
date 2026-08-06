package palette

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Semantic holds measured colors for conventional message meanings. Like the
// entity palettes, every color clears 4.5:1 contrast against its target
// background, and the colors are mutually distinguishable. Which strings
// warrant a semantic color is the caller's policy; the set only guarantees
// the colors themselves are legible and distinct.
//
// To keep entity colors perceptually clear of the set, reserve it: pass the
// [Semantic.Colors] slice to [WithReserved] or [Palette.Avoiding].
type Semantic struct {
	Danger  color.Color
	Warning color.Color
	Success color.Color
	Info    color.Color
}

// SemanticDark returns the semantic set tuned for dark backgrounds, measured
// against #1e1e1e like [TrueColorDark].
func SemanticDark() Semantic {
	return Semantic{
		Danger:  lipgloss.Color("#f87171"),
		Warning: lipgloss.Color("#fbbf24"),
		Success: lipgloss.Color("#4ade80"),
		Info:    lipgloss.Color("#60a5fa"),
	}
}

// SemanticLight returns the semantic set tuned for light backgrounds,
// measured against #fafafa like [TrueColorLight].
func SemanticLight() Semantic {
	return Semantic{
		Danger:  lipgloss.Color("#b91c1c"),
		Warning: lipgloss.Color("#975a04"),
		Success: lipgloss.Color("#15803d"),
		Info:    lipgloss.Color("#1d4ed8"),
	}
}

// Colors returns the set as a slice, in Danger, Warning, Success, Info
// order - convenient for [Palette.Avoiding] and [WithReserved].
func (s Semantic) Colors() []color.Color {
	return []color.Color{s.Danger, s.Warning, s.Success, s.Info}
}
