package bytes

import "bytes"

// Indent prefixes every non-blank line of `s` with `prefix`. Blank and
// whitespace-only lines are normalized to empty, and CRLF line endings to LF.
//
//	Indent([]byte("foo\nbar"), []byte("  "))      // "  foo\n  bar"
//	Indent([]byte("foo\n\nbar"), []byte("> "))    // "> foo\n\n> bar"
//	Indent([]byte("foo\n   \nbar"), []byte("> ")) // "> foo\n\n> bar"
func Indent(s, prefix []byte) []byte {
	if len(s) == 0 {
		return s
	}
	lines := SplitLinesRaw(s)
	var b bytes.Buffer
	b.Grow(len(s) + len(prefix)*len(lines))
	for i, line := range lines {
		if i > 0 {
			b.WriteByte('\n')
		}
		if IsBlank(line) {
			continue
		}
		b.Write(prefix)
		b.Write(line)
	}
	return b.Bytes()
}

// Dedent strips the longest common leading-whitespace prefix from non-empty
// lines. Whitespace-only lines are normalized to empty (Python textwrap.dedent)
// and CRLF line endings to LF.
//
//	Dedent([]byte("    foo\n      bar\n    baz")) // "foo\n  bar\nbaz"
func Dedent(s []byte) []byte {
	if len(s) == 0 {
		return s
	}
	lines := SplitLinesRaw(s)

	var prefix []byte
	first := true
	for _, line := range lines {
		if IsBlank(line) {
			continue
		}
		lead := line[:len(line)-len(bytes.TrimLeft(line, " \t"))]
		if first {
			prefix = lead
			first = false
			continue
		}
		prefix = commonPrefix(prefix, lead)
		if len(prefix) == 0 {
			break
		}
	}

	if len(prefix) == 0 {
		for i, line := range lines {
			if IsBlank(line) {
				lines[i] = nil
			}
		}
		return bytes.Join(lines, []byte("\n"))
	}

	for i, line := range lines {
		if IsBlank(line) {
			lines[i] = nil
			continue
		}
		lines[i] = bytes.TrimPrefix(line, prefix)
	}
	return bytes.Join(lines, []byte("\n"))
}

func commonPrefix(a, b []byte) []byte {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
	return a[:i]
}
