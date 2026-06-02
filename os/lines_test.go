package os_test

import (
	stdos "os"
	stdpath "path/filepath"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := stdpath.Join(dir, "lines.txt")
	require.NoError(t, stdos.WriteFile(path, []byte("  alpha\n\nbeta  \n\n  \ngamma\n"), 0o600))

	got, err := xos.ReadLines(path)
	require.NoError(t, err)
	require.Equal(t, []string{"alpha", "beta", "gamma"}, got)
}

func TestReadLines_Missing(t *testing.T) {
	t.Parallel()

	_, err := xos.ReadLines(stdpath.Join(t.TempDir(), "nope"))
	require.Error(t, err)
}

func TestWriteLines(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := stdpath.Join(dir, "out.txt")

	require.NoError(t, xos.WriteLines(path, []string{"one", "two", "three"}, 0o600))

	got, err := stdos.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "one\ntwo\nthree\n", string(got))
}

func TestWriteLines_RoundTrip(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := stdpath.Join(dir, "rt.txt")
	want := []string{"x", "y", "z"}

	require.NoError(t, xos.WriteLines(path, want, 0o600))
	got, err := xos.ReadLines(path)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
