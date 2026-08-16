package terminal_test

import (
	"os"
	"testing"

	"github.com/gechr/x/terminal"
	"github.com/stretchr/testify/require"
)

func TestDetectTrueColor(t *testing.T) {
	t.Run("COLORTERM truecolor", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{
			"TERM=xterm-256color", "COLORTERM=truecolor",
		}))
	})

	t.Run("COLORTERM 24bit", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{
			"TERM=xterm-256color", "COLORTERM=24bit",
		}))
	})

	t.Run("256color is not true color", func(t *testing.T) {
		require.False(t, terminal.DetectTrueColor([]string{"TERM=xterm-256color"}))
	})

	t.Run("TERM advertising truecolor", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{"TERM=xterm-truecolor"}))
	})

	t.Run("direct color TERM", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{"TERM=xterm-direct"}))
	})

	t.Run("terminal known to render 24-bit", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{"TERM=xterm-ghostty"}))
	})

	t.Run("Windows Terminal", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{
			"TERM=xterm-256color", "WT_SESSION=6ec7d5d1",
		}))
	})

	t.Run("NO_COLOR is preference, not capability", func(t *testing.T) {
		require.True(t, terminal.DetectTrueColor([]string{
			"TERM=xterm-256color", "COLORTERM=truecolor", "NO_COLOR=1",
		}))
	})

	t.Run("dumb terminal", func(t *testing.T) {
		require.False(t, terminal.DetectTrueColor([]string{
			"TERM=dumb", "COLORTERM=truecolor",
		}))
	})

	t.Run("empty environment", func(t *testing.T) {
		require.False(t, terminal.DetectTrueColor(nil))
	})
}

func TestSupportsTrueColor(t *testing.T) {
	require.Equal(t, terminal.DetectTrueColor(os.Environ()), terminal.SupportsTrueColor())
}
