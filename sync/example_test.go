package sync_test

import (
	"fmt"

	xsync "github.com/gechr/x/sync"
)

func ExampleParallel() {
	// Each call receives a distinct index, so writing results[i]
	// needs no lock.
	results := make([]int, 5)
	xsync.Parallel(3, len(results), func(i int) {
		results[i] = i * i
	})
	fmt.Println(results)
	// Output:
	// [0 1 4 9 16]
}

// A single worker runs the calls one at a time, in index order.
func ExampleParallel_singleWorker() {
	xsync.Parallel(1, 3, func(i int) {
		fmt.Println("call", i)
	})
	// Output:
	// call 0
	// call 1
	// call 2
}
