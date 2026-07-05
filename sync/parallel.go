// Package sync provides concurrency helpers.
package sync

import "sync"

// Parallel runs fn(0) through fn(n-1) concurrently with at most workers in
// flight, blocking until all complete. Each call receives a distinct index,
// so a goroutine writing results[i] needs no lock; fn must otherwise be safe
// to call concurrently. workers < 1 runs one call at a time.
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
