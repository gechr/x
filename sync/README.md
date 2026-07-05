# sync

```go
import "github.com/gechr/x/sync"
```

Package sync provides concurrency helpers.

## Index

- [func Parallel(workers, n int, fn func(i int))](<#Parallel>)

<a name="Parallel"></a>

## func [Parallel](<https://github.com/gechr/x/blob/main/sync/parallel.go#L10>)

```go
func Parallel(workers, n int, fn func(i int))
```

**Parallel** runs `fn(0)` through `fn(n-1)` concurrently with at most `workers` in flight, blocking until all complete. Each call receives a distinct index, so a goroutine writing results\[i\] needs no lock; `fn` must otherwise be safe to call concurrently. `workers` \< 1 runs one call at a time.
