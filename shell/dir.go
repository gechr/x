package shell

import (
	"os"
	"path/filepath"
	"strings"
)

// baseDir returns `env` when it is set to an absolute path, otherwise the
// result of `fallback`. Per the XDG Base Directory spec, a relative path in the
// environment variable is invalid and ignored. The XDG_* variables are
// honored on every platform; only the `fallback` is OS-specific.
func baseDir(env string, fallback func() (string, error)) (string, error) {
	if dir := os.Getenv(env); filepath.IsAbs(dir) {
		return dir, nil
	}
	return fallback()
}

// searchDirs splits the value of `env` on the OS list separator, dropping empty
// and relative entries per the XDG Base Directory spec, and falls back to the
// OS-specific `fallback` when no valid entries remain.
func searchDirs(env string, fallback []string) []string {
	var dirs []string
	for dir := range strings.SplitSeq(os.Getenv(env), string(os.PathListSeparator)) {
		if filepath.IsAbs(dir) {
			dirs = append(dirs, dir)
		}
	}
	if len(dirs) == 0 {
		return fallback
	}
	return dirs
}

// CacheDir returns the user cache directory: `$XDG_CACHE_HOME` when set to an
// absolute path, otherwise an OS-specific default.
func CacheDir() (string, error) {
	return baseDir("XDG_CACHE_HOME", cacheDirDefault)
}

// ConfigDir returns the user config directory: `$XDG_CONFIG_HOME` when set to an
// absolute path, otherwise an OS-specific default.
func ConfigDir() (string, error) {
	return baseDir("XDG_CONFIG_HOME", configDirDefault)
}

// DataDir returns the user data directory: `$XDG_DATA_HOME` when set to an
// absolute path, otherwise an OS-specific default.
func DataDir() (string, error) {
	return baseDir("XDG_DATA_HOME", dataDirDefault)
}

// StateDir returns the user state directory: `$XDG_STATE_HOME` when set to an
// absolute path, otherwise an OS-specific default.
func StateDir() (string, error) {
	return baseDir("XDG_STATE_HOME", stateDirDefault)
}

// ConfigDirs returns the ordered, read-only config search directories:
// `$XDG_CONFIG_DIRS` when it has absolute entries, otherwise OS-specific
// defaults. These are searched after [ConfigDir], so a user's config overrides
// the system defaults.
func ConfigDirs() []string {
	return searchDirs("XDG_CONFIG_DIRS", configDirsDefault())
}

// DataDirs returns the ordered, read-only data search directories:
// `$XDG_DATA_DIRS` when it has absolute entries, otherwise OS-specific defaults.
// These are searched after [DataDir], so a user's data overrides the system
// defaults.
func DataDirs() []string {
	return searchDirs("XDG_DATA_DIRS", dataDirsDefault())
}
