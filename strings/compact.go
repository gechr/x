package strings

import "strings"

// CompactLines trims lines, drops blank lines, removes duplicate lines while
// preserving first-seen order, and joins the remaining lines with sep.
func CompactLines(s, sep string) string {
	lines := strings.Split(s, "\n")
	parts := make([]string, 0, len(lines))
	seen := make(map[string]struct{}, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, ok := seen[line]; ok {
			continue
		}
		seen[line] = struct{}{}
		parts = append(parts, line)
	}
	return strings.Join(parts, sep)
}
