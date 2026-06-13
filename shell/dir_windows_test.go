//go:build windows

package shell_test

import (
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func TestConfigDir_EnvHonored(t *testing.T) {
	// An absolute XDG_CONFIG_HOME wins on Windows too.
	t.Setenv("XDG_CONFIG_HOME", `C:\custom\config`)

	got, err := shell.ConfigDir()
	require.NoError(t, err)
	require.Equal(t, `C:\custom\config`, got)
}

func TestConfigDir_FallsBackToAppData(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("AppData", `C:\Users\test\AppData\Roaming`)

	got, err := shell.ConfigDir()
	require.NoError(t, err)
	require.Equal(t, `C:\Users\test\AppData\Roaming`, got)
}

func TestDataDir_FallsBackToLocalAppData(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "")
	t.Setenv("LocalAppData", `C:\Users\test\AppData\Local`)

	got, err := shell.DataDir()
	require.NoError(t, err)
	require.Equal(t, `C:\Users\test\AppData\Local`, got)
}

func TestCacheDir_FallsBackToLocalAppDataCache(t *testing.T) {
	t.Setenv("XDG_CACHE_HOME", "")
	t.Setenv("LocalAppData", `C:\Users\test\AppData\Local`)

	got, err := shell.CacheDir()
	require.NoError(t, err)
	require.Equal(t, `C:\Users\test\AppData\Local\cache`, got)
}

func TestStateDir_FallsBackToLocalAppDataState(t *testing.T) {
	t.Setenv("XDG_STATE_HOME", "")
	t.Setenv("LocalAppData", `C:\Users\test\AppData\Local`)

	got, err := shell.StateDir()
	require.NoError(t, err)
	require.Equal(t, `C:\Users\test\AppData\Local\state`, got)
}

func TestConfigDir_ErrorsWhenAppDataUnset(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("AppData", "")

	_, err := shell.ConfigDir()
	require.Error(t, err)
}

func TestDataDirs_FallsBackToProgramData(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "")
	t.Setenv("ProgramData", `C:\ProgramData`)

	require.Equal(t, []string{`C:\ProgramData`}, shell.DataDirs())
}

func TestConfigDirs_FallsBackToProgramData(t *testing.T) {
	t.Setenv("XDG_CONFIG_DIRS", "")
	t.Setenv("ProgramData", `C:\ProgramData`)

	require.Equal(t, []string{`C:\ProgramData`}, shell.ConfigDirs())
}

func TestDataDirs_EnvHonored(t *testing.T) {
	// On Windows the search path is split on ';'.
	t.Setenv("XDG_DATA_DIRS", `C:\a;C:\b`)

	require.Equal(t, []string{`C:\a`, `C:\b`}, shell.DataDirs())
}
