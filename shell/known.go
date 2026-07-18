package shell

import (
	"slices"

	"github.com/gechr/x/set"
)

// Recognized shell names, as returned by [Known].
const (
	Ash    = "ash"
	Bash   = "bash"
	Dash   = "dash"
	Elvish = "elvish"
	Fish   = "fish"
	Ksh    = "ksh"
	Nu     = "nu"
	Pwsh   = "pwsh"
	Sh     = "sh"
	Tcsh   = "tcsh"
	Zsh    = "zsh"
)

var knownShells = []string{
	Ash,
	Bash,
	Dash,
	Elvish,
	Fish,
	Ksh,
	Nu,
	Pwsh,
	Sh,
	Tcsh,
	Zsh,
}

var knownShellSet = set.New(knownShells...)

// Known returns the set of recognized shell names.
func Known() []string {
	return slices.Clone(knownShells)
}

// IsKnown reports whether `name` matches a known shell.
func IsKnown(name string) bool {
	return knownShellSet.Contains(name)
}
