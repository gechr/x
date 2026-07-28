//go:build windows

package filepath_test

import (
	"testing"

	xfilepath "github.com/gechr/x/filepath"
	"github.com/stretchr/testify/require"
)

func TestIsWithin_CaseInsensitive(t *testing.T) {
	// Windows paths are case-insensitive, so containment must ignore case.
	require.True(t, xfilepath.IsWithin(`C:\foo`, `C:\FOO\bar.go`))
	require.True(t, xfilepath.IsWithin(`C:\Foo`, `c:\foo`))
	require.False(t, xfilepath.IsWithin(`C:\foo`, `C:\bar\baz.go`))
}

func TestLooksLikePath_Windows(t *testing.T) {
	t.Parallel()

	// Windows recognises backslash-rooted, UNC, and drive-letter paths on top
	// of the cross-platform "." / "/" / "~" prefixes.
	require.True(t, xfilepath.LooksLikePath(`\foo`))
	require.True(t, xfilepath.LooksLikePath(`\\server\share`))
	require.True(t, xfilepath.LooksLikePath(`C:\foo`))
	require.True(t, xfilepath.LooksLikePath(`c:/foo`))
	require.True(t, xfilepath.LooksLikePath(`Z:relative`))

	// A bare name with no marker is still not path-like.
	require.False(t, xfilepath.LooksLikePath("build"))
	require.False(t, xfilepath.LooksLikePath("owner/repo"))
}
