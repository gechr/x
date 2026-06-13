//go:build !windows

package filepath

import "strings"

// equalPath reports whether two paths are equal. Unix filesystems are
// case-sensitive, so the comparison is exact.
func equalPath(a, b string) bool {
	return a == b
}

// hasPathPrefix reports whether path begins with prefix. Unix filesystems are
// case-sensitive, so the comparison is exact.
func hasPathPrefix(path, prefix string) bool {
	return strings.HasPrefix(path, prefix)
}
