package os_test

import (
	stdos "os"
	stdpath "path/filepath"
	"strings"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestTrash(t *testing.T) {
	// Not parallel: redirects XDG_DATA_HOME so the Unix trash stays hermetic.
	dataHome := t.TempDir()
	t.Setenv("XDG_DATA_HOME", dataHome)

	file, err := stdos.CreateTemp(t.TempDir(), "xos-trash-*")
	require.NoError(t, err)
	require.NoError(t, file.Close())
	path := file.Name()
	t.Cleanup(func() { cleanupTrash(dataHome, stdpath.Base(path)) })

	require.NoError(t, xos.Trash(path))

	// The original is gone, but trashed (recoverable), not merely deleted.
	exists, err := xos.Exists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

func TestTrashSpacesAndSpecialChars(t *testing.T) {
	// Not parallel: redirects XDG_DATA_HOME so the Unix trash stays hermetic.
	dataHome := t.TempDir()
	t.Setenv("XDG_DATA_HOME", dataHome)

	// No shell is involved (exec passes argv directly; the other platforms move
	// the file via syscalls), so spaces and shell metacharacters must survive.
	base := "a file with spaces & 'quotes'.txt"
	path := stdpath.Join(t.TempDir(), base)
	require.NoError(t, stdos.WriteFile(path, []byte("x"), 0o600))
	t.Cleanup(func() { cleanupTrash(dataHome, base) })

	require.NoError(t, xos.Trash(path))

	exists, err := xos.Exists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

func TestTrashMissing(t *testing.T) {
	t.Parallel()

	err := xos.Trash(stdpath.Join(t.TempDir(), "does-not-exist"))
	require.Error(t, err)
}

// cleanupTrash removes a trashed entry so a real run does not pollute the user's
// Trash. It covers the Unix home trash (under the redirected XDG_DATA_HOME) and
// the macOS ~/.Trash, which NSFileManager uses regardless of XDG.
func cleanupTrash(dataHome, base string) {
	dirs := []string{stdpath.Join(dataHome, "Trash", "files")}
	if home, err := stdos.UserHomeDir(); err == nil {
		dirs = append(dirs, stdpath.Join(home, ".Trash"))
	}
	for _, dir := range dirs {
		entries, err := stdos.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			// The trash may disambiguate a collision (e.g. "name 2"), so match on
			// prefix, not just the exact base name.
			if strings.HasPrefix(entry.Name(), base) {
				_ = stdos.RemoveAll(stdpath.Join(dir, entry.Name()))
			}
		}
	}
}
