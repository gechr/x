//go:build !aix && !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris

package terminal

import "os"

func queryBackground(_ *os.File) (uint8, uint8, uint8, bool) {
	return 0, 0, 0, false
}
