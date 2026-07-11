package os_test

import (
	"runtime"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestPlatformPredicates(t *testing.T) {
	t.Parallel()

	// Each single-GOOS predicate must agree with runtime.GOOS via its constant.
	require.Equal(t, runtime.GOOS == xos.PlatformWindows, xos.IsWindows())
	require.Equal(t, runtime.GOOS == xos.PlatformDarwin, xos.IsDarwin())
	require.Equal(t, runtime.GOOS == xos.PlatformLinux, xos.IsLinux())
	require.Equal(t, runtime.GOOS == xos.PlatformAndroid, xos.IsAndroid())
	require.Equal(t, runtime.GOOS == xos.PlatformIOS, xos.IsIOS())
}

func TestIsBSD(t *testing.T) {
	t.Parallel()

	bsd := map[string]bool{
		xos.PlatformFreeBSD:   true,
		xos.PlatformNetBSD:    true,
		xos.PlatformOpenBSD:   true,
		xos.PlatformDragonfly: true,
	}
	require.Equal(t, bsd[runtime.GOOS], xos.IsBSD())

	// Darwin is BSD-derived but deliberately excluded from IsBSD.
	require.False(t, xos.IsDarwin() && xos.IsBSD())
}

func TestIsUnix(t *testing.T) {
	t.Parallel()

	// Every family predicate implies the Unix-like grouping.
	if xos.IsLinux() || xos.IsDarwin() || xos.IsAndroid() || xos.IsIOS() || xos.IsBSD() {
		require.True(t, xos.IsUnix())
	}

	// Windows is never Unix-like.
	require.False(t, xos.IsWindows() && xos.IsUnix())
}
