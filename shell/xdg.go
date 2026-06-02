package shell

import (
	"os"
	"path/filepath"
	"strings"
)

// XDGCacheHome returns the XDG cache home directory, defaulting to
// ~/.cache.
func XDGCacheHome() (string, error) {
	if cacheDir := os.Getenv("XDG_CACHE_HOME"); cacheDir != "" {
		return cacheDir, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache"), nil
}

// XDGConfigDirs returns the XDG config search directories, defaulting to
// [/etc/xdg].
func XDGConfigDirs() []string {
	if dirs := os.Getenv("XDG_CONFIG_DIRS"); dirs != "" {
		return strings.Split(dirs, ":")
	}
	return []string{"/etc/xdg"}
}

// XDGConfigHome returns the XDG config home directory, defaulting to
// ~/.config.
func XDGConfigHome() (string, error) {
	if configDir := os.Getenv("XDG_CONFIG_HOME"); configDir != "" {
		return configDir, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config"), nil
}

// XDGDataDirs returns the XDG data search directories, defaulting to
// [/usr/local/share, /usr/share].
func XDGDataDirs() []string {
	if dirs := os.Getenv("XDG_DATA_DIRS"); dirs != "" {
		return strings.Split(dirs, ":")
	}
	return []string{"/usr/local/share", "/usr/share"}
}

// XDGDataHome returns the XDG data home directory, defaulting to
// ~/.local/share.
func XDGDataHome() (string, error) {
	if dataDir := os.Getenv("XDG_DATA_HOME"); dataDir != "" {
		return dataDir, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share"), nil
}

// XDGStateHome returns the XDG state home directory, defaulting to
// ~/.local/state.
func XDGStateHome() (string, error) {
	if stateDir := os.Getenv("XDG_STATE_HOME"); stateDir != "" {
		return stateDir, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "state"), nil
}
