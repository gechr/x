package ansi_test

import (
	"os"
	"testing"

	"github.com/gechr/x/ansi"
	"github.com/stretchr/testify/require"
)

func TestNew_NoOptions(t *testing.T) {
	w := ansi.New()
	require.False(t, w.Terminal())
}

func TestNew_WithTerminalTrue(t *testing.T) {
	w := ansi.New(ansi.WithTerminal(true))
	require.True(t, w.Terminal())
}

func TestNew_WithTerminalFalse(t *testing.T) {
	w := ansi.New(ansi.WithTerminal(false))
	require.False(t, w.Terminal())
}

func TestNever(t *testing.T) {
	w := ansi.Never()
	require.False(t, w.Terminal())
}

func TestForce(t *testing.T) {
	w := ansi.Force()
	require.True(t, w.Terminal())
}

func TestAuto_DefaultNonTerminal(t *testing.T) {
	// In test environments, stdout is not a terminal.
	w := ansi.Auto()
	require.False(t, w.Terminal())
}

func TestAuto_ExplicitNonTerminal(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "non-terminal")
	require.NoError(t, err)
	defer f.Close()

	w := ansi.Auto(f)
	require.False(t, w.Terminal())
}

func TestAuto_MultipleFiles(t *testing.T) {
	f1, err := os.CreateTemp(t.TempDir(), "a")
	require.NoError(t, err)
	defer f1.Close()

	f2, err := os.CreateTemp(t.TempDir(), "b")
	require.NoError(t, err)
	defer f2.Close()

	w := ansi.Auto(f1, f2)
	require.False(t, w.Terminal())
}

func TestAuto_NilFile(t *testing.T) {
	w := ansi.Auto(nil)
	require.False(t, w.Terminal())
}

func TestAuto_NilAmongFiles(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "valid")
	require.NoError(t, err)
	defer f.Close()

	w := ansi.Auto(f, nil)
	require.False(t, w.Terminal())
}
