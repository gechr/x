package terminal

import (
	"os"
	"syscall"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// openPTY opens a pseudo-terminal and returns the path of its subsidiary side.
// The controlling side stays open until the test ends.
func openPTY(t *testing.T) string {
	t.Helper()

	ptmx, err := os.OpenFile("/dev/ptmx", os.O_RDWR|unix.O_NOCTTY, 0)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, ptmx.Close()) })

	fd := int(ptmx.Fd())
	require.NoError(t, unix.IoctlSetInt(fd, unix.TIOCPTYGRANT, 0))
	require.NoError(t, unix.IoctlSetInt(fd, unix.TIOCPTYUNLK, 0))

	// x/sys has no wrapper for TIOCPTYGNAME, which fills a caller buffer.
	var name [128]byte
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		ptmx.Fd(),
		unix.TIOCPTYGNAME,
		uintptr(unsafe.Pointer(&name[0])),
	)
	if errno != 0 {
		require.NoError(t, errno)
	}
	return unix.ByteSliceToString(name[:])
}
