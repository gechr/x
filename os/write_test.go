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

func TestAtomicWrite_WithNoFollowReplacesSymlink(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "target.txt")
	link := filepath.Join(dir, "link.txt")
	require.NoError(t, os.WriteFile(target, []byte("original"), 0o600))
	require.NoError(t, os.Symlink(target, link))

	require.NoError(t, xos.AtomicWrite(link, []byte("new"), 0o600, xos.WithNoFollow()))

	// The rename detaches the link rather than writing through it.
	info, err := os.Lstat(link)
	require.NoError(t, err)
	require.Zero(t, info.Mode()&os.ModeSymlink, "link should have been replaced by a regular file")

	got, err := os.ReadFile(link)
	require.NoError(t, err)
	require.Equal(t, "new", string(got))

	got, err = os.ReadFile(target)
	require.NoError(t, err)
	require.Equal(t, "original", string(got))
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

func TestEnsureDir_EnforcesPermOnExisting(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "private")
	require.NoError(t, os.Mkdir(target, 0o777))

	require.NoError(t, xos.EnsureDir(target, 0o700))
	info, err := os.Stat(target)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o700), info.Mode().Perm())
}

func TestEnsureDir_EnforcesPermOnCreation(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	parent := filepath.Join(dir, "parent")
	target := filepath.Join(parent, "child")

	require.NoError(t, xos.EnsureDir(target, 0o750))
	info, err := os.Stat(target)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o750), info.Mode().Perm())
}

func TestEnsureFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "a", "b", "config.txt")

	require.NoError(t, xos.EnsureFile(target, 0o600))
	info, err := os.Stat(target)
	require.NoError(t, err)
	require.True(t, info.Mode().IsRegular())
	requireMode(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestEnsureFile_LeavesExistingUntouched(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	target := filepath.Join(dir, "data.txt")
	require.NoError(t, os.WriteFile(target, []byte("payload"), 0o640))

	require.NoError(t, xos.EnsureFile(target, 0o600))

	got, err := os.ReadFile(target)
	require.NoError(t, err)
	require.Equal(t, "payload", string(got))
	info, err := os.Stat(target)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o640), info.Mode().Perm())
}

func TestEnsureFile_PreservesParentMode(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	parent := filepath.Join(dir, "secrets")
	require.NoError(t, os.Mkdir(parent, 0o700))

	require.NoError(t, xos.EnsureFile(filepath.Join(parent, "token"), 0o600))

	info, err := os.Stat(parent)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o700), info.Mode().Perm())
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

func TestAtomicWrite_WithNoSync(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")
	require.NoError(t, os.WriteFile(path, []byte("old"), 0o644))

	require.NoError(t, xos.AtomicWrite(path, []byte("new"), 0o600, xos.WithNoSync()))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, "new", string(got))

	info, err := os.Stat(path)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o600), info.Mode().Perm())

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "config.txt", entries[0].Name())
}

func TestAtomicWrite_WritesThroughSymlinks(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	// link1 -> links/link2 -> real/target.txt, all relative.
	require.NoError(t, os.Mkdir(filepath.Join(dir, "real"), 0o755))
	require.NoError(t, os.Mkdir(filepath.Join(dir, "links"), 0o755))
	target := filepath.Join(dir, "real", "target.txt")
	link2 := filepath.Join(dir, "links", "link2.txt")
	link1 := filepath.Join(dir, "link1.txt")
	require.NoError(t, os.WriteFile(target, []byte("original"), 0o644))
	require.NoError(t, os.Symlink(filepath.Join("..", "real", "target.txt"), link2))
	require.NoError(t, os.Symlink(filepath.Join("links", "link2.txt"), link1))

	require.NoError(t, xos.AtomicWrite(link1, []byte("new"), 0o600))

	for _, link := range []string{link1, link2} {
		info, err := os.Lstat(link)
		require.NoError(t, err)
		require.Equal(t, os.ModeSymlink, info.Mode().Type(), "%s should still be a symlink", link)
	}
	got, err := os.ReadFile(target)
	require.NoError(t, err)
	require.Equal(t, "new", string(got))
	info, err := os.Stat(target)
	require.NoError(t, err)
	requireMode(t, os.FileMode(0o600), info.Mode().Perm())

	requireDirEntries(t, filepath.Join(dir, "real"), "target.txt")
	requireDirEntries(t, filepath.Join(dir, "links"), "link2.txt")
}

func TestAtomicWrite_TempFileGoesBesideTarget(t *testing.T) {
	t.Parallel()

	// The link's directory is read-only, so the temp file must go beside the
	// target.
	dir := t.TempDir()
	links := filepath.Join(dir, "links")
	require.NoError(t, os.Mkdir(links, 0o755))
	target := filepath.Join(dir, "target.txt")
	require.NoError(t, os.WriteFile(target, []byte("original"), 0o644))
	require.NoError(t, os.Symlink(target, filepath.Join(links, "link.txt")))
	require.NoError(t, os.Chmod(links, 0o555))
	t.Cleanup(func() { _ = os.Chmod(links, 0o755) })

	require.NoError(t, xos.AtomicWrite(filepath.Join(links, "link.txt"), []byte("new"), 0o644))

	got, err := os.ReadFile(target)
	require.NoError(t, err)
	require.Equal(t, "new", string(got))
}

func TestAtomicWrite_ResolvesDotDotPhysically(t *testing.T) {
	t.Parallel()

	// a -> x/y, and x/y/link -> ../file. Through a, the kernel resolves the
	// link to x/file, where lexical cleaning would give file.
	dir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "x", "y"), 0o755))
	require.NoError(t, os.Symlink(filepath.Join("x", "y"), filepath.Join(dir, "a")))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "x", "file"), []byte("original"), 0o644))
	require.NoError(
		t,
		os.Symlink(filepath.Join("..", "file"), filepath.Join(dir, "x", "y", "link")),
	)

	require.NoError(t, xos.AtomicWrite(filepath.Join(dir, "a", "link"), []byte("new"), 0o644))

	got, err := os.ReadFile(filepath.Join(dir, "x", "file"))
	require.NoError(t, err)
	require.Equal(t, "new", string(got))
	_, err = os.Stat(filepath.Join(dir, "file"))
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestAtomicWrite_DanglingSymlinkCreatesTarget(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	link := filepath.Join(dir, "link.txt")
	require.NoError(t, os.Symlink("missing.txt", link))

	require.NoError(t, xos.AtomicWrite(link, []byte("new"), 0o644))

	got, err := os.ReadFile(filepath.Join(dir, "missing.txt"))
	require.NoError(t, err)
	require.Equal(t, "new", string(got))
	info, err := os.Lstat(link)
	require.NoError(t, err)
	require.Equal(t, os.ModeSymlink, info.Mode().Type())
}

func TestAtomicWrite_SymlinkLoop(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	a := filepath.Join(dir, "a")
	b := filepath.Join(dir, "b")
	require.NoError(t, os.Symlink("b", a))
	require.NoError(t, os.Symlink("a", b))

	err := xos.AtomicWrite(a, []byte("new"), 0o644)
	require.Error(t, err)
	requireDirEntries(t, dir, "a", "b")
}

func requireDirEntries(t *testing.T, dir string, want ...string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	got := make([]string, 0, len(entries))
	for _, e := range entries {
		got = append(got, e.Name())
	}
	require.Equal(t, want, got)
}
