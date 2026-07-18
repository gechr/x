package strings

import (
	"strings"

	xslices "github.com/gechr/x/slices"
)

// CompactLines trims lines, drops blank lines, removes duplicate lines while
// preserving first-seen order, and joins the remaining lines with `sep`.
func CompactLines(s, sep string) string {
	return strings.Join(xslices.Unique(SplitLines(s)), sep)
}
