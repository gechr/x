package shell

import (
	"fmt"
	"path/filepath"
)

// CompletionFile returns the standard completion file path for the given
// command and shell.
func CompletionFile(command, sh string) (string, error) {
	switch sh {
	case Bash:
		dir, err := DataDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "bash-completion", "completions", command), nil
	case Zsh:
		dir, err := DataDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "zsh", "site-functions", "_"+command), nil
	case Fish:
		dir, err := ConfigDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "fish", "completions", command+".fish"), nil
	case Nu:
		// Nushell auto-sources every .nu file under its vendor autoload
		// directories on startup; the user-writable one lives under the data dir.
		dir, err := DataDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "nushell", "vendor", "autoload", command+".nu"), nil
	default:
		return "", fmt.Errorf("unsupported shell %q", sh)
	}
}
