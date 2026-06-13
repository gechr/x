//go:build !windows

package shell

import (
	"os"
	"path/filepath"
)

func cacheDirDefault() (string, error)  { return homeDir(".cache") }
func configDirDefault() (string, error) { return homeDir(".config") }
func dataDirDefault() (string, error)   { return homeDir(".local", "share") }
func stateDirDefault() (string, error)  { return homeDir(".local", "state") }

func configDirsDefault() []string { return []string{"/etc/xdg"} }
func dataDirsDefault() []string   { return []string{"/usr/local/share", "/usr/share"} }

// homeDir joins parts onto the user's home directory.
func homeDir(parts ...string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{home}, parts...)...), nil
}
