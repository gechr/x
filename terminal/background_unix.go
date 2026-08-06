//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package terminal

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

const (
	// backgroundQueryTimeout bounds the wait for a terminal that answers
	// neither query. It is a fault backstop rather than a latency budget:
	// terminals that speak DA1 - effectively all of them - end the read as
	// soon as the reply lands, so this is only reached behind a pty with no
	// emulator on the other end, such as a CI runner or script(1). It is
	// generous enough for a DA1 round trip over a high-latency link, since
	// giving up early is the failure that leaks escape sequences.
	backgroundQueryTimeout = 500 * time.Millisecond
	// backgroundQuery asks for the background colour (OSC 11) and follows it
	// with a Primary Device Attributes request (DA1). Terminals answer queries
	// in order, so the DA1 reply marks the end of the batch: seeing it with no
	// OSC 11 reply before it means the terminal does not support the query,
	// which a timeout alone cannot distinguish from a terminal that is merely
	// slow. Reading through to DA1 also guarantees no reply is left behind in
	// the input buffer, where the shell would echo it once the process exits.
	backgroundQuery = "\x1b]11;?\x1b\\\x1b[0c"
	// deviceAttributesPrefix and deviceAttributesFinal bracket a DA1 reply,
	// which takes the form CSI ? Ps ; ... c.
	deviceAttributesPrefix = "\x1b[?"
	deviceAttributesFinal  = 'c'
	backgroundResponse     = "\x1b]11;rgb:"
	maxBackgroundResponse  = 128
	rgbComponentCount      = 3
	bitsPerHexDigit        = 4
)

func queryBackground(f *os.File) (uint8, uint8, uint8, bool) {
	fd := int(f.Fd())
	state, err := term.MakeRaw(fd)
	if err != nil {
		return 0, 0, 0, false
	}
	defer term.Restore(fd, state) //nolint:errcheck // Best-effort restoration after probing.

	if _, err := f.WriteString(backgroundQuery); err != nil {
		return 0, 0, 0, false
	}

	return parseBackgroundResponse(readQueryResponse(f, fd))
}

// readQueryResponse reads the terminal's replies to backgroundQuery, stopping
// at the DA1 reply that trails them. It reads a byte at a time so it consumes
// exactly the replies and no more, leaving any input the user typed ahead of
// them in the buffer rather than swallowing it in a bulk read.
func readQueryResponse(f *os.File, fd int) string {
	deadline := time.Now().Add(backgroundQueryTimeout)
	var response strings.Builder
	buffer := make([]byte, 1)

	for response.Len() < maxBackgroundResponse {
		remaining := time.Until(deadline)
		if remaining <= 0 {
			break
		}

		timeout := max(int(remaining.Milliseconds()), 1)
		//nolint:gosec // File descriptors fit PollFd.Fd.
		pollFDs := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
		ready, err := unix.Poll(pollFDs, timeout)
		if errors.Is(err, unix.EINTR) {
			continue
		}
		if err != nil || ready == 0 {
			break
		}

		count, err := f.Read(buffer)
		if err != nil || count == 0 {
			break
		}
		response.Write(buffer[:count])
		if deviceAttributesComplete(response.String()) {
			break
		}
	}

	return response.String()
}

// deviceAttributesComplete reports whether response ends with a complete DA1
// reply. The search starts after the CSI ? introducer so that hex digits in an
// OSC 11 reply cannot be mistaken for the final byte.
func deviceAttributesComplete(response string) bool {
	_, after, ok := strings.Cut(response, deviceAttributesPrefix)
	if !ok {
		return false
	}
	return strings.ContainsRune(after, deviceAttributesFinal)
}

func parseBackgroundResponse(response string) (uint8, uint8, uint8, bool) {
	_, after, ok := strings.Cut(response, backgroundResponse)
	if !ok {
		return 0, 0, 0, false
	}

	color := after
	end := strings.IndexAny(color, "\a\x1b")
	if end < 0 {
		return 0, 0, 0, false
	}

	components := strings.Split(color[:end], "/")
	if len(components) != rgbComponentCount {
		return 0, 0, 0, false
	}

	red, ok := parseColorComponent(components[0])
	if !ok {
		return 0, 0, 0, false
	}
	green, ok := parseColorComponent(components[1])
	if !ok {
		return 0, 0, 0, false
	}
	blue, ok := parseColorComponent(components[2])
	if !ok {
		return 0, 0, 0, false
	}

	return red, green, blue, true
}

func parseColorComponent(component string) (uint8, bool) {
	if len(component) < 1 || len(component) > 4 {
		return 0, false
	}

	value, err := strconv.ParseUint(component, 16, 16)
	if err != nil {
		return 0, false
	}

	maximum := uint64(1)<<(bitsPerHexDigit*len(component)) - 1
	normalized := value * colorChannelMaximum / maximum
	//nolint:gosec // The normalized result is at most 255.
	return uint8(normalized), true
}
