package emulator_test

import (
	"testing"

	"github.com/gechr/x/terminal/emulator"
	"github.com/stretchr/testify/require"
)

func TestSupportsGraphemes_Term(t *testing.T) {
	tests := []struct {
		term string
		want bool
	}{
		{"alacritty", false},
		{"contour", true},
		{"foot", true},
		{"rio", true},
		{"rxvt-unicode", false},
		{"st", false},
		{"wezterm", true},
		{"xterm-ghostty", true},
		{"xterm-kitty", true},
	}
	for _, tt := range tests {
		t.Run(tt.term, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(emulator.EnvTerm, tt.term)

			require.Equal(t, tt.want, emulator.SupportsGraphemes())
		})
	}
}

func TestSupportsGraphemes_MarkerVars(t *testing.T) {
	tests := []struct {
		env  string
		want bool
	}{
		{"ITERM_SESSION_ID", false},
		{"KITTY_WINDOW_ID", true},
		{"WEZTERM_PANE", true},
		{"WT_SESSION", true},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(tt.env, "some-value")

			require.Equal(t, tt.want, emulator.SupportsGraphemes())
		})
	}
}

// TestSupportsGraphemes_Tmux: inside tmux the multiplexer's own wcwidth-based
// screen model governs, no matter how capable the outer terminal is.
func TestSupportsGraphemes_Tmux(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv("TMUX", "/private/tmp/tmux-501/default,4864,0")
	t.Setenv(emulator.EnvTerm, "tmux-256color")
	t.Setenv("KITTY_WINDOW_ID", "1")

	require.False(t, emulator.SupportsGraphemes())
}

func TestSupportsGraphemes_Undetected(t *testing.T) {
	clearDetectEnv(t)

	require.False(t, emulator.SupportsGraphemes())
}
