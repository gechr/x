package shell

import "strings"

// Quote returns a shell-escaped version of s. The returned value can safely be
// used as one token in a POSIX shell command line.
func Quote(s string) string {
	if s == "" {
		return "''"
	}
	if strings.IndexFunc(s, needsQuote) == -1 {
		return s
	}
	return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
}

func needsQuote(r rune) bool {
	return (r < 'a' || r > 'z') &&
		(r < 'A' || r > 'Z') &&
		(r < '0' || r > '9') &&
		r != '_' &&
		!strings.ContainsRune("@%+=:,./-", r)
}
