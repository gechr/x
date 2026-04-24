package fs

import (
	"errors"
	"os"
)

// Exists reports whether path exists. Permission and other stat errors are
// returned to the caller.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}
