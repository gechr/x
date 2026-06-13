//go:build windows

package shell

import "strings"

// normalizeExecName returns the canonical form of an executable's base name.
// Windows executables are case-insensitive and carry a ".exe" suffix, so the
// name is lower-cased and the suffix stripped to match the known shell names.
func normalizeExecName(base string) string {
	return strings.TrimSuffix(strings.ToLower(base), ".exe")
}
