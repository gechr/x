//go:build unix

package os

import (
	"fmt"
	"os"
)

// syncDir flushes `dir`'s own entries to stable storage. Syncing a file persists
// its contents; the name pointing at those contents lives in the parent
// directory, so a rename is only durable once the directory itself is synced.
// Unix permits `fsync(2)` on a read-only directory descriptor, which is the one
// portable way to ask for that.
func syncDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return fmt.Errorf("failed to open directory: %w", err)
	}
	defer d.Close()

	if err := d.Sync(); err != nil {
		return fmt.Errorf("failed to sync directory: %w", err)
	}
	return nil
}
