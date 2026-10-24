package terminal

import (
	"os"
	"strconv"
	"testing"

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
	require.NoError(t, unix.IoctlSetPointerInt(fd, unix.TIOCSPTLCK, 0))
	number, err := unix.IoctlGetUint32(fd, unix.TIOCGPTN)
	require.NoError(t, err)
	return "/dev/pts/" + strconv.FormatUint(uint64(number), 10)
}
