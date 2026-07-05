package time_test

import (
	"testing"
	"time"

	xtime "github.com/gechr/x/time"
	"github.com/stretchr/testify/require"
)

func TestDurations(t *testing.T) {
	t.Parallel()

	require.Equal(t, 24*time.Hour, xtime.Day)
	require.Equal(t, 7*24*time.Hour, xtime.Week)
	require.Equal(t, 365*24*time.Hour, xtime.Year)
}
