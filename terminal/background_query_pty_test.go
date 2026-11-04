//go:build darwin || linux

package terminal

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

func TestQueryBackground_Replies(t *testing.T) {
	const (
		color = "\x1b]11;rgb:0000/0000/0000\x1b\\"
		da1   = "\x1b[?62;4c"
	)
	tests := map[string]struct {
		response   string
		fragmented bool
		valid      bool
	}{
		"normal":     {response: color + da1, valid: true},
		"fragmented": {response: color + da1, fragmented: true, valid: true},
		"leading noise": {
			response: strings.Repeat("x", maxBackgroundResponse*2) + color + da1,
			valid:    true,
		},
		"oversized OSC": {
			response: "\x1b]11;" + strings.Repeat("x", maxBackgroundResponse*2) + "\a" + da1,
		},
		"oversized DA1": {
			response: color + "\x1b[?" + strings.Repeat("1;", maxBackgroundResponse) + "c",
			valid:    true,
		},
		"aborted CSI":      {response: "\x1b[?\x1b[0mc" + color + da1, valid: true},
		"other CSI final":  {response: "\x1b[?25hc" + color + da1, valid: true},
		"incomplete color": {response: "\x1b]11;rgb:0000/0000/0000" + da1},
		"unsupported OSC":  {response: da1},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			master, slave := backgroundPTY(t)
			state, err := term.GetState(int(slave.Fd()))
			require.NoError(t, err)

			done := make(chan backgroundResult, 1)
			go func() {
				dark, ok := detectBackground(slave)
				done <- backgroundResult{dark: dark, ok: ok}
			}()

			require.Equal(t, backgroundQuery, readPTYBytes(t, master, len(backgroundQuery)))
			input := test.response + "USER"
			if test.fragmented {
				for i := range len(input) {
					_, err = master.WriteString(input[i : i+1])
					require.NoError(t, err)
				}
			} else {
				_, err = master.WriteString(input)
				require.NoError(t, err)
			}

			select {
			case result := <-done:
				require.Equal(t, test.valid, result.ok)
				if result.ok {
					require.True(t, result.dark)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("background query did not finish")
			}
			restored, err := term.GetState(int(slave.Fd()))
			require.NoError(t, err)
			require.Equal(t, state, restored)

			_, err = term.MakeRaw(int(slave.Fd()))
			require.NoError(t, err)
			t.Cleanup(func() { require.NoError(t, term.Restore(int(slave.Fd()), state)) })
			require.Equal(t, "USER", readPTYBytes(t, slave, len("USER")))
		})
	}
}

func TestQueryBackground_NoResponse(t *testing.T) {
	master, slave := backgroundPTY(t)
	state, err := term.GetState(int(slave.Fd()))
	require.NoError(t, err)
	done := make(chan backgroundResult, 1)
	go func() {
		dark, ok := detectBackground(slave)
		done <- backgroundResult{dark: dark, ok: ok}
	}()
	require.Equal(t, backgroundQuery, readPTYBytes(t, master, len(backgroundQuery)))
	select {
	case result := <-done:
		require.False(t, result.ok)
	case <-time.After(2 * time.Second):
		t.Fatal("background query did not time out")
	}
	restored, err := term.GetState(int(slave.Fd()))
	require.NoError(t, err)
	require.Equal(t, state, restored)
}

func backgroundPTY(t *testing.T) (*os.File, *os.File) {
	t.Helper()
	master, path := openPTYPair(t)
	slave, err := os.OpenFile(path, os.O_RDWR|unix.O_NOCTTY, 0)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, slave.Close()) })
	return master, slave
}

func readPTYBytes(t *testing.T, f *os.File, length int) string {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	buffer := make([]byte, length)
	for count := 0; count < length; {
		remaining := time.Until(deadline)
		require.Positive(t, remaining, "timed out reading pseudo-terminal")
		fds := []unix.PollFd{{Fd: int32(f.Fd()), Events: unix.POLLIN}}
		ready, err := unix.Poll(fds, max(int(remaining.Milliseconds()), 1))
		require.NoError(t, err)
		require.Positive(t, ready, "timed out reading pseudo-terminal")
		n, err := f.Read(buffer[count:])
		require.NoError(t, err)
		require.Positive(t, n)
		count += n
	}
	return string(buffer)
}
