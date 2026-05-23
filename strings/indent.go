package strings

import "strings"

// Indent prefixes every non-blank line of s with prefix. Blank and
// whitespace-only lines are normalized to empty.
//
//	Indent("foo\nbar", "  ")      // "  foo\n  bar"
//	Indent("foo\n\nbar", "> ")    // "> foo\n\n> bar"
//	Indent("foo\n   \nbar", "> ") // "> foo\n\n> bar"
func Indent(s, prefix string) string {
	if s == "" {
		return s
	}
	lines := strings.Split(s, "\n")
	var b strings.Builder
	b.Grow(len(s) + len(prefix)*len(lines))
	for i, line := range lines {
		if i > 0 {
			b.WriteByte('\n')
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		b.WriteString(prefix)
		b.WriteString(line)
	}
	return b.String()
}

// Dedent strips the longest common leading-whitespace prefix from non-empty
// lines. Whitespace-only lines are normalized to empty (Python textwrap.dedent).
//
//	Dedent("    foo\n      bar\n    baz") // "foo\n  bar\nbaz"
func Dedent(s string) string {
	if s == "" {
		return s
	}
	lines := strings.Split(s, "\n")

	prefix := ""
	first := true
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		lead := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
		if first {
			prefix = lead
			first = false
			continue
		}
		prefix = commonPrefix(prefix, lead)
		if prefix == "" {
			break
		}
	}

	if prefix == "" {
		for i, line := range lines {
			if strings.TrimSpace(line) == "" {
				lines[i] = ""
			}
		}
		return strings.Join(lines, "\n")
	}

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lines[i] = ""
			continue
		}
		lines[i] = strings.TrimPrefix(line, prefix)
	}
	return strings.Join(lines, "\n")
}

func commonPrefix(a, b string) string {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
	return a[:i]
}
