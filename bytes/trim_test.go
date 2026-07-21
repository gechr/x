package bytes_test

import (
	"testing"

	xbytes "github.com/gechr/x/bytes"
	"github.com/stretchr/testify/require"
)

func TestTrimPrefixes(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		"example.com/pkg",
		string(
			xbytes.TrimPrefixes(
				[]byte("https://example.com/pkg"),
				[]byte("https://"),
				[]byte("http://"),
			),
		),
	)
	require.Equal(
		t,
		"example.com/pkg",
		string(
			xbytes.TrimPrefixes(
				[]byte("http://example.com/pkg"),
				[]byte("https://"),
				[]byte("http://"),
			),
		),
	)
	require.Equal(
		t,
		"example.com/pkg",
		string(
			xbytes.TrimPrefixes([]byte("example.com/pkg"), []byte("https://"), []byte("http://")),
		),
	)
	require.Equal(
		t,
		"b-a-s",
		string(xbytes.TrimPrefixes([]byte("a-b-a-s"), []byte("a-"), []byte("b-"))),
	)
	require.Equal(t, "foo", string(xbytes.TrimPrefixes([]byte("foo"))))
	require.Empty(t, xbytes.TrimPrefixes([]byte(""), []byte("x")))
}

func TestTrimSuffixes(t *testing.T) {
	t.Parallel()

	require.Equal(
		t,
		"archive",
		string(xbytes.TrimSuffixes([]byte("archive.tar.gz"), []byte(".tar.gz"), []byte(".tgz"))),
	)
	require.Equal(
		t,
		"archive",
		string(xbytes.TrimSuffixes([]byte("archive.tgz"), []byte(".tar.gz"), []byte(".tgz"))),
	)
	require.Equal(
		t,
		"archive",
		string(xbytes.TrimSuffixes([]byte("archive"), []byte(".tar.gz"), []byte(".tgz"))),
	)
	require.Equal(
		t,
		"s-a-b",
		string(xbytes.TrimSuffixes([]byte("s-a-b-a"), []byte("-a"), []byte("-b"))),
	)
	require.Equal(t, "foo", string(xbytes.TrimSuffixes([]byte("foo"))))
	require.Empty(t, xbytes.TrimSuffixes([]byte(""), []byte("x")))
}
