//go:build darwin || linux

package terminal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

// A stream reopened on the terminal in one direction passes isatty but cannot
// both send the query and read the reply, so it must not be chosen.
func TestTerminalFile_SkipsOneWayTerminal(t *testing.T) {
	path := openPTY(t)
	open := func(flag int) *os.File {
		f, err := os.OpenFile(path, flag|unix.O_NOCTTY, 0)
		require.NoError(t, err)
		t.Cleanup(func() { require.NoError(t, f.Close()) })
		return f
	}

	readOnly := open(os.O_RDONLY)
	writeOnly := open(os.O_WRONLY)
	readWrite := open(os.O_RDWR)

	require.True(t, Is(readOnly))
	require.True(t, Is(writeOnly))
	require.Same(t, readWrite, terminalFile(readOnly, writeOnly, readWrite))
	require.Nil(t, terminalFile(readOnly, writeOnly))
}
