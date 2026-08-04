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
	backgroundQueryTimeout = 10 * time.Millisecond
	backgroundQuery        = "\x1b]11;?\x1b\\"
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

	deadline := time.Now().Add(backgroundQueryTimeout)
	response := make([]byte, 0, maxBackgroundResponse)
	buffer := make([]byte, maxBackgroundResponse)
	for len(response) < maxBackgroundResponse {
		remaining := time.Until(deadline)
		if remaining <= 0 {
			return 0, 0, 0, false
		}

		timeout := max(int(remaining.Milliseconds()), 1)
		//nolint:gosec // File descriptors fit PollFd.Fd.
		pollFDs := []unix.PollFd{{Fd: int32(fd), Events: unix.POLLIN}}
		ready, err := unix.Poll(pollFDs, timeout)
		if errors.Is(err, unix.EINTR) {
			continue
		}
		if err != nil || ready == 0 {
			return 0, 0, 0, false
		}

		count, err := f.Read(buffer[:min(len(buffer), maxBackgroundResponse-len(response))])
		if err != nil || count == 0 {
			return 0, 0, 0, false
		}
		response = append(response, buffer[:count]...)
		if red, green, blue, ok := parseBackgroundResponse(string(response)); ok {
			return red, green, blue, true
		}
	}

	return 0, 0, 0, false
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
