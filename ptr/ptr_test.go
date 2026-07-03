package ptr_test

import (
	"testing"

	"github.com/gechr/x/ptr"
	"github.com/stretchr/testify/require"
)

func TestDeref(t *testing.T) {
	t.Parallel()

	require.True(t, ptr.Deref(new(true)))
	require.False(t, ptr.Deref(new(false)))
	require.False(t, ptr.Deref((*bool)(nil)))

	require.Equal(t, 42, ptr.Deref(new(42)))
	require.Equal(t, 0, ptr.Deref((*int)(nil)))
	require.Empty(t, ptr.Deref((*string)(nil)))
}
