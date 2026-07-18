package agent

// EnvVars returns every environment variable consulted by Is and Detect.
// Used by tests to isolate the environment.
func EnvVars() []string {
	vars := make([]string, 0, len(standardVars)+len(markerVars))
	vars = append(vars, standardVars...)
	for _, m := range markerVars {
		vars = append(vars, m.env)
	}
	return vars
}
