//go:build !unix && !windows

package os

import (
	"errors"
	"fmt"
	"runtime"
)

// trash returns an error wrapping [errors.ErrUnsupported]: platforms outside
// Unix and Windows (e.g. Plan 9, WASM) have no trash facility, so a caller can
// detect the case and fall back (e.g. to [os.Remove]).
func trash(path string) error {
	return fmt.Errorf("cannot trash %q on %s: %w", path, runtime.GOOS, errors.ErrUnsupported)
}
