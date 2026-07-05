# sync

```go
import "github.com/gechr/x/sync"
```

Package `sync` provides concurrency helpers.

## Index

- [func Parallel(workers, n int, fn func(i int))](<#Parallel>)

<a name="Parallel"></a>

## func [Parallel](<https://github.com/gechr/x/blob/main/sync/parallel.go#L10>)

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
