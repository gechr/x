package emulator_test

import (
	"testing"

	"github.com/gechr/x/terminal/emulator"
	"github.com/stretchr/testify/require"
)

// clearDetectEnv blanks every environment variable consulted by Detect so
// tests are isolated from the terminal actually running them.
func clearDetectEnv(t *testing.T) {
	t.Helper()
	for _, v := range emulator.EnvVars() {
		t.Setenv(v, "")
	}
}

func TestDetect_TermProgram(t *testing.T) {
	tests := []struct {
		program string
		want    string
	}{
		{"Apple_Terminal", emulator.AppleTerminal},
		{"Hyper", emulator.Hyper},
		{"Tabby", emulator.Tabby},
		{"WarpTerminal", emulator.Warp},
		{"WezTerm", emulator.WezTerm},
		{"contour", emulator.Contour},
		{"ghostty", emulator.Ghostty},
		{"iTerm.app", emulator.ITerm2},
		{"mintty", emulator.Mintty},
		{"vscode", emulator.VSCode},
		{"zed", emulator.Zed},
	}
	for _, tt := range tests {
		t.Run(tt.program, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(emulator.EnvTermProgram, tt.program)

			require.Equal(t, tt.want, emulator.Detect())
		})
	}
}

func TestDetect_MarkerVars(t *testing.T) {
	tests := []struct {
		env  string
		want string
	}{
		{"ALACRITTY_LOG", emulator.Alacritty},
		{"ALACRITTY_WINDOW_ID", emulator.Alacritty},
		{"ConEmuPID", emulator.ConEmu},
		{"GHOSTTY_RESOURCES_DIR", emulator.Ghostty},
		{"GNOME_TERMINAL_SCREEN", emulator.GNOMETerminal},
		{"ITERM_SESSION_ID", emulator.ITerm2},
		{"KITTY_WINDOW_ID", emulator.Kitty},
		{"KONSOLE_DBUS_SERVICE", emulator.Konsole},
		{"KONSOLE_VERSION", emulator.Konsole},
		{"TERMINATOR_UUID", emulator.Terminator},
		{"TERMUX_VERSION", emulator.Termux},
		{"TILIX_ID", emulator.Tilix},
		{"WEZTERM_PANE", emulator.WezTerm},
		{"WT_SESSION", emulator.WindowsTerminal},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(tt.env, "some-value")

			require.Equal(t, tt.want, emulator.Detect())
		})
	}
}

func TestDetect_Multiplexers(t *testing.T) {
	tests := []struct {
		env  string
		want string
	}{
		{"TMUX", emulator.Tmux},
		{"STY", emulator.Screen},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(tt.env, "some-value")

			require.Equal(t, tt.want, emulator.Detect())
		})
	}
}

// TestDetect_MultiplexerTakesPrecedence: a multiplexer owns the screen model
// of everything inside it, so it must win over the hosting emulator's
// variables - including a tmux config that sets TERM to "screen-256color".
func TestDetect_MultiplexerTakesPrecedence(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv("TMUX", "/private/tmp/tmux-501/default,4864,0")
	t.Setenv(emulator.EnvTerm, "screen-256color")
	t.Setenv(emulator.EnvTermProgram, "iTerm.app")
	t.Setenv("KITTY_WINDOW_ID", "1")

	require.Equal(t, emulator.Tmux, emulator.Detect())
}

// TestDetect_NestedMultiplexers: when multiplexers are nested, both markers
// are present (neither scrubs the other's variable) and the environment
// cannot identify the innermost one. tmux is reported as the more common
// case; this pins the documented tie-break, not innermost-mux detection.
func TestDetect_NestedMultiplexers(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv("TMUX", "/private/tmp/tmux-501/default,4864,0")
	t.Setenv("STY", "12345.pts-0.host")
	t.Setenv(emulator.EnvTerm, "screen")

	require.Equal(t, emulator.Tmux, emulator.Detect())
}

func TestDetect_Term(t *testing.T) {
	tests := []struct {
		term string
		want string
	}{
		{"alacritty", emulator.Alacritty},
		{"alacritty-direct", emulator.Alacritty},
		{"contour", emulator.Contour},
		{"foot", emulator.Foot},
		{"foot-extra", emulator.Foot},
		{"foot-extra-direct", emulator.Foot},
		{"rxvt-unicode", emulator.URxvt},
		{"rxvt-unicode-256color", emulator.URxvt},
		{"screen", emulator.Screen},
		{"screen-256color", emulator.Screen},
		{"st", emulator.ST},
		{"st-256color", emulator.ST},
		{"tmux", emulator.Tmux},
		{"tmux-256color", emulator.Tmux},
		{"wezterm", emulator.WezTerm},
		{"xterm-ghostty", emulator.Ghostty},
		{"xterm-kitty", emulator.Kitty},
	}
	for _, tt := range tests {
		t.Run(tt.term, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(emulator.EnvTerm, tt.term)

			require.Equal(t, tt.want, emulator.Detect())
		})
	}
}

func TestDetect_TerminalEmulatorJetBrains(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTerminalEmulator, "JetBrains-JediTerm")

	require.Equal(t, emulator.JetBrains, emulator.Detect())
}

func TestDetect_TerminalEmulatorUnknownIgnored(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTerminalEmulator, "SomethingElse")

	require.Empty(t, emulator.Detect())
}

// TestDetect_TermTakesPrecedence: TERM names the innermost emulator; an
// inherited TERM_PROGRAM or marker variable from the launching terminal must
// not override it. TERM=xterm-kitty alongside TERM_PROGRAM=ghostty means
// kitty was launched from Ghostty (kitty sets TERM but scrubs neither).
func TestDetect_TermTakesPrecedence(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTermProgram, "ghostty")
	t.Setenv("ITERM_SESSION_ID", "w0t0p0")
	t.Setenv(emulator.EnvTerm, "xterm-kitty")

	require.Equal(t, emulator.Kitty, emulator.Detect())
}

// TestDetect_KittyLaunchedFromITerm2 mirrors the environment observed in a
// real kitty 0.47.4 instance launched from iTerm2: kitty sets TERM and its
// own markers but inherits iTerm2's TERM_PROGRAM and ITERM_SESSION_ID.
func TestDetect_KittyLaunchedFromITerm2(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTermProgram, "iTerm.app")
	t.Setenv("ITERM_SESSION_ID", "w0t1p1:FACC5BAE")
	t.Setenv("KITTY_WINDOW_ID", "1")
	t.Setenv(emulator.EnvTerm, "xterm-kitty")

	require.Equal(t, emulator.Kitty, emulator.Detect())
}

func TestDetect_TermProgramTakesPrecedenceOverMarkers(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTermProgram, "iTerm.app")
	t.Setenv("KITTY_WINDOW_ID", "1")
	t.Setenv(emulator.EnvTerm, "xterm-256color")

	require.Equal(t, emulator.ITerm2, emulator.Detect())
}

func TestDetect_UnknownTermProgramFallsThrough(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTermProgram, "tmux")
	t.Setenv("ITERM_SESSION_ID", "w0t0p0")

	require.Equal(t, emulator.ITerm2, emulator.Detect())
}

func TestDetect_EmptyWhenNothingSet(t *testing.T) {
	clearDetectEnv(t)

	require.Empty(t, emulator.Detect())
}

func TestDetect_UnknownValuesIgnored(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(emulator.EnvTermProgram, "not-a-terminal")
	t.Setenv(emulator.EnvTerm, "xterm-256color")

	require.Empty(t, emulator.Detect())
}

func TestKnown(t *testing.T) {
	known := emulator.Known()
	require.NotEmpty(t, known)
	for _, name := range known {
		require.True(t, emulator.IsKnown(name))
	}
}

func TestIsKnown_Unknown(t *testing.T) {
	require.False(t, emulator.IsKnown("xterm-256color"))
	require.False(t, emulator.IsKnown(""))
}
