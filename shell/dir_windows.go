//go:build windows

package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

// Environment variables naming the standard Windows known folders.
const (
	dirAppData      = "AppData"      // roaming application data
	dirLocalAppData = "LocalAppData" // machine-local application data
	dirProgramData  = "ProgramData"  // system-wide application data
)

// Windows fallbacks map the XDG roles onto the standard Windows locations.
// They are only used when the corresponding XDG_* variable is unset or
// relative; an absolute XDG_* value is always honored (see [baseDir]). The
// locations match what [os.UserConfigDir] and [os.UserCacheDir] return, but those
// stdlib helpers ignore the XDG_* variables entirely, so we resolve them here.
//
//	ConfigDir -> %AppData%             (roaming, like os.UserConfigDir)
//	DataDir   -> %LocalAppData%
//	CacheDir  -> %LocalAppData%\cache  (%LocalAppData% is os.UserCacheDir)
//	StateDir  -> %LocalAppData%\state
func cacheDirDefault() (string, error)  { return knownFolder(dirLocalAppData, "cache") }
func configDirDefault() (string, error) { return knownFolder(dirAppData) }
func dataDirDefault() (string, error)   { return knownFolder(dirLocalAppData) }
func stateDirDefault() (string, error)  { return knownFolder(dirLocalAppData, "state") }

func configDirsDefault() []string { return programData() }
func dataDirsDefault() []string   { return programData() }

// knownFolder joins `parts` onto the directory named by the given environment
// variable (e.g. AppData, LocalAppData), erroring if it is undefined.
func knownFolder(env string, parts ...string) (string, error) {
	dir := os.Getenv(env)
	if dir == "" {
		return "", fmt.Errorf("%%%s%% is not defined", env)
	}
	return filepath.Join(append([]string{dir}, parts...)...), nil
}

// programData returns the system-wide ProgramData directory as a single-entry
// search path, or nil when it is undefined.
func programData() []string {
	if dir := os.Getenv(dirProgramData); filepath.IsAbs(dir) {
		return []string{dir}
	}
	return nil
}
