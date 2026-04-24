package ansi_test

import (
	"testing"

	"github.com/gechr/x/ansi"
	"github.com/stretchr/testify/assert"
)

// TestEscapeLiterals pins the byte-exact values of escapes we own outright.
func TestEscapeLiterals(t *testing.T) {
	assert.Equal(t, "\x1b[", ansi.CSI)
	assert.Equal(t, "\x1b[2K\r", ansi.ClearLine)
}

func TestEraseFunctions(t *testing.T) {
	assert.Equal(t, "\x1b[J", ansi.EraseDisplay(0))
	assert.Equal(t, "\x1b[2K", ansi.EraseLine(2))
	assert.Equal(t, "\x1b[3X", ansi.EraseCharacter(3))
}

func TestCursorMovement(t *testing.T) {
	assert.Equal(t, "\x1b[3A", ansi.CursorUp(3))
	assert.Equal(t, "\x1b[5B", ansi.CursorDown(5))
	assert.Equal(t, "\x1b[2C", ansi.CursorForward(2))
	assert.Equal(t, "\x1b[4D", ansi.CursorBackward(4))
	assert.Equal(t, "\x1b[7E", ansi.CursorNextLine(7))
	assert.Equal(t, "\x1b[2F", ansi.CursorPreviousLine(2))
	assert.Equal(t, "\x1b[9G", ansi.CursorHorizontalAbsolute(9))
	assert.Equal(t, "\x1b[3;10H", ansi.CursorPosition(10, 3))
}

func TestModeConstants(t *testing.T) {
	assert.Equal(t, "\x1b[?25h", ansi.ShowCursor)
	assert.Equal(t, "\x1b[?25l", ansi.HideCursor)
	assert.Equal(t, "\x1b[?1049h", ansi.EnterAltScreen)
	assert.Equal(t, "\x1b[?1049l", ansi.ExitAltScreen)
	assert.Equal(t, "\x1b[?2004h", ansi.EnableBracketedPaste)
	assert.Equal(t, "\x1b[?2004l", ansi.DisableBracketedPaste)
	assert.Equal(t, "\x1b[?1004h", ansi.EnableFocusEvents)
	assert.Equal(t, "\x1b[?1004l", ansi.DisableFocusEvents)
	assert.Equal(t, "\x1b[6n", ansi.RequestCursorPosition)
	assert.Equal(t, "\x1b[?6n", ansi.RequestExtendedCursorPosition)
	assert.Equal(t, "\x1b[I", ansi.Focus)
	assert.Equal(t, "\x1b[O", ansi.Blur)
	assert.Equal(t, "\x1b[H", ansi.CursorHomePosition)
	assert.Equal(t, "\x1b[s", ansi.SaveCursorPosition)
	assert.Equal(t, "\x1b[u", ansi.RestoreCursorPosition)
}

func TestScrollAndEdit(t *testing.T) {
	assert.Equal(t, "\x1b[2S", ansi.ScrollUp(2))
	assert.Equal(t, "\x1b[3T", ansi.ScrollDown(3))
	assert.Equal(t, "\x1b[2L", ansi.InsertLine(2))
	assert.Equal(t, "\x1b[4M", ansi.DeleteLine(4))
	assert.Equal(t, "\x1b[2@", ansi.InsertCharacter(2))
	assert.Equal(t, "\x1b[5P", ansi.DeleteCharacter(5))
}

func TestStyleAndCursorStyle(t *testing.T) {
	assert.Equal(t, "\x1b[m", ansi.ResetStyle)
	assert.Equal(t, "\x1b[2 q", ansi.SetCursorStyle(2))
	assert.Equal(t, "\x1b[5 q", ansi.SetCursorStyle(5))
}

func TestTitleSequences(t *testing.T) {
	assert.Equal(t, "\x1b]2;hello\x07", ansi.SetWindowTitle("hello"))
	assert.Equal(t, "\x1b]1;hello\x07", ansi.SetIconName("hello"))
	assert.Equal(t, "\x1b]0;hello\x07", ansi.SetIconNameWindowTitle("hello"))
}

func TestKeypadAndVersionQuery(t *testing.T) {
	assert.Equal(t, "\x1b=", ansi.KeypadApplicationMode)
	assert.Equal(t, "\x1b>", ansi.KeypadNumericMode)
	assert.Equal(t, "\x1b[>q", ansi.RequestNameVersion)
}

func TestMouseModes(t *testing.T) {
	assert.Equal(t, "\x1b[?9h", ansi.EnableMouseX10)
	assert.Equal(t, "\x1b[?9l", ansi.DisableMouseX10)
	assert.Equal(t, "\x1b[?1000h", ansi.EnableMouseNormal)
	assert.Equal(t, "\x1b[?1000l", ansi.DisableMouseNormal)
	assert.Equal(t, "\x1b[?1002h", ansi.EnableMouseButtonEvent)
	assert.Equal(t, "\x1b[?1002l", ansi.DisableMouseButtonEvent)
	assert.Equal(t, "\x1b[?1003h", ansi.EnableMouseAnyEvent)
	assert.Equal(t, "\x1b[?1003l", ansi.DisableMouseAnyEvent)
	assert.Equal(t, "\x1b[?1006h", ansi.EnableMouseSGR)
	assert.Equal(t, "\x1b[?1006l", ansi.DisableMouseSGR)
}
