//go:build !unix

package os

// syncDir is a no-op outside Go's `unix` build constraint. Nothing there exposes
// a directory sync a process can request: [os.File.Sync] maps to
// `FlushFileBuffers` on Windows, which rejects a directory handle, and Plan 9
// and WASM have no equivalent. Directory-entry durability is left to whatever
// the filesystem provides - NTFS journals metadata operations, so a rename is
// recovered whole or not at all - rather than to the caller.
func syncDir(string) error {
	return nil
}
