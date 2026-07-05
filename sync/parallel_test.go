package sync_test

import (
	stdsync "sync"
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
	var mu stdsync.Mutex
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
