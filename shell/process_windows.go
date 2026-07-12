//go:build windows

package shell

import (
	"path/filepath"

	"golang.org/x/sys/windows"
)

// processName returns the name of the process with the given PID using the
// Win32 process image path on Windows. The returned name is the bare base
// name without directory or `.exe` suffix, matching the darwin and linux
// implementations (so it compares against the known shell names).
func processName(pid int) string {
	h, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, uint32(pid))
	if err != nil {
		return ""
	}
	defer windows.CloseHandle(h)

	buf := make([]uint16, windows.MAX_PATH)
	size := uint32(len(buf))
	if err := windows.QueryFullProcessImageName(h, 0, &buf[0], &size); err != nil {
		return ""
	}

	return normalizeExecName(filepath.Base(windows.UTF16ToString(buf[:size])))
}
