package ansi

import xansi "github.com/charmbracelet/x/ansi"

// HyperlinkFallback controls how hyperlinks render when the output is not a terminal.
type HyperlinkFallback int

const (
	// HyperlinkFallbackExpanded renders "text (url)".
	HyperlinkFallbackExpanded HyperlinkFallback = iota
	// HyperlinkFallbackMarkdown renders "[text](url)".
	HyperlinkFallbackMarkdown
	// HyperlinkFallbackText renders only the display text, discarding the URL.
	HyperlinkFallbackText
	// HyperlinkFallbackURL renders only the URL, discarding the display text.
	HyperlinkFallbackURL
)

// WithHyperlinkFallback sets how hyperlinks render when the output is not a terminal.
func WithHyperlinkFallback(fallback HyperlinkFallback) Option {
	return func(w *ANSI) {
		w.hyperlinkFallback = fallback
	}
}

// Hyperlink creates an OSC 8 terminal hyperlink.
// When the output is not a terminal, the HyperlinkFallback mode controls
// how the link is rendered in plain text.
func (w *ANSI) Hyperlink(url, text string) string {
	if !w.terminal {
		switch w.hyperlinkFallback {
		case HyperlinkFallbackExpanded:
			return text + " (" + url + ")"
		case HyperlinkFallbackMarkdown:
			return "[" + text + "](" + url + ")"
		case HyperlinkFallbackText:
			return text
		case HyperlinkFallbackURL:
			return url
		default:
			return text + " (" + url + ")"
		}
	}
	return xansi.SetHyperlink(url) + text + xansi.ResetHyperlink()
}
