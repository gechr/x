package bytes

import "bytes"

// EnsureTrailingNewline trims any trailing newlines from `s` and appends exactly
// one, so the result always ends in a single `\n`. An empty slice becomes
// `\n`. The returned slice never aliases `s`.
func EnsureTrailingNewline(s []byte) []byte {
	trimmed := bytes.TrimRight(s, "\n")
	// Build on a fresh backing array: `trimmed` shares `s`, so appending onto
	// it directly could clobber `s` or its spare capacity.
	return append(append(make([]byte, 0, len(trimmed)+1), trimmed...), '\n')
}
