package palette_test

import (
	"fmt"
	"image/color"
	"sync"
	"testing"

	"github.com/gechr/x/palette"
	"github.com/stretchr/testify/require"
)

func TestAssignerStablePerKey(t *testing.T) {
	a := palette.NewAssigner(palette.DefaultDark()...)

	first := a.Assign("alpha")
	require.NotNil(t, first)
	require.Equal(t, first, a.Assign("alpha"))
}

func TestNewAssignerDefaultsToAuto(t *testing.T) {
	a := palette.NewAssigner()

	require.NotEmpty(t, a.Palette())
	require.NotNil(t, a.Assign("alpha"))
}

func TestNewAssignerUsesGivenColors(t *testing.T) {
	pal := palette.DefaultDark()
	require.Equal(t, pal, palette.NewAssigner(pal...).Palette())
}

func TestAssignerFirstKeyGetsHashColor(t *testing.T) {
	pal := palette.DefaultDark()
	a := palette.NewAssigner(pal...)

	// With nothing assigned yet, the first key lands on its stable hash color -
	// the same one [Palette.Color] would pick.
	require.Equal(t, pal.Color("alpha"), a.Assign("alpha"))
}

func TestAssignerNoDuplicatesUntilExhausted(t *testing.T) {
	pal := palette.Palette{color.Black, color.White, color.Opaque}
	a := palette.NewAssigner(pal...)

	seen := map[color.Color]bool{}
	for _, key := range []string{"alpha", "beta", "gamma"} {
		seen[a.Assign(key)] = true
	}
	// All three distinct keys map to distinct colors despite hash collisions.
	require.Len(t, seen, 3)

	// Palette exhausted: a fourth key must reuse an existing color.
	require.Contains(t, []color.Color(pal), a.Assign("delta"))
}

func TestAssignerConcurrent(t *testing.T) {
	a := palette.NewAssigner(palette.DefaultDark()...)

	var wg sync.WaitGroup
	for i := range 64 {
		wg.Go(func() {
			a.Assign(fmt.Sprintf("key-%d", i%8))
		})
	}
	wg.Wait()

	// Every one of the 8 distinct keys resolves to a stable color.
	for i := range 8 {
		key := fmt.Sprintf("key-%d", i)
		c := a.Assign(key)
		require.NotNil(t, c)
		require.Equal(t, c, a.Assign(key))
	}
}
