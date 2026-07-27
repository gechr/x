//go:build windows

package shell

// nuDataDirDefault returns the roaming `%AppData%` directory. Nushell resolves
// its data directory with the Rust `dirs` crate, which maps to the roaming
// known folder on Windows, not the machine-local one [DataDir] uses.
func nuDataDirDefault() (string, error) { return knownFolder(dirAppData) }
