package os_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestAtomicWrite(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")

	require.NoError(t, xos.AtomicWrite(path, []byte("hello"), 0o600))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "hello", string(got))

	info, err := os.Stat(path)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestAtomicWrite_OverwritesExisting(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")
	require.NoError(t, os.WriteFile(path, []byte("old"), 0o600))

	require.NoError(t, xos.AtomicWrite(path, []byte("new"), 0o600))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "new", string(got))
}

func TestAtomicWrite_NoTempLeftBehind(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")
	require.NoError(t, xos.AtomicWrite(path, []byte("x"), 0o600))

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "config.txt", entries[0].Name())
}

func TestEnsureDir(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "a", "b", "c")

	require.NoError(t, xos.EnsureDir(target, 0o755))
	info, err := os.Stat(target)
	require.NoError(t, err)
	require.True(t, info.IsDir())

	require.NoError(t, xos.EnsureDir(target, 0o755))
}

func TestCopyFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "src")
	dst := filepath.Join(dir, "dst")
	require.NoError(t, os.WriteFile(src, []byte("payload"), 0o640))

	require.NoError(t, xos.CopyFile(src, dst))

	got, err := os.ReadFile(dst)
	require.NoError(t, err)
	require.Equal(t, "payload", string(got))

	info, err := os.Stat(dst)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o640), info.Mode().Perm())
}

func TestCopyFile_TruncatesDestination(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "src")
	dst := filepath.Join(dir, "dst")
	require.NoError(t, os.WriteFile(src, []byte("short"), 0o600))
	require.NoError(t, os.WriteFile(dst, []byte("much longer existing content"), 0o600))

	require.NoError(t, xos.CopyFile(src, dst))

	got, err := os.ReadFile(dst)
	require.NoError(t, err)
	require.Equal(t, "short", string(got))
}

func TestCopyFile_NonRegularSource(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	srcDir := filepath.Join(dir, "srcdir")
	require.NoError(t, os.Mkdir(srcDir, 0o755))
	dst := filepath.Join(dir, "dst")

	err := xos.CopyFile(srcDir, dst)
	require.Error(t, err)

	_, err = os.Stat(dst)
	require.True(t, os.IsNotExist(err), "dst should not be created, got err=%v", err)
}

func TestCopyFile_NonRegularSourceLeavesDstUntouched(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	srcDir := filepath.Join(dir, "srcdir")
	require.NoError(t, os.Mkdir(srcDir, 0o755))
	dst := filepath.Join(dir, "dst")
	require.NoError(t, os.WriteFile(dst, []byte("existing"), 0o600))

	err := xos.CopyFile(srcDir, dst)
	require.Error(t, err)

	got, err := os.ReadFile(dst)
	require.NoError(t, err)
	require.Equal(t, "existing", string(got))
}

func TestCopyFile_MissingSource(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	err := xos.CopyFile(filepath.Join(dir, "missing"), filepath.Join(dir, "dst"))
	require.Error(t, err)
}

func TestCopyFile_PreservesModeOnExistingDst(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "src")
	dst := filepath.Join(dir, "dst")
	require.NoError(t, os.WriteFile(src, []byte("payload"), 0o640))
	require.NoError(t, os.WriteFile(dst, []byte("old"), 0o600))

	require.NoError(t, xos.CopyFile(src, dst))

	info, err := os.Stat(dst)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o640), info.Mode().Perm())
}

func requireMode(t *testing.T, want, got os.FileMode) {
	t.Helper()
	if runtime.GOOS == "windows" {
		return
	}
	require.Equal(t, want, got)
}

func TestCopyFile_SameFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "file")
	require.NoError(t, os.WriteFile(src, []byte("payload"), 0o600))

	require.NoError(t, xos.CopyFile(src, src))

	got, err := os.ReadFile(src)
	require.NoError(t, err)
	require.Equal(t, "payload", string(got))
}

func TestCopyFile_SameFileViaSymlink(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "src")
	link := filepath.Join(dir, "link")
	require.NoError(t, os.WriteFile(src, []byte("payload"), 0o600))
	require.NoError(t, os.Symlink(src, link))

	require.NoError(t, xos.CopyFile(src, link))

	got, err := os.ReadFile(src)
	require.NoError(t, err)
	require.Equal(t, "payload", string(got))
}

func TestCopyFile_SameFileViaHardlink(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	src := filepath.Join(dir, "src")
	link := filepath.Join(dir, "link")
	require.NoError(t, os.WriteFile(src, []byte("payload"), 0o600))
	require.NoError(t, os.Link(src, link))

	require.NoError(t, xos.CopyFile(src, link))

	got, err := os.ReadFile(src)
	require.NoError(t, err)
	require.Equal(t, "payload", string(got))
}
