package terminal_test

import (
	"os"
	"testing"

	"github.com/gechr/x/terminal"
	"github.com/stretchr/testify/require"
)

func TestIs_Nil(t *testing.T) {
	require.False(t, terminal.Is(nil))
}

func TestIs_RegularFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "not-a-terminal")
	require.NoError(t, err)
	defer f.Close()

	require.False(t, terminal.Is(f))
}

func TestIs_Pipe(t *testing.T) {
	r, w, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()
	defer w.Close()

	require.False(t, terminal.Is(r))
	require.False(t, terminal.Is(w))
}

func TestWidth_Nil(t *testing.T) {
	require.Zero(t, terminal.Width(nil))
}

func TestWidth_RegularFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "not-a-terminal")
	require.NoError(t, err)
	defer f.Close()

	require.Zero(t, terminal.Width(f))
}

func TestWidth_Pipe(t *testing.T) {
	r, w, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()
	defer w.Close()

	require.Zero(t, terminal.Width(r))
	require.Zero(t, terminal.Width(w))
}

func TestHeight_Nil(t *testing.T) {
	require.Zero(t, terminal.Height(nil))
}

func TestHeight_RegularFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "not-a-terminal")
	require.NoError(t, err)
	defer f.Close()

	require.Zero(t, terminal.Height(f))
}

func TestSize_Nil(t *testing.T) {
	w, h := terminal.Size(nil)
	require.Zero(t, w)
	require.Zero(t, h)
}

func TestSize_Pipe(t *testing.T) {
	r, pw, err := os.Pipe()
	require.NoError(t, err)
	defer r.Close()
	defer pw.Close()

	w, h := terminal.Size(r)
	require.Zero(t, w)
	require.Zero(t, h)
}
