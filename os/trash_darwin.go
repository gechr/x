//go:build darwin

package os

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
)

// systemTrash is the command-line tool macOS ships (since macOS 15 Sequoia) to
// move a file to the Trash, recording its original location so the Finder's
// "Put Back" can restore it and trashing to the correct per-volume location for
// a file on an external disk.
const systemTrash = "/usr/bin/trash"

// trash moves path to the Trash via the system trash tool. On a macOS too old to
// ship that tool it returns an error wrapping [errors.ErrUnsupported], so a
// caller can detect the case and fall back (e.g. to [os.Remove]).
func trash(path string) error {
	if _, err := exec.LookPath(systemTrash); err != nil {
		return fmt.Errorf("%s not found, needs macOS 15+: %w", systemTrash, errors.ErrUnsupported)
	}
	out, err := exec.CommandContext(context.Background(), systemTrash, path).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to move %q to the trash: %w: %s", path, err, bytes.TrimSpace(out))
	}
	return nil
}
