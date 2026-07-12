package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestTrimPrefixes(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		"example.com/pkg",
		xstrings.TrimPrefixes("https://example.com/pkg", "https://", "http://"),
	)
	require.Equal(
		t,
		"example.com/pkg",
		xstrings.TrimPrefixes("http://example.com/pkg", "https://", "http://"),
	)
	require.Equal(
		t,
		"example.com/pkg",
		xstrings.TrimPrefixes("example.com/pkg", "https://", "http://"),
	)
	require.Equal(t, "b-a-s", xstrings.TrimPrefixes("a-b-a-s", "a-", "b-"))
	require.Equal(t, "foo", xstrings.TrimPrefixes("foo"))
	require.Empty(t, xstrings.TrimPrefixes("", "x"))
}

func TestTrimSuffixes(t *testing.T) {
	t.Parallel()

	require.Equal(t, "archive", xstrings.TrimSuffixes("archive.tar.gz", ".tar.gz", ".tgz"))
	require.Equal(t, "archive", xstrings.TrimSuffixes("archive.tgz", ".tar.gz", ".tgz"))
	require.Equal(t, "archive", xstrings.TrimSuffixes("archive", ".tar.gz", ".tgz"))
	require.Equal(t, "s-a-b", xstrings.TrimSuffixes("s-a-b-a", "-a", "-b"))
	require.Equal(t, "foo", xstrings.TrimSuffixes("foo"))
	require.Empty(t, xstrings.TrimSuffixes("", "x"))
}
