package maps_test

import (
	"testing"

	xmaps "github.com/gechr/x/maps"
	"github.com/stretchr/testify/require"
)

func TestInvert(t *testing.T) {
	t.Parallel()

	require.Equal(t, map[int]string{1: "alpha", 2: "beta"},
		xmaps.Invert(map[string]int{"alpha": 1, "beta": 2}))

	// Empty and nil maps yield empty (non-nil) maps.
	require.Empty(t, xmaps.Invert(map[string]int{}))
	require.Empty(t, xmaps.Invert(map[string]int(nil)))

	// Duplicate values collapse to a single arbitrary key.
	inverted := xmaps.Invert(map[string]int{"a": 1, "b": 1})
	require.Len(t, inverted, 1)
	require.Contains(t, []string{"a", "b"}, inverted[1])

	// Named map types satisfy the constraint.
	type codes map[string]int
	require.Equal(t, map[int]string{200: "ok"}, xmaps.Invert(codes{"ok": 200}))
}
