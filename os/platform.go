package os

import "runtime"

// Platform constants are the recognized [runtime.GOOS] values. Go exposes GOOS
// only as a string, so these name the tokens to avoid scattering string
// literals across build-time comparisons.
const (
	PlatformAIX       = "aix"
	PlatformAndroid   = "android"
	PlatformDarwin    = "darwin"
	PlatformDragonfly = "dragonfly"
	PlatformFreeBSD   = "freebsd"
	PlatformIllumos   = "illumos"
	PlatformIOS       = "ios"
	PlatformJS        = "js"
	PlatformLinux     = "linux"
	PlatformNetBSD    = "netbsd"
	PlatformOpenBSD   = "openbsd"
	PlatformPlan9     = "plan9"
	PlatformSolaris   = "solaris"
	PlatformWASIP1    = "wasip1"
	PlatformWindows   = "windows"
)

// IsWindows reports whether the program is running on Windows.
func IsWindows() bool {
	return runtime.GOOS == PlatformWindows
}

// IsDarwin reports whether the program is running on macOS. It matches Go's
// `darwin` GOOS token; iOS is a separate GOOS (see [IsIOS]) and reports false.
func IsDarwin() bool {
	return runtime.GOOS == PlatformDarwin
}

// IsLinux reports whether the program is running on Linux.
func IsLinux() bool {
	return runtime.GOOS == PlatformLinux
}

// IsAndroid reports whether the program is running on Android. Android is its
// own GOOS but is also Unix-like, so it additionally satisfies [IsUnix].
func IsAndroid() bool {
	return runtime.GOOS == PlatformAndroid
}

// IsIOS reports whether the program is running on iOS. iOS is its own GOOS -
// distinct from macOS (see [IsDarwin]) - but is also Unix-like, so it
// additionally satisfies [IsUnix].
func IsIOS() bool {
	return runtime.GOOS == PlatformIOS
}

// IsBSD reports whether the program is running on a BSD-family OS: FreeBSD,
// NetBSD, OpenBSD, or DragonFly BSD. Go has no `bsd` build constraint, so this
// is a fixed GOOS set. macOS is deliberately excluded even though Darwin is
// BSD-derived; use [IsDarwin] for it. Every OS reported here is also Unix-like,
// so it satisfies [IsUnix].
func IsBSD() bool {
	switch runtime.GOOS {
	case PlatformFreeBSD, PlatformNetBSD, PlatformOpenBSD, PlatformDragonfly:
		return true
	default:
		return false
	}
}

// IsUnix reports whether the program is running on a Unix-like OS. It mirrors
// Go's `unix` build constraint - which spans Linux, macOS, the BSDs, and mobile
// GOOSes, among others - rather than any single GOOS, so [IsLinux], [IsDarwin],
// [IsAndroid], [IsIOS], and [IsBSD] all imply IsUnix.
func IsUnix() bool {
	return isUnix
}
