package human

import (
	"os"
	"path/filepath"
	"strings"
)

// ContractHome replaces the user's home directory prefix with ~.
func ContractHome(path string) string {
	if home, err := os.UserHomeDir(); err == nil {
		rel, err := filepath.Rel(home, path)
		if err != nil {
			return path
		}
		if rel == "." {
			return "~"
		}
		if rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return "~/" + filepath.ToSlash(rel)
		}
	}
	return path
}
