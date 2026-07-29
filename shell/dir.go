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
func baseDir(env string, fallback func() (string, error), elem ...string) (string, error) {
	dir := os.Getenv(env)
	if !filepath.IsAbs(dir) {
		var err error
		dir, err = fallback()
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(append([]string{dir}, elem...)...), nil
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

// CacheDir returns a path within the user cache directory:
// `$XDG_CACHE_HOME` when set to an absolute path, otherwise an OS-specific
// default. Each element is joined to the directory.
func CacheDir(elem ...string) (string, error) {
	return baseDir("XDG_CACHE_HOME", cacheDirDefault, elem...)
}

// ConfigDir returns a path within the user config directory:
// `$XDG_CONFIG_HOME` when set to an absolute path, otherwise an OS-specific
// default. Each element is joined to the directory.
func ConfigDir(elem ...string) (string, error) {
	return baseDir("XDG_CONFIG_HOME", configDirDefault, elem...)
}

// DataDir returns a path within the user data directory: `$XDG_DATA_HOME` when
// set to an absolute path, otherwise an OS-specific default. Each element is
// joined to the directory.
func DataDir(elem ...string) (string, error) {
	return baseDir("XDG_DATA_HOME", dataDirDefault, elem...)
}

// StateDir returns a path within the user state directory: `$XDG_STATE_HOME`
// when set to an absolute path, otherwise an OS-specific default. Each element
// is joined to the directory.
func StateDir(elem ...string) (string, error) {
	return baseDir("XDG_STATE_HOME", stateDirDefault, elem...)
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
