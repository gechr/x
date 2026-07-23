package terminal_test

import (
	"testing"

	"github.com/gechr/x/terminal"
	"github.com/stretchr/testify/require"
)

func TestSupportsTrueColor(t *testing.T) {
	t.Run("COLORTERM truecolor", func(t *testing.T) {
		t.Setenv("COLORTERM", "truecolor")
		require.True(t, terminal.SupportsTrueColor())
	})

	t.Run("COLORTERM 24bit", func(t *testing.T) {
		t.Setenv("COLORTERM", "24bit")
		require.True(t, terminal.SupportsTrueColor())
	})

	t.Run("256color is not true color", func(t *testing.T) {
		t.Setenv("COLORTERM", "")
		t.Setenv("TERM", "xterm-256color")
		require.False(t, terminal.SupportsTrueColor())
	})

	t.Run("TERM advertising truecolor", func(t *testing.T) {
		t.Setenv("COLORTERM", "")
		t.Setenv("TERM", "xterm-truecolor")
		require.True(t, terminal.SupportsTrueColor())
	})
}
