package shell_test

import (
	"testing"

	"github.com/gechr/x/shell"
	"github.com/stretchr/testify/require"
)

func fakeProcess(name string) func() string {
	return func() string { return name }
}

func TestDetect_CompleteShellTakesPrecedence(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "zsh")
	t.Setenv(shell.EnvShell, "/bin/bash")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("fish")))

	got := shell.Detect()
	if got != "zsh" {
		t.Errorf("Detect() = %q, want %q", got, "zsh")
	}
}

func TestDetect_CompleteShellPathReturnsBase(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "/usr/local/bin/fish")
	t.Setenv(shell.EnvShell, "/bin/bash")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("")))

	got := shell.Detect()
	if got != "fish" {
		t.Errorf("Detect() = %q, want %q", got, "fish")
	}
}

func TestDetect_ProcessTakesPrecedenceOverShellEnv(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "")
	t.Setenv(shell.EnvShell, "/bin/bash")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("zsh")))

	got := shell.Detect()
	if got != "zsh" {
		t.Errorf("Detect() = %q, want %q", got, "zsh")
	}
}

func TestDetect_FallsBackToShellEnv(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "")
	t.Setenv(shell.EnvShell, "/bin/zsh")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("")))

	got := shell.Detect()
	if got != "zsh" {
		t.Errorf("Detect() = %q, want %q", got, "zsh")
	}
}

func TestDetect_EmptyWhenNothingSet(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "")
	t.Setenv(shell.EnvShell, "")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("")))

	got := shell.Detect()
	if got != "" {
		t.Errorf("Detect() = %q, want empty", got)
	}
}

func TestDetect_UnknownProcessIgnored(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "")
	t.Setenv(shell.EnvShell, "/bin/fish")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("node")))

	got := shell.Detect()
	if got != "fish" {
		t.Errorf("Detect() = %q, want %q", got, "fish")
	}
}

func TestDetectFromEnv(t *testing.T) {
	t.Setenv("TEST_SHELL", "/usr/bin/fish")

	got := shell.DetectFromEnv("TEST_SHELL")
	if got != "fish" {
		t.Errorf("DetectFromEnv() = %q, want %q", got, "fish")
	}
}

func TestDetectFromEnv_Empty(t *testing.T) {
	t.Setenv("TEST_SHELL", "")

	got := shell.DetectFromEnv("TEST_SHELL")
	if got != "" {
		t.Errorf("DetectFromEnv() = %q, want empty", got)
	}
}

func TestDetectFromEnv_Unset(t *testing.T) {
	got := shell.DetectFromEnv("NONEXISTENT_SHELL_VAR_FOR_TEST")
	if got != "" {
		t.Errorf("DetectFromEnv() = %q, want empty", got)
	}
}

func TestDetectFromEnv_UnknownIgnored(t *testing.T) {
	t.Setenv("TEST_SHELL", "/usr/bin/node")

	got := shell.DetectFromEnv("TEST_SHELL")
	if got != "" {
		t.Errorf("DetectFromEnv() = %q, want empty", got)
	}
}

func TestDetect_UnknownShellEnvIgnored(t *testing.T) {
	t.Setenv("COMPLETE_SHELL", "")
	t.Setenv(shell.EnvShell, "/usr/bin/node")
	t.Cleanup(shell.SetParentProcessName(fakeProcess("")))

	got := shell.Detect()
	if got != "" {
		t.Errorf("Detect() = %q, want empty", got)
	}
}

func TestDetectFromProcess(t *testing.T) {
	t.Cleanup(shell.SetParentProcessName(fakeProcess("fish")))

	got := shell.DetectFromProcess()
	if got != "fish" {
		t.Errorf("DetectFromProcess() = %q, want %q", got, "fish")
	}
}

func TestDetectFromProcess_UnknownIgnored(t *testing.T) {
	t.Cleanup(shell.SetParentProcessName(fakeProcess("node")))

	got := shell.DetectFromProcess()
	if got != "" {
		t.Errorf("DetectFromProcess() = %q, want empty", got)
	}
}

func TestDetectFromProcess_Empty(t *testing.T) {
	t.Cleanup(shell.SetParentProcessName(fakeProcess("")))

	got := shell.DetectFromProcess()
	if got != "" {
		t.Errorf("DetectFromProcess() = %q, want empty", got)
	}
}

func TestCompletionFile_UnsupportedShell(t *testing.T) {
	_, err := shell.CompletionFile("myapp", "unsupported")
	require.Error(t, err)
	require.EqualError(t, err, `unsupported shell "unsupported"`)
}

func TestKnown(t *testing.T) {
	known := shell.Known()
	require.NotEmpty(t, known)
	require.Contains(t, known, "bash")
	require.Contains(t, known, "zsh")
	require.Contains(t, known, "fish")

	// Verify mutation safety: modifying the returned slice should not affect future calls.
	known[0] = "mutated"
	fresh := shell.Known()
	require.NotEqual(t, "mutated", fresh[0])
}

func TestIsKnown(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"bash", true},
		{"nonexistent", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, shell.IsKnown(tt.name))
		})
	}
}
