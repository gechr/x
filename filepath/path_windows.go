//go:build windows

package filepath

import (
	"strings"

	xbytes "github.com/gechr/x/bytes"
)

// equalPath reports whether two paths are equal, compared case-insensitively
// because Windows filesystems are case-insensitive.
func equalPath(a, b string) bool {
	return strings.EqualFold(a, b)
}

// hasPathPrefix reports whether `path` begins with `prefix`, compared
// case-insensitively because Windows filesystems are case-insensitive.
func hasPathPrefix(path, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(prefix))
}

// looksLikePathOS reports whether `s` begins with a Windows-specific path
// marker for [LooksLikePath]: a leading backslash (including a UNC "\\") or a
// drive-letter prefix such as "C:".
func looksLikePathOS(s string) bool {
	if s[0] == '\\' {
		return true
	}
	return len(s) >= 2 && s[1] == ':' && xbytes.IsAlphaChar(s[0])
}
