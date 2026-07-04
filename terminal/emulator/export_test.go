package emulator

// EnvVars returns every environment variable consulted by Detect.
// Used by tests to isolate the environment.
func EnvVars() []string {
	vars := []string{EnvTerm, EnvTermProgram, EnvTerminalEmulator}
	for _, m := range markerVars {
		vars = append(vars, m.env)
	}
	return vars
}
