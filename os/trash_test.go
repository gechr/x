package os_test

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestTrash(t *testing.T) {
	// Not parallel: redirects XDG_DATA_HOME so the Unix trash stays hermetic.
	dataHome := t.TempDir()
	t.Setenv("XDG_DATA_HOME", dataHome)

	file, err := os.CreateTemp(t.TempDir(), "xos-trash-*")
	require.NoError(t, err)
	require.NoError(t, file.Close())
	path := file.Name()
	t.Cleanup(func() { cleanupTrash(dataHome, filepath.Base(path)) })

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
	path := filepath.Join(t.TempDir(), base)
	require.NoError(t, os.WriteFile(path, []byte("x"), 0o600))
	t.Cleanup(func() { cleanupTrash(dataHome, base) })

	require.NoError(t, xos.Trash(path))

	exists, err := xos.Exists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

func TestTrashMissing(t *testing.T) {
	t.Parallel()

	err := xos.Trash(filepath.Join(t.TempDir(), "does-not-exist"))
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestTrashConcurrentRace(t *testing.T) {
	// Not parallel: redirects XDG_DATA_HOME so the Unix trash stays hermetic.
	dataHome := t.TempDir()
	t.Setenv("XDG_DATA_HOME", dataHome)

	base := "contended.txt"
	path := filepath.Join(t.TempDir(), base)
	require.NoError(t, os.WriteFile(path, []byte("x"), 0o600))
	t.Cleanup(func() { cleanupTrash(dataHome, base) })

	// Several callers trashing the same file at once: exactly one wins, and every
	// loser sees the file already gone rather than an opaque platform failure.
	const racers = 8
	errs := make(chan error, racers)
	var start sync.WaitGroup
	start.Add(1)
	for range racers {
		go func() {
			start.Wait()
			errs <- xos.Trash(path)
		}()
	}
	start.Done()

	won := 0
	for range racers {
		if err := <-errs; err == nil {
			won++
		} else {
			require.ErrorIs(t, err, os.ErrNotExist)
		}
	}
	require.Equal(t, 1, won)

	exists, err := xos.Exists(path)
	require.NoError(t, err)
	require.False(t, exists)
}

// cleanupTrash removes a trashed entry so a real run does not pollute the user's
// Trash. It covers the Unix home trash (under the redirected XDG_DATA_HOME) and
// the macOS ~/.Trash, which NSFileManager uses regardless of XDG.
func cleanupTrash(dataHome, base string) {
	dirs := []string{filepath.Join(dataHome, "Trash", "files")}
	if home, err := os.UserHomeDir(); err == nil {
		dirs = append(dirs, filepath.Join(home, ".Trash"))
	}
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			// The trash may disambiguate a collision (e.g. "name 2"), so match on
			// prefix, not just the exact base name.
			if strings.HasPrefix(entry.Name(), base) {
				_ = os.RemoveAll(filepath.Join(dir, entry.Name()))
			}
		}
	}
}
