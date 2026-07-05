package os

import (
	"fmt"
	"os"
	"path/filepath"
)

// Trash asks the operating system to move `path` to its trash (or recycle bin)
// rather than removing it permanently like [os.Remove], so it can typically be
// recovered. The `path` is resolved to an absolute path first, so a relative
// path trashes the intended file regardless of the working directory.
//
// The mechanism is platform-specific: the system trash tool on macOS (so the
// Finder's "Put Back" works), the FreeDesktop.org trash specification on Linux
// and other Unix systems, and the shell file operation that targets the Recycle
// Bin on Windows. Recoverability is the OS's to honor, not a guarantee: an
// environment with the Recycle Bin disabled, for instance, may delete outright.
//
// Where the platform cannot trash, it returns an error wrapping
// [errors.ErrUnsupported], so a caller can detect the case and decide what to do
// (e.g. fall back to [os.Remove]). This covers a macOS older than 15 (which lacks
// the system trash tool) and a Unix file with no usable same-device trash.
func Trash(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}
	// [os.Lstat], not [os.Stat]: a symlink is trashed as itself, never
	// followed.
	if _, err := os.Lstat(abs); err != nil {
		return err
	}
	return trash(abs)
}
