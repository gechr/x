package shell

import (
	"os"
	"path/filepath"
	"strings"
)

// xdgHome returns the value of env when it is set to an absolute path,
// falling back to filepath.Join(home, defaults...). The XDG Base Directory
// spec requires relative paths to be treated as invalid and ignored.
func xdgHome(env string, defaults ...string) (string, error) {
	if dir := os.Getenv(env); filepath.IsAbs(dir) {
		return dir, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{home}, defaults...)...), nil
}

// xdgDirs splits the colon-separated value of env, dropping empty and
// relative entries per the XDG Base Directory spec, and falls back to
// defaults when no valid entries remain.
func xdgDirs(env string, defaults ...string) []string {
	var dirs []string
	for dir := range strings.SplitSeq(os.Getenv(env), ":") {
		if filepath.IsAbs(dir) {
			dirs = append(dirs, dir)
		}
	}
	if len(dirs) == 0 {
		return defaults
	}
	return dirs
}

// XDGCacheHome returns the XDG cache home directory, defaulting to
// ~/.cache.
func XDGCacheHome() (string, error) {
	return xdgHome("XDG_CACHE_HOME", ".cache")
}

// XDGConfigDirs returns the XDG config search directories, defaulting to
// [/etc/xdg].
func XDGConfigDirs() []string {
	return xdgDirs("XDG_CONFIG_DIRS", "/etc/xdg")
}

// XDGConfigHome returns the XDG config home directory, defaulting to
// ~/.config.
func XDGConfigHome() (string, error) {
	return xdgHome("XDG_CONFIG_HOME", ".config")
}

// XDGDataDirs returns the XDG data search directories, defaulting to
// [/usr/local/share, /usr/share].
func XDGDataDirs() []string {
	return xdgDirs("XDG_DATA_DIRS", "/usr/local/share", "/usr/share")
}

// XDGDataHome returns the XDG data home directory, defaulting to
// ~/.local/share.
func XDGDataHome() (string, error) {
	return xdgHome("XDG_DATA_HOME", ".local", "share")
}

// XDGStateHome returns the XDG state home directory, defaulting to
// ~/.local/state.
func XDGStateHome() (string, error) {
	return xdgHome("XDG_STATE_HOME", ".local", "state")
}
