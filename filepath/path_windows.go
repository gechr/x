//go:build windows

package filepath

import "strings"

// equalPath reports whether two paths are equal, compared case-insensitively
// because Windows filesystems are case-insensitive.
func equalPath(a, b string) bool {
	return strings.EqualFold(a, b)
}

// hasPathPrefix reports whether path begins with prefix, compared
// case-insensitively because Windows filesystems are case-insensitive.
func hasPathPrefix(path, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(prefix))
}
