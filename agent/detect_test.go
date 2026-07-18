package agent_test

import (
	"testing"

	"github.com/gechr/x/agent"
	"github.com/stretchr/testify/require"
)

// clearDetectEnv blanks every environment variable consulted by Is and Detect
// so tests are isolated from any agent actually running them.
func clearDetectEnv(t *testing.T) {
	t.Helper()
	for _, v := range agent.EnvVars() {
		t.Setenv(v, "")
	}
}

func TestDetect_AgentNamed(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{"Goose", agent.Goose},
		{"amp", agent.Amp},
		{"claude", agent.Claude},
		{"goose", agent.Goose},
		{"some-future-agent", "some-future-agent"},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(agent.EnvAgent, tt.value)

			require.Equal(t, tt.want, agent.Detect())
			require.True(t, agent.Is())
		})
	}
}

func TestDetect_MarkerVars(t *testing.T) {
	tests := []struct {
		env  string
		want string
	}{
		{"CLAUDECODE", agent.Claude},
		{"CLINE_ACTIVE", agent.Cline},
		{"CODEX_SANDBOX", agent.Codex},
		{"CURSOR_AGENT", agent.Cursor},
		{"GEMINI_CLI", agent.Gemini},
		{"OPENCODE", agent.OpenCode},
		{"REPL_ID", agent.Replit},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(tt.env, "1")

			require.Equal(t, tt.want, agent.Detect())
			require.True(t, agent.Is())
		})
	}
}

// TestDetect_AnonymousTruthy: AGENT=1 asserts presence without identity, so
// Is reports true while Detect has no name to return.
func TestDetect_AnonymousTruthy(t *testing.T) {
	for _, value := range []string{"1", "true", "yes", "on", "TRUE"} {
		t.Run(value, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(agent.EnvAgent, value)

			require.True(t, agent.Is())
			require.Empty(t, agent.Detect())
		})
	}
}

// TestDetect_AnonymousTruthyFallsThroughToMarkers: AGENT=1 alone carries no
// name, but an agent-specific marker set alongside it still identifies the
// agent.
func TestDetect_AnonymousTruthyFallsThroughToMarkers(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(agent.EnvAgent, "1")
	t.Setenv("CLAUDECODE", "1")

	require.True(t, agent.Is())
	require.Equal(t, agent.Claude, agent.Detect())
}

// TestDetect_FalsyOptOut: a falsy AGENT value is an explicit opt-out that
// wins even over agent-specific markers, mirroring the CI=false escape hatch.
func TestDetect_FalsyOptOut(t *testing.T) {
	for _, value := range []string{"0", "false", "no", "off", "FALSE"} {
		t.Run(value, func(t *testing.T) {
			clearDetectEnv(t)
			t.Setenv(agent.EnvAgent, value)
			t.Setenv("CLAUDECODE", "1")

			require.False(t, agent.Is())
			require.Empty(t, agent.Detect())
		})
	}
}

func TestDetect_AIAgentVariant(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(agent.EnvAIAgent, "cursor")

	require.True(t, agent.Is())
	require.Equal(t, agent.Cursor, agent.Detect())
}

// TestDetect_AgentTakesPrecedenceOverAIAgent: the primary convention wins
// when both cross-tool variables are set.
func TestDetect_AgentTakesPrecedenceOverAIAgent(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(agent.EnvAgent, "goose")
	t.Setenv(agent.EnvAIAgent, "cursor")

	require.Equal(t, agent.Goose, agent.Detect())
}

// TestDetect_NamedAgentTakesPrecedenceOverMarkers: a named AGENT value is
// authoritative even when another agent's marker leaks through (e.g. one
// agent spawned from another's session).
func TestDetect_NamedAgentTakesPrecedenceOverMarkers(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(agent.EnvAgent, "goose")
	t.Setenv("CLAUDECODE", "1")

	require.Equal(t, agent.Goose, agent.Detect())
}

func TestDetect_EmptyWhenNothingSet(t *testing.T) {
	clearDetectEnv(t)

	require.False(t, agent.Is())
	require.Empty(t, agent.Detect())
}

// TestDetect_WhitespaceOnlyIgnored: a whitespace-only value is treated as
// unset, not as a named agent.
func TestDetect_WhitespaceOnlyIgnored(t *testing.T) {
	clearDetectEnv(t)
	t.Setenv(agent.EnvAgent, "   ")

	require.False(t, agent.Is())
	require.Empty(t, agent.Detect())
}

func TestKnown(t *testing.T) {
	known := agent.Known()
	require.NotEmpty(t, known)
	for _, name := range known {
		require.True(t, agent.IsKnown(name))
	}
}

func TestIsKnown_Unknown(t *testing.T) {
	require.False(t, agent.IsKnown("some-future-agent"))
	require.False(t, agent.IsKnown(""))
}
