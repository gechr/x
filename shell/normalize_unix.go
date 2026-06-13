//go:build !windows

package shell

// normalizeExecName returns the canonical form of an executable's base name.
// On Unix the base name is already canonical, so it is returned unchanged.
func normalizeExecName(base string) string {
	return base
}
