package slices_test

import (
	"testing"

	xslices "github.com/gechr/x/slices"
	"github.com/stretchr/testify/require"
)

func TestFormat(t *testing.T) {
	t.Parallel()

	require.Equal(t,
		[]string{"Hello, Valentina!", "Hello, Ander!"},
		xslices.Format("Hello, %s!", []string{"Valentina", "Ander"}),
	)

	// Scalar arguments repeat for every element.
	require.Equal(t,
		[]string{"Salutations, Valentina!", "Salutations, Ander!"},
		xslices.Format("%s, %s!", "Salutations", []string{"Valentina", "Ander"}),
	)

	// The shortest slice determines the result length.
	require.Equal(t,
		[]string{"1: a", "2: b"},
		xslices.Format("%d: %s", []int{1, 2, 3}, []string{"a", "b"}),
	)

	// No slice arguments produces a single formatted result.
	require.Equal(t, []string{"a, b"}, xslices.Format("%s, %s", "a", "b"))

	require.Empty(t, xslices.Format("%s", []string(nil)))

	// Byte slices are treated as a single scalar, not iterated per byte.
	require.Equal(t,
		[]string{"x=hi", "y=hi"},
		xslices.Format("%s=%s", []string{"x", "y"}, []byte("hi")),
	)
	require.Equal(t, []string{"hi"}, xslices.Format("%s", []byte("hi")))
}
