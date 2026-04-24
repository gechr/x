package strings

// SplitLines splits s into non-empty trimmed lines.
func SplitLines(s string) []string {
	return SplitBy(s, "\n")
}
