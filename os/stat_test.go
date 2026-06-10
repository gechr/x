package os_test

import (
	stdos "os"
	stdpath "path/filepath"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := stdpath.Join(dir, "file")
	require.NoError(t, stdos.WriteFile(file, []byte("x"), 0o600))

	got, err := xos.Exists(dir)
	require.NoError(t, err)
	require.True(t, got)

	got, err = xos.Exists(file)
	require.NoError(t, err)
	require.True(t, got)

	got, err = xos.Exists(stdpath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, got)

	// A path under a regular file does not exist (ENOTDIR, not an error).
	got, err = xos.Exists(stdpath.Join(file, "sub"))
	require.NoError(t, err)
	require.False(t, got)
}

func TestIsFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := stdpath.Join(dir, "f")
	require.NoError(t, stdos.WriteFile(file, []byte("x"), 0o600))

	ok, err := xos.IsFile(file)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xos.IsFile(dir)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xos.IsFile(stdpath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestIsDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := stdpath.Join(dir, "f")
	require.NoError(t, stdos.WriteFile(file, []byte("x"), 0o600))

	ok, err := xos.IsDir(dir)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xos.IsDir(file)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xos.IsDir(stdpath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestIsSymlink(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := stdpath.Join(dir, "f")
	link := stdpath.Join(dir, "l")
	require.NoError(t, stdos.WriteFile(file, []byte("x"), 0o600))
	require.NoError(t, stdos.Symlink(file, link))

	ok, err := xos.IsSymlink(link)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xos.IsSymlink(file)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xos.IsSymlink(stdpath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xos.IsSymlink(stdpath.Join(file, "sub"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestIsWritableDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.True(t, xos.IsWritableDir(dir))
	require.False(t, xos.IsWritableDir(stdpath.Join(dir, "missing")))
}
