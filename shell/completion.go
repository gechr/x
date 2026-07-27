package shell

import (
	"fmt"
	"path/filepath"
)

// CompletionFile returns the standard completion file path for the given
// `command` and shell.
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
		// directories on startup; the user-writable one lives under its own data
		// directory. That honors `$XDG_DATA_HOME` on every platform, but its
		// fallback is the platform-idiomatic location rather than the XDG one
		// [DataDir] uses, so it is resolved separately.
		dir, err := baseDir("XDG_DATA_HOME", nuDataDirDefault)
		if err != nil {
			return "", err
		}
		return filepath.Join(dir, "nushell", "vendor", "autoload", command+".nu"), nil
	default:
		return "", fmt.Errorf("unsupported shell %q", sh)
	}
}
