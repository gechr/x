//go:build windows

package shell_test

import (
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func TestDetectFromEnv_WindowsExePath(t *testing.T) {
	t.Setenv("TEST_SHELL", `C:\Program Files\PowerShell\7\pwsh.exe`)
	require.Equal(t, shell.Pwsh, shell.DetectFromEnv("TEST_SHELL"))
}

func TestDetectFromEnv_WindowsUpperCase(t *testing.T) {
	t.Setenv("TEST_SHELL", `PWSH.EXE`)
	require.Equal(t, shell.Pwsh, shell.DetectFromEnv("TEST_SHELL"))
}
