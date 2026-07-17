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

func ExampleParallelErr() {
	items := []int{2, 7, 4, 9}
	err := xsync.ParallelErr(2, len(items), func(i int) error {
		if items[i]%2 != 0 {
			return fmt.Errorf("odd value %d", items[i])
		}
		return nil
	})
	fmt.Println(err)
	// Output:
	// task 1: odd value 7
	// task 3: odd value 9
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
