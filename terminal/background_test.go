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

func TestIsDarkColor(t *testing.T) {
	tests := map[string]struct {
		red, green, blue uint8
		dark             bool
	}{
		"black":    {0, 0, 0, true},
		"white":    {255, 255, 255, false},
		"gray 117": {117, 117, 117, true},
		"gray 118": {118, 118, 118, false},
		// HSL lightness puts every pure primary at 0.5, the old threshold, so
		// it cannot tell dark blue from bright green.
		"pure red":        {255, 0, 0, false},
		"pure green":      {0, 255, 0, false},
		"pure blue":       {0, 0, 255, true},
		"solarized dark":  {0x00, 0x2b, 0x36, true},
		"solarized light": {0xfd, 0xf6, 0xe3, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, test.dark, terminal.IsDarkColor(test.red, test.green, test.blue))
		})
	}
}
