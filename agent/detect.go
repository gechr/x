// Package agent detects AI coding agents hosting the process.
package agent

import (
	"os"
	"strings"

	xstrings "github.com/gechr/x/strings"
)

// Environment variables consulted by [Is] and [Detect].
const (
	// EnvAgent is the proposed cross-tool convention
	// (https://github.com/agentsmd/agents.md/issues/136), set either to the
	// agent's name (e.g. `goose`, `amp`) or a truthy token like `1`.
	EnvAgent = "AGENT"
	// EnvAIAgent is the more specific variant used by Vercel's detection
	// tooling; consulted with the same semantics as [EnvAgent].
	EnvAIAgent = "AI_AGENT"
)

// standardVars lists the cross-tool variables, checked in order before any
// agent-specific marker.
var standardVars = []string{EnvAgent, EnvAIAgent}

// markerVars lists agent-specific environment variables, checked in order,
// covering tools that predate the `AGENT` convention. `REPL_ID` marks a Replit
// workspace rather than the Replit Agent itself, but is the established signal
// (a human in a Replit shell is reported as agent-hosted).
var markerVars = []struct {
	env  string
	name string
}{
	{"CLAUDECODE", Claude},
	{"CLINE_ACTIVE", Cline},
	{"CODEX_SANDBOX", Codex},
	{"CURSOR_AGENT", Cursor},
	{"GEMINI_CLI", Gemini},
	{"OPENCODE", OpenCode},
	{"REPL_ID", Replit},
}

// Is reports whether the process appears to be running under an AI coding
// agent. Detection is best-effort, based on environment variables set by the
// hosting agent: the cross-tool `AGENT` / `AI_AGENT` convention first, then
// agent-specific markers (`CLAUDECODE`, `CODEX_SANDBOX`, `CURSOR_AGENT`, ...).
// A falsy `AGENT` / `AI_AGENT` value (`0`, `false`, `no`, `off`) is an
// explicit opt-out and reports false even when a marker variable is present,
// mirroring the `CI=false` escape hatch.
func Is() bool {
	for _, env := range standardVars {
		if v := os.Getenv(env); strings.TrimSpace(v) != "" {
			return !xstrings.IsFalsy(v)
		}
	}
	for _, m := range markerVars {
		if os.Getenv(m.env) != "" {
			return true
		}
	}
	return false
}

// Detect returns the name of the AI coding agent hosting the process, or
// empty if there is none or it cannot be identified. A named `AGENT` /
// `AI_AGENT` value is returned verbatim (lowercased), so agents unknown to
// this package still identify themselves; a bare truthy token like `AGENT=1`
// signals presence without identity, so Detect falls through to the marker
// variables and may return empty even though [Is] reports true.
func Detect() string {
	for _, env := range standardVars {
		v := strings.ToLower(strings.TrimSpace(os.Getenv(env)))
		switch {
		case v == "":
			continue
		case xstrings.IsFalsy(v):
			return ""
		case xstrings.IsTruthy(v):
			// Anonymous opt-in; an agent-specific marker may still name it.
		default:
			return v
		}
	}
	for _, m := range markerVars {
		if os.Getenv(m.env) != "" {
			return m.name
		}
	}
	return ""
}
