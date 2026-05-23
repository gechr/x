package fs_test

import (
	"os"
	"path/filepath"
	"testing"

	xfs "github.com/gechr/x/fs"
	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "lines.txt")
	require.NoError(t, os.WriteFile(path, []byte("  alpha\n\nbeta  \n\n  \ngamma\n"), 0o600))

	got, err := xfs.ReadLines(path)
	require.NoError(t, err)
	require.Equal(t, []string{"alpha", "beta", "gamma"}, got)
}

func TestReadLines_Missing(t *testing.T) {
	t.Parallel()

	_, err := xfs.ReadLines(filepath.Join(t.TempDir(), "nope"))
	require.Error(t, err)
}

func TestWriteLines(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	require.NoError(t, xfs.WriteLines(path, []string{"one", "two", "three"}, 0o600))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "one\ntwo\nthree\n", string(got))
}

func TestWriteLines_RoundTrip(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "rt.txt")
	want := []string{"x", "y", "z"}

	require.NoError(t, xfs.WriteLines(path, want, 0o600))
	got, err := xfs.ReadLines(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
