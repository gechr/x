package agent

import (
	"slices"

	"github.com/gechr/x/set"
)

// Recognized AI coding agent names, as returned by [Detect].
const (
	Amp      = "amp"
	Claude   = "claude"
	Cline    = "cline"
	Codex    = "codex"
	Cursor   = "cursor"
	Gemini   = "gemini"
	Goose    = "goose"
	OpenCode = "opencode"
	Replit   = "replit"
)

var knownAgents = []string{
	Amp,
	Claude,
	Cline,
	Codex,
	Cursor,
	Gemini,
	Goose,
	OpenCode,
	Replit,
}

var knownAgentSet = set.New(knownAgents...)

// Known returns the set of recognized AI coding agent names.
func Known() []string {
	return slices.Clone(knownAgents)
}

// IsKnown reports whether `name` matches a known AI coding agent.
func IsKnown(name string) bool {
	return knownAgentSet.Contains(name)
}
