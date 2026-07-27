//go:build darwin

package shell

// nuDataDirDefault returns `~/Library/Application Support` on macOS. Nushell
// resolves its data directory with the Rust `dirs` crate, which is
// platform-idiomatic rather than XDG-conformant, so it never falls back to
// `~/.local/share` here.
func nuDataDirDefault() (string, error) { return homeDir("Library", "Application Support") }
