package sync_test

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	xsync "github.com/gechr/x/sync"
	"github.com/stretchr/testify/require"
)

func TestParallel(t *testing.T) {
	t.Parallel()

	const n = 100
	results := make([]int, n)
	xsync.Parallel(8, n, func(i int) {
		results[i] = i * 2
	})
	for i, got := range results {
		require.Equal(t, i*2, got)
	}
}

func TestParallelBoundsWorkers(t *testing.T) {
	t.Parallel()

	const workers = 4
	var inFlight, peak atomic.Int64
	var mu sync.Mutex
	xsync.Parallel(workers, 64, func(int) {
		cur := inFlight.Add(1)
		defer inFlight.Add(-1)
		mu.Lock()
		if cur > peak.Load() {
			peak.Store(cur)
		}
		mu.Unlock()
	})
	require.LessOrEqual(t, peak.Load(), int64(workers))
	require.Positive(t, peak.Load())
}

func TestParallelZeroWorkersRunsSerially(t *testing.T) {
	t.Parallel()

	var inFlight, peak atomic.Int64
	xsync.Parallel(0, 16, func(int) {
		cur := inFlight.Add(1)
		defer inFlight.Add(-1)
		if cur > peak.Load() {
			peak.Store(cur)
		}
	})
	require.Equal(t, int64(1), peak.Load())
}

func TestParallelZeroCalls(t *testing.T) {
	t.Parallel()

	called := false
	xsync.Parallel(4, 0, func(int) { called = true })
	require.False(t, called)
}

func TestParallelErr(t *testing.T) {
	t.Parallel()

	errOdd := errors.New("odd")
	var calls atomic.Int64
	err := xsync.ParallelErr(4, 10, func(i int) error {
		calls.Add(1)
		if i%2 == 1 {
			return errOdd
		}
		return nil
	})
	require.Equal(t, int64(10), calls.Load(), "a failure must not cancel other tasks")
	require.ErrorIs(t, err, errOdd)
	for _, i := range []int{1, 3, 5, 7, 9} {
		require.ErrorContains(t, err, fmt.Sprintf("task %d: odd", i))
	}
}

func TestParallelErrAllSucceed(t *testing.T) {
	t.Parallel()

	err := xsync.ParallelErr(4, 10, func(int) error { return nil })
	require.NoError(t, err)
}

func TestParallelErrZeroCalls(t *testing.T) {
	t.Parallel()

	err := xsync.ParallelErr(4, 0, func(int) error { return errors.New("unreachable") })
	require.NoError(t, err)
}
