package time_test

import (
	"testing"
	stdtime "time"

	xtime "github.com/gechr/x/time"
	"github.com/stretchr/testify/require"
)

func TestDurations(t *testing.T) {
	t.Parallel()

	require.Equal(t, 24*stdtime.Hour, xtime.Day)
	require.Equal(t, 7*24*stdtime.Hour, xtime.Week)
	require.Equal(t, 365*24*stdtime.Hour, xtime.Year)
}
