//go:build unix && !darwin

package os_test

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

// trashRoot points the home trash at a temp dir and returns its files/ and info/
// directories. The work dir holds both XDG_DATA_HOME and the subjects, so they
// share a device and the home-trash (absolute Path) branch is exercised.
func trashRoot(t *testing.T) (string, string, string) {
	t.Helper()
	work := t.TempDir()
	t.Setenv("XDG_DATA_HOME", filepath.Join(work, "xdg"))
	root := filepath.Join(work, "xdg", "Trash")
	return work, filepath.Join(root, "files"), filepath.Join(root, "info")
}

func TestTrashUnixWritesSpecRecord(t *testing.T) {
	work, filesDir, infoDir := trashRoot(t)

	base := "café déjà vu.txt"
	path := filepath.Join(work, base)
	require.NoError(t, os.WriteFile(path, []byte("x"), 0o600))

	require.NoError(t, xos.Trash(path))

	// The file moves into files/ keeping its real (unencoded) name.
	require.FileExists(t, filepath.Join(filesDir, base))

	// info/<base>.trashinfo records the URL-encoded absolute path and a date.
	info, err := os.ReadFile(filepath.Join(infoDir, base+".trashinfo"))
	require.NoError(t, err)
	lines := strings.Split(strings.TrimSpace(string(info)), "\n")
	require.Equal(t, "[Trash Info]", lines[0])
	require.Equal(t, "Path="+(&url.URL{Path: path}).EscapedPath(), lines[1])

	date, ok := strings.CutPrefix(lines[2], "DeletionDate=")
	require.True(t, ok)
	_, err = time.Parse("2006-01-02T15:04:05", date)
	require.NoError(t, err)
}

func TestTrashUnixSymlinkTrashedAsSymlink(t *testing.T) {
	work, filesDir, _ := trashRoot(t)

	target := filepath.Join(work, "target.txt")
	require.NoError(t, os.WriteFile(target, []byte("x"), 0o600))
	link := filepath.Join(work, "link.txt")
	require.NoError(t, os.Symlink(target, link))

	require.NoError(t, xos.Trash(link))

	// The symlink is trashed as itself; the target stays put.
	info, err := os.Lstat(filepath.Join(filesDir, "link.txt"))
	require.NoError(t, err)
	require.NotZero(t, info.Mode()&os.ModeSymlink)
	require.FileExists(t, target)
}

func TestTrashUnixNameCollision(t *testing.T) {
	work, filesDir, infoDir := trashRoot(t)

	for _, want := range []string{"dup.txt", "dup.txt_1"} {
		path := filepath.Join(work, "dup.txt")
		require.NoError(t, os.WriteFile(path, []byte("x"), 0o600))
		require.NoError(t, xos.Trash(path))
		require.FileExists(t, filepath.Join(filesDir, want))
		require.FileExists(t, filepath.Join(infoDir, want+".trashinfo"))
	}
}

func TestTrashUnixOrphanFilesSlotSkipped(t *testing.T) {
	work, filesDir, infoDir := trashRoot(t)

	// An orphan in files/ with no info record must not be clobbered: the next
	// trashed file of the same name takes the _1 slot instead.
	require.NoError(t, os.MkdirAll(filesDir, 0o700))
	require.NoError(t, os.WriteFile(filepath.Join(filesDir, "orphan.txt"), []byte("old"), 0o600))

	path := filepath.Join(work, "orphan.txt")
	require.NoError(t, os.WriteFile(path, []byte("new"), 0o600))
	require.NoError(t, xos.Trash(path))

	require.FileExists(t, filepath.Join(filesDir, "orphan.txt_1"))
	require.FileExists(t, filepath.Join(infoDir, "orphan.txt_1.trashinfo"))
	old, err := os.ReadFile(filepath.Join(filesDir, "orphan.txt"))
	require.NoError(t, err)
	require.Equal(t, "old", string(old)) // untouched
}
