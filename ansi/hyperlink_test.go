package ansi_test

import (
	"testing"

	xansi "github.com/charmbracelet/x/ansi"
	"github.com/gechr/x/ansi"
	"github.com/stretchr/testify/require"
)

func TestHyperlink_Never(t *testing.T) {
	w := ansi.Never()
	got := w.Hyperlink("https://example.com", "link")
	// Default fallback is parentheses.
	require.Equal(t, "link (https://example.com)", got)
}

func TestHyperlink_FallbackMarkdown(t *testing.T) {
	w := ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallbackMarkdown))
	got := w.Hyperlink("https://example.com", "link")
	require.Equal(t, "[link](https://example.com)", got)
}

func TestHyperlink_FallbackExpanded(t *testing.T) {
	w := ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallbackExpanded))
	got := w.Hyperlink("https://example.com", "link")
	require.Equal(t, "link (https://example.com)", got)
}

func TestHyperlink_FallbackText(t *testing.T) {
	w := ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallbackText))
	got := w.Hyperlink("https://example.com", "link")
	require.Equal(t, "link", got)
}

func TestHyperlink_FallbackURL(t *testing.T) {
	w := ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallbackURL))
	got := w.Hyperlink("https://example.com", "link")
	require.Equal(t, "https://example.com", got)
}

func TestHyperlink_FallbackUnknownDefaultsToExpanded(t *testing.T) {
	w := ansi.New(ansi.WithHyperlinkFallback(ansi.HyperlinkFallback(99)))
	got := w.Hyperlink("https://example.com", "link")
	require.Equal(t, "link (https://example.com)", got)
}

func TestHyperlink_Force(t *testing.T) {
	got := ansi.Force().Hyperlink("https://example.com", "link")

	// Should contain OSC 8 sequences.
	expected := xansi.SetHyperlink("https://example.com") + "link" + xansi.ResetHyperlink()
	require.Equal(t, expected, got)

	// Visible text should still be "link".
	require.Equal(t, "link", xansi.Strip(got))
}
