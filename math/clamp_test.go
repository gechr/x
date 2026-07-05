package math_test

import (
	"math"
	"testing"

	xmath "github.com/gechr/x/math"
	"github.com/stretchr/testify/require"
)

func TestClamp(t *testing.T) {
	t.Parallel()

	require.Equal(t, 5, xmath.Clamp(5, 0, 10))
	require.Equal(t, 0, xmath.Clamp(-3, 0, 10))
	require.Equal(t, 10, xmath.Clamp(42, 0, 10))
	require.Equal(t, "b", xmath.Clamp("a", "b", "d"))
	require.InEpsilon(t, 1.5, xmath.Clamp(1.5, 1.0, 2.0), 0)
	require.InEpsilon(t, 2.0, xmath.Clamp(math.Inf(1), 1.0, 2.0), 0)
	require.InEpsilon(t, 1.0, xmath.Clamp(math.Inf(-1), 1.0, 2.0), 0)
	// NaN clamps to lo instead of propagating.
	require.InEpsilon(t, 1.0, xmath.Clamp(math.NaN(), 1.0, 2.0), 0)
}

func TestClamp01(t *testing.T) {
	t.Parallel()

	require.InEpsilon(t, 0.5, xmath.Clamp01(0.5), 0)
	require.Zero(t, xmath.Clamp01(-0.1))
	require.InEpsilon(t, 1.0, xmath.Clamp01(1.1), 0)
	require.Zero(t, xmath.Clamp01(math.NaN()))
	require.InEpsilon(t, 1.0, xmath.Clamp01(math.Inf(1)), 0)
	require.Zero(t, xmath.Clamp01(math.Inf(-1)))
}
