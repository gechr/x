package os

import (
	"fmt"
	"os"
	"strings"

	xstrings "github.com/gechr/x/strings"
)

// ReadLines reads `path` and returns its non-empty, trimmed lines.
func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return xstrings.SplitLines(string(data)), nil
}

// WriteLines atomically writes `lines` to `path`, one per line, with a trailing
// newline.
func WriteLines(path string, lines []string, perm os.FileMode) error {
	var b strings.Builder
	for _, line := range lines {
		b.WriteString(line)
		b.WriteByte('\n')
	}
	return AtomicWrite(path, []byte(b.String()), perm)
}
