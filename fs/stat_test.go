package fs_test

import (
	ifs "io/fs"
	"os"
	"path/filepath"
	"testing"

	xfs "github.com/gechr/x/fs"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := filepath.Join(dir, "file")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))

	got, err := xfs.Exists(dir)
	require.NoError(t, err)
	require.True(t, got)

	got, err = xfs.Exists(file)
	require.NoError(t, err)
	require.True(t, got)

	got, err = xfs.Exists(filepath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, got)
}

func TestIsFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := filepath.Join(dir, "f")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))

	ok, err := xfs.IsFile(file)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xfs.IsFile(dir)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xfs.IsFile(filepath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestIsDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := filepath.Join(dir, "f")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))

	ok, err := xfs.IsDir(dir)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xfs.IsDir(file)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xfs.IsDir(filepath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestIsSymlink(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := filepath.Join(dir, "f")
	link := filepath.Join(dir, "l")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))
	require.NoError(t, os.Symlink(file, link))

	ok, err := xfs.IsSymlink(link)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = xfs.IsSymlink(file)
	require.NoError(t, err)
	require.False(t, ok)

	ok, err = xfs.IsSymlink(filepath.Join(dir, "missing"))
	require.NoError(t, err)
	require.False(t, ok)
}

func TestResolve(t *testing.T) {
	t.Parallel()

	dir, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	file := filepath.Join(dir, "f")
	link1 := filepath.Join(dir, "l1")
	link2 := filepath.Join(dir, "l2")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))
	require.NoError(t, os.Symlink(file, link1))
	require.NoError(t, os.Symlink(link1, link2))

	got, err := xfs.Resolve(link2)
	require.NoError(t, err)
	require.Equal(t, file, got)

	got, err = xfs.Resolve(link1)
	require.NoError(t, err)
	require.Equal(t, file, got)

	got, err = xfs.Resolve(file)
	require.NoError(t, err)
	require.Equal(t, file, got)

	missing := filepath.Join(dir, "missing")
	got, err = xfs.Resolve(missing)
	require.ErrorIs(t, err, ifs.ErrNotExist)
	require.Equal(t, missing, got)
}

func TestIsWritableDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	require.True(t, xfs.IsWritableDir(dir))
	require.False(t, xfs.IsWritableDir(filepath.Join(dir, "missing")))
}

func TestIsWithin(t *testing.T) {
	t.Parallel()

	require.True(t, xfs.IsWithin(".", "README.md"))
	require.True(t, xfs.IsWithin(".", "a/b.go", "c/d.go"))
	require.False(t, xfs.IsWithin("src", "lib/foo.go"))
	require.False(t, xfs.IsWithin("src"))

	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	require.NoError(t, os.Mkdir(sub, 0o755))

	require.True(t, xfs.IsWithin(dir, sub))
	require.True(t, xfs.IsWithin(dir, dir))
	require.False(t, xfs.IsWithin(sub, dir))
}
