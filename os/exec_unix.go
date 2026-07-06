//go:build unix

package os

import "golang.org/x/sys/unix"

// executable reports whether the current process may execute `path`, following
// symlinks. It uses `access(2)` with `X_OK` so the effective owner/group/other bit
// is chosen based on the process's real uid/gid rather than assuming the owner
// bit applies.
func executable(path string) bool {
	return unix.Access(path, unix.X_OK) == nil
}
