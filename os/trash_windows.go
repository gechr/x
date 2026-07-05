//go:build windows

package os

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// SHFileOperationW constants: the delete operation and the flags that send the
// file to the Recycle Bin (rather than deleting it) without UI or prompts.
const (
	foDelete          = 0x0003
	fofSilent         = 0x0004
	fofNoConfirmation = 0x0010
	fofAllowUndo      = 0x0040 // route through the Recycle Bin
	fofNoErrorUI      = 0x0400
)

// shFileOpStruct mirrors the Win32 SHFILEOPSTRUCTW. `pFrom` is a list of source
// paths terminated by a double NUL.
type shFileOpStruct struct {
	hwnd                  windows.Handle
	wFunc                 uint32
	pFrom                 *uint16
	pTo                   *uint16
	fFlags                uint16
	fAnyOperationsAborted int32
	hNameMappings         uintptr
	lpszProgressTitle     *uint16
}

var (
	shell32              = windows.NewLazySystemDLL("shell32.dll")
	procSHFileOperationW = shell32.NewProc("SHFileOperationW")
)

// trash moves `path` to the Recycle Bin via SHFileOperationW with FOF_ALLOWUNDO,
// which selects the correct per-volume bin and records the original location so
// the file can be restored. FOF_ALLOWUNDO is a request the OS may decline (e.g.
// the Recycle Bin disabled, or a path over MAX_PATH that this legacy API cannot
// address), in which case the item may be deleted outright or the call fails.
func trash(path string) error {
	from, err := doubleNullTerminated(path)
	if err != nil {
		return err
	}
	op := shFileOpStruct{
		wFunc:  foDelete,
		pFrom:  &from[0],
		fFlags: fofAllowUndo | fofNoConfirmation | fofNoErrorUI | fofSilent,
	}
	ret, _, _ := procSHFileOperationW.Call(uintptr(unsafe.Pointer(&op)))
	if ret != 0 {
		return fmt.Errorf(
			"failed to move %q to the recycle bin: SHFileOperationW returned 0x%x",
			path,
			ret,
		)
	}
	if op.fAnyOperationsAborted != 0 {
		return fmt.Errorf("moving %q to the recycle bin was aborted", path)
	}
	return nil
}

// doubleNullTerminated encodes `path` as a UTF-16 string with the extra trailing
// NUL SHFileOperationW's `pFrom` list requires.
func doubleNullTerminated(path string) ([]uint16, error) {
	encoded, err := windows.UTF16FromString(path)
	if err != nil {
		return nil, fmt.Errorf("failed to encode path: %w", err)
	}
	return append(encoded, 0), nil
}
