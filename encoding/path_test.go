package encoding_test

import (
	"strconv"
	"testing"

	xencoding "github.com/gechr/x/encoding"
	"github.com/stretchr/testify/require"
)

func TestPathString(t *testing.T) {
	t.Parallel()

	p := xencoding.NewPath("items")
	require.Equal(t, "items", p.String())

	p = p.Index(0)
	require.Equal(t, "items[0]", p.String())

	p = p.Child("foo", "bar")
	require.Equal(t, "items[0].foo.bar", p.String())

	p = p.Wildcard()
	require.Equal(t, "items[0].foo.bar[*]", p.String())

	p = p.Key("name")
	require.Equal(t, `items[0].foo.bar[*]["name"]`, p.String())

	// Multi-name constructor.
	require.Equal(t, "spec.template.spec", xencoding.NewPath("spec", "template", "spec").String())

	// Negative indices render verbatim.
	require.Equal(t, "items[-1]", xencoding.NewPath("items").Index(-1).String())
}

func TestPathQuoting(t *testing.T) {
	t.Parallel()

	// Letters, digits, '_', and '-' stay in dot notation.
	require.Equal(
		t,
		"spec.some-key.a_b.v2",
		xencoding.NewPath("spec").Child("some-key", "a_b", "v2").String(),
	)

	// Unicode letters stay bare, including decomposed accents (NFD).
	require.Equal(t, "spec.caf\u00e9", xencoding.NewPath("spec").Child("caf\u00e9").String())
	require.Equal(t, "spec.cafe\u0301", xencoding.NewPath("spec").Child("cafe\u0301").String())

	// A combining mark cannot start a bare name.
	require.Equal(
		t,
		"spec["+strconv.Quote("\u0301x")+"]",
		xencoding.NewPath("spec").Child("\u0301x").String(),
	)

	// Control characters render with Go escapes.
	require.Equal(t, `spec["a\x01b"]`, xencoding.NewPath("spec").Child("a\x01b").String())

	// Reserved characters force bracket-quoting.
	require.Equal(
		t,
		`labels["kubernetes.io/hostname"]`,
		xencoding.NewPath("labels").Child("kubernetes.io/hostname").String(),
	)
	require.Equal(t, `spec["a b"]`, xencoding.NewPath("spec").Child("a b").String())
	require.Equal(t, `spec[""]`, xencoding.NewPath("spec").Child("").String())
	require.Equal(t, `spec["say \"hi\""]`, xencoding.NewPath("spec").Child(`say "hi"`).String())

	// A quoted segment can start the path.
	require.Equal(t, `["a.b"].name`, xencoding.NewPath("a.b").Child("name").String())

	// Key always quotes, even bare names.
	require.Equal(t, `spec["name"]`, xencoding.NewPath("spec").Key("name").String())
}

func TestPathBranching(t *testing.T) {
	t.Parallel()

	// Extending a path never mutates it: branches share the prefix.
	base := xencoding.NewPath("items").Index(0)
	left := base.Child("foo")
	right := base.Child("bar").Wildcard()

	require.Equal(t, "items[0]", base.String())
	require.Equal(t, "items[0].foo", left.String())
	require.Equal(t, "items[0].bar[*]", right.String())
}

func TestPathNil(t *testing.T) {
	t.Parallel()

	// A nil *Path is the empty path: it renders as "" and addresses the
	// document itself.
	var p *xencoding.Path
	require.Empty(t, p.String())
	require.Equal(t, "$", p.Render(xencoding.WithRoot()))

	v, ok := p.Lookup("doc")
	require.True(t, ok)
	require.Equal(t, "doc", v)
	require.Equal(t, []any{"doc"}, p.LookupAll("doc"))
}

func TestPathRender(t *testing.T) {
	t.Parallel()

	p := xencoding.NewPath("items").Index(0)
	require.Equal(t, "items[0]", p.Render())
	require.Equal(t, "$.items[0]", p.Render(xencoding.WithRoot()))
	require.Equal(t, "@.items[0]", p.Render(xencoding.WithRoot('@')))

	// The root marker abuts a bracket-quoted first segment.
	require.Equal(t, `$["a.b"]`, xencoding.NewPath("a.b").Render(xencoding.WithRoot()))
}
