package errors_test

import (
	"errors"
	"fmt"
	"testing"

	xerrors "github.com/gechr/x/errors"
	"github.com/stretchr/testify/require"
)

func TestIsAny(t *testing.T) {
	t.Parallel()

	errA := errors.New("a")
	errB := errors.New("b")
	errC := errors.New("c")

	require.True(t, xerrors.IsAny(errA, errA))
	require.True(t, xerrors.IsAny(fmt.Errorf("wrapped: %w", errB), errA, errB))
	require.True(t, xerrors.IsAny(errors.Join(errA, errB), errB, errC))
	require.True(t, xerrors.IsAny(nil, errA, nil))

	require.False(t, xerrors.IsAny(errA, errB, errC))
	require.False(t, xerrors.IsAny(errA))
	require.False(t, xerrors.IsAny(errA, nil))
}
