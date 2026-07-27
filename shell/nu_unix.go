//go:build !windows && !darwin

package shell

// nuDataDirDefault matches [DataDir] on Unix, where Nushell's data directory
// follows the XDG default of `~/.local/share`.
func nuDataDirDefault() (string, error) { return dataDirDefault() }
