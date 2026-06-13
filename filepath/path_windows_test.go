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
