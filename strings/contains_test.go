package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestContainsAll(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.ContainsAll("foo bar baz", "foo", "bar", "baz"))
	require.True(t, xstrings.ContainsAll("foo", "foo"))
	require.True(t, xstrings.ContainsAll("anything"))
	require.False(t, xstrings.ContainsAll("foo bar", "foo", "baz"))
	require.False(t, xstrings.ContainsAll("", "x"))
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	require.True(t, xstrings.ContainsAny("foo bar baz", "baz", "qux"))
	require.True(t, xstrings.ContainsAny("foo", "foo", "bar"))
	require.False(t, xstrings.ContainsAny("foo bar"))
	require.False(t, xstrings.ContainsAny("foo bar", "baz", "qux"))
	require.False(t, xstrings.ContainsAny("", "x"))
}
