package terminal_test

import (
	"os"
	"testing"

	"github.com/gechr/x/terminal"
	"github.com/stretchr/testify/require"
)

func TestDetectBackground_Nil(t *testing.T) {
	dark, ok := terminal.DetectBackground(nil)
	require.False(t, ok)
	require.False(t, dark)
}

func TestDetectBackground_RegularFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "not-a-terminal")
	require.NoError(t, err)
	defer f.Close()

	dark, ok := terminal.DetectBackground(f)
	require.False(t, ok)
	require.False(t, dark)
}

func TestDetectBackground_Pipe(t *testing.T) {
	r, w, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()
	defer w.Close()

	for _, f := range []*os.File{r, w} {
		dark, ok := terminal.DetectBackground(f)
		require.False(t, ok)
		require.False(t, dark)
	}
}

// Under `go test` no standard stream is a terminal, so both queries report
// ok=false rather than guessing.
func TestIsDark_NoTerminal(t *testing.T) {
	dark, ok := terminal.IsDark()
	require.False(t, ok)
	require.False(t, dark)

	light, ok := terminal.IsLight()
	require.False(t, ok)
	require.False(t, light)
}
