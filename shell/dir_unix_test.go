//go:build !windows

package shell_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func TestCacheDir_EnvSet(t *testing.T) {
	t.Setenv("XDG_CACHE_HOME", "/custom/cache")

	got, err := shell.CacheDir()
	require.NoError(t, err)
	require.Equal(t, "/custom/cache", got)
}

func TestCacheDir_EnvUnset(t *testing.T) {
	t.Setenv("XDG_CACHE_HOME", "")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	got, err := shell.CacheDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, ".cache"), got)
}

func TestConfigDir_EnvSet(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "/custom/config")

	got, err := shell.ConfigDir()
	require.NoError(t, err)
	require.Equal(t, "/custom/config", got)
}

func TestConfigDir_EnvUnset(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	got, err := shell.ConfigDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, ".config"), got)
}

func TestDataDir_EnvSet(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/custom/data")

	got, err := shell.DataDir()
	require.NoError(t, err)
	require.Equal(t, "/custom/data", got)
}

func TestDataDir_EnvUnset(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	got, err := shell.DataDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, ".local", "share"), got)
}

func TestStateDir_EnvSet(t *testing.T) {
	t.Setenv("XDG_STATE_HOME", "/custom/state")

	got, err := shell.StateDir()
	require.NoError(t, err)
	require.Equal(t, "/custom/state", got)
}

func TestStateDir_EnvUnset(t *testing.T) {
	t.Setenv("XDG_STATE_HOME", "")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	got, err := shell.StateDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, ".local", "state"), got)
}

func TestCacheDir_EnvRelative(t *testing.T) {
	// The XDG spec requires relative paths to be treated as invalid.
	t.Setenv("XDG_CACHE_HOME", "relative/cache")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	got, err := shell.CacheDir()
	require.NoError(t, err)
	require.Equal(t, filepath.Join(home, ".cache"), got)
}

func TestDataDirs_EnvSet(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "/a:/b:/c")

	require.Equal(t, []string{"/a", "/b", "/c"}, shell.DataDirs())
}

func TestDataDirs_SkipsInvalidEntries(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "/a::relative:/b:")

	require.Equal(t, []string{"/a", "/b"}, shell.DataDirs())
}

func TestDataDirs_AllInvalid(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "relative:also/relative")

	require.Equal(t, []string{"/usr/local/share", "/usr/share"}, shell.DataDirs())
}

func TestDataDirs_EnvUnset(t *testing.T) {
	t.Setenv("XDG_DATA_DIRS", "")

	require.Equal(t, []string{"/usr/local/share", "/usr/share"}, shell.DataDirs())
}

func TestConfigDirs_EnvSet(t *testing.T) {
	t.Setenv("XDG_CONFIG_DIRS", "/a:/b")

	require.Equal(t, []string{"/a", "/b"}, shell.ConfigDirs())
}

func TestConfigDirs_EnvUnset(t *testing.T) {
	t.Setenv("XDG_CONFIG_DIRS", "")

	require.Equal(t, []string{"/etc/xdg"}, shell.ConfigDirs())
}

func TestCompletionFile(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", "/data")
	t.Setenv("XDG_CONFIG_HOME", "/config")

	tests := []struct {
		name    string
		command string
		shell   string
		want    string
	}{
		{
			name:    "bash",
			command: "myapp",
			shell:   shell.Bash,
			want:    "/data/bash-completion/completions/myapp",
		},
		{
			name:    "zsh",
			command: "myapp",
			shell:   shell.Zsh,
			want:    "/data/zsh/site-functions/_myapp",
		},
		{
			name:    "fish",
			command: "myapp",
			shell:   shell.Fish,
			want:    "/config/fish/completions/myapp.fish",
		},
		{
			name:    "nu",
			command: "myapp",
			shell:   shell.Nu,
			want:    "/data/nushell/vendor/autoload/myapp.nu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shell.CompletionFile(tt.command, tt.shell)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCompletionFile_NuEnvUnset(t *testing.T) {
	// Nushell resolves its data directory with the Rust `dirs` crate, so the
	// fallback is platform-idiomatic rather than the XDG `~/.local/share`.
	t.Setenv("XDG_DATA_HOME", "")

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	want := filepath.Join(home, ".local", "share")
	if runtime.GOOS == "darwin" {
		want = filepath.Join(home, "Library", "Application Support")
	}

	got, err := shell.CompletionFile("myapp", shell.Nu)
	require.NoError(t, err)
	require.Equal(t, filepath.Join(want, "nushell", "vendor", "autoload", "myapp.nu"), got)
}
