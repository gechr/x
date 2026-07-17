# sync

```go
import "github.com/gechr/x/sync"
```

Package `sync` provides concurrency helpers.

## Index

- [func Parallel(workers, n int, fn func(i int))](<#Parallel>)
- [func ParallelErr(workers, n int, fn func(i int) error) error](<#ParallelErr>)

<a name="Parallel"></a>

## func [Parallel](<https://github.com/gechr/x/blob/main/sync/parallel.go#L14>)

```go
func Parallel(workers, n int, fn func(i int))
```

**Parallel** runs `fn(0)` through `fn(n-1)` concurrently with at most `workers` in flight, blocking until all complete. Each call receives a distinct index, so a goroutine writing `results[i]` needs no lock; `fn` must otherwise be safe to call concurrently. `workers` \< 1 runs one call at a time.

<details><summary><b>Example</b></summary>

```go
// Each call receives a distinct index, so writing results[i]
// needs no lock.
results := make([]int, 5)
xsync.Parallel(3, len(results), func(i int) {
    results[i] = i * i
})
fmt.Println(results)
```

Output:

```text
[0 1 4 9 16]
```

</details>

<details><summary><b>Example (SingleWorker)</b></summary>

A single worker runs the calls one at a time, in index order.

```go
xsync.Parallel(1, 3, func(i int) {
    fmt.Println("call", i)
})
```

Output:

```text
call 0
call 1
call 2
```

</details>

<a name="ParallelErr"></a>

## func [ParallelErr](<https://github.com/gechr/x/blob/main/sync/parallel.go#L37>)

```go
func ParallelErr(workers, n int, fn func(i int) error) error
```

**ParallelErr** is [Parallel](<#Parallel>) for tasks that can fail. All `n` calls run regardless of failures - one task's error does not cancel the others. It returns nil if every call succeeded, otherwise an error joining each failure wrapped with its task index; [errors.Is](<https://pkg.go.dev/errors#Is>) and [errors.As](<https://pkg.go.dev/errors#As>) reach every cause through the join.

<details><summary><b>Example</b></summary>

```go
items := []int{2, 7, 4, 9}
err := xsync.ParallelErr(2, len(items), func(i int) error {
    if items[i]%2 != 0 {
        return fmt.Errorf("odd value %d", items[i])
    }
    return nil
})
fmt.Println(err)
```

Output:

```text
task 1: odd value 7
task 3: odd value 9
```

</details>
