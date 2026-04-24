package strings

import "strings"

// SplitBy splits s by sep, trims whitespace from each part, and drops empty
// values.
func SplitBy(s, sep string) []string {
	raw := strings.Split(strings.TrimSpace(s), sep)
	parts := make([]string, 0, len(raw))
	for _, part := range raw {
		part = strings.TrimSpace(part)
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}
