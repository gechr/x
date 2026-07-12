// Package encoding provides helpers for structured document formats such as
// JSON and YAML: building field paths for diagnostics and lookup.
package encoding

import (
	"strconv"
	"strings"
	"unicode"
)

// segmentKind discriminates how a [Path] segment addresses into a document.
type segmentKind uint8

const (
	segmentChild segmentKind = iota
	segmentIndex
	segmentKey
	segmentWildcard
)

// Path is an immutable field path into a structured document (JSON, YAML,
// ...), built segment by segment for diagnostics and lookup:
//
//	NewPath("items").Index(0).Child("foo", "bar").Wildcard() // items[0].foo.bar[*]
//
// Each method allocates a new Path referencing its receiver as parent, so
// multiple children can safely branch off a shared prefix - handy in
// recursive document walkers.
//
// A nil `*Path` is the empty path: it renders as `""` and addresses the
// document itself in lookups. All methods are nil-safe.
type Path struct {
	parent *Path
	kind   segmentKind
	name   string
	index  int
}

// NewPath returns a [Path] rooted at `name`, extended with `moreNames` as
// nested children.
func NewPath(name string, moreNames ...string) *Path {
	var root *Path
	return root.Child(name, moreNames...)
}

// Child returns `p` extended with `name` (and `moreNames`) as nested field
// segments. Names render in dot notation (`.name`), or bracket-quoted
// (`["a.b"]`) when they contain characters other than letters, digits, `_`,
// and `-`.
func (p *Path) Child(name string, moreNames ...string) *Path {
	child := &Path{parent: p, kind: segmentChild, name: name}
	for _, n := range moreNames {
		child = &Path{parent: child, kind: segmentChild, name: n}
	}
	return child
}

// Index returns `p` extended with an array index segment, rendered as `[3]`.
func (p *Path) Index(index int) *Path {
	return &Path{parent: p, kind: segmentIndex, index: index}
}

// Key returns `p` extended with an explicit key segment, always rendered
// bracket-quoted (`["name"]`) even when `key` would be valid in dot notation.
// Lookup treats it exactly like [Path.Child].
func (p *Path) Key(key string) *Path {
	return &Path{parent: p, kind: segmentKey, name: key}
}

// Wildcard returns `p` extended with a `[*]` segment matching every element
// of an array or every value of a map. [Path.LookupAll] fans out at wildcard
// segments; [Path.Lookup] cannot resolve them.
func (p *Path) Wildcard() *Path {
	return &Path{parent: p, kind: segmentWildcard}
}

// Render returns the path in dot/bracket notation, e.g. `items[0].foo.bar[*]`.
// Names that cannot appear in dot notation are bracket-quoted: `spec["a.b"]`.
// Pass [WithRoot] to prefix a JSONPath-style root marker: `$.items[0]`. The
// output is a human-readable diagnostic notation, not strict RFC 9535
// JSONPath (bare names are broader, and quoting follows Go syntax).
func (p *Path) Render(opts ...RenderOption) string {
	var cfg renderConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	var sb strings.Builder
	if cfg.root {
		sb.WriteRune(cfg.marker)
	}
	for i, seg := range p.segments() {
		seg.render(&sb, i > 0 || cfg.root)
	}
	return sb.String()
}

// String implements [fmt.Stringer]; it is [Path.Render] with default options.
func (p *Path) String() string {
	return p.Render()
}

// render writes the segment to `sb`. `dot` says whether a dot-notation name
// needs a leading `.` - false only for the first segment of an unrooted path.
func (p *Path) render(sb *strings.Builder, dot bool) {
	switch p.kind {
	case segmentChild:
		if !isBareName(p.name) {
			sb.WriteString("[" + strconv.Quote(p.name) + "]")
			return
		}
		if dot {
			sb.WriteByte('.')
		}
		sb.WriteString(p.name)
	case segmentIndex:
		sb.WriteString("[" + strconv.Itoa(p.index) + "]")
	case segmentKey:
		sb.WriteString("[" + strconv.Quote(p.name) + "]")
	case segmentWildcard:
		sb.WriteString("[*]")
	}
}

// segments returns the chain of segments from the root to `p`.
func (p *Path) segments() []*Path {
	depth := 0
	for q := p; q != nil; q = q.parent {
		depth++
	}
	segs := make([]*Path, depth)
	for q := p; q != nil; q = q.parent {
		depth--
		segs[depth] = q
	}
	return segs
}

// isBareName reports whether `name` can render unquoted in dot notation: one
// or more letters, digits, `_`, or `-`, with combining marks allowed after
// the first rune (so decomposed accents, as in NFD `café`, stay bare).
func isBareName(name string) bool {
	if name == "" {
		return false
	}
	for i, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			continue
		}
		if i > 0 && unicode.IsMark(r) {
			continue
		}
		return false
	}
	return true
}
