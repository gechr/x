// Package sync provides concurrency helpers.
package sync

import (
	"errors"
	"fmt"
	"sync"
)

// Parallel runs `fn(0)` through `fn(n-1)` concurrently with at most `workers` in
// flight, blocking until all complete. Each call receives a distinct index,
// so a goroutine writing `results[i]` needs no lock; `fn` must otherwise be safe
// to call concurrently. `workers` < 1 runs one call at a time.
func Parallel(workers, n int, fn func(i int)) {
	if workers < 1 {
		workers = 1
	}
	slots := make(chan struct{}, workers)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		slots <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-slots }()
			fn(i)
		}()
	}
	wg.Wait()
}

// ParallelErr is [Parallel] for tasks that can fail. All `n` calls run
// regardless of failures - one task's error does not cancel the others. It
// returns nil if every call succeeded, otherwise an error joining each failure
// wrapped with its task index; [errors.Is] and [errors.As] reach every cause
// through the join.
func ParallelErr(workers, n int, fn func(i int) error) error {
	errs := make([]error, n)
	Parallel(workers, n, func(i int) {
		if err := fn(i); err != nil {
			errs[i] = fmt.Errorf("task %d: %w", i, err)
		}
	})
	return errors.Join(errs...)
}
