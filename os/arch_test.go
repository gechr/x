package os_test

import (
	"runtime"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestIsWasm(t *testing.T) {
	t.Parallel()

	require.Equal(t, runtime.GOARCH == xos.ArchWASM, xos.IsWasm())
}
