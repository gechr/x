# time

```go
import "github.com/gechr/x/time"
```

Package time provides time helpers.

## Index

- [Constants](<#constants>)

## Constants

<a name="Day"></a>Calendar-scaled durations beyond what the standard time package names. A day is 24 hours, a week is 7 days, and a year is 365 days, matching [github.com/gechr/x/human.FormatDuration](<https://pkg.go.dev/github.com/gechr/x/human#FormatDuration>) and [github.com/gechr/x/human.ParseDuration](<https://pkg.go.dev/github.com/gechr/x/human#ParseDuration>).

```go
const (
    Day  = 24 * time.Hour
    Week = 7 * Day
    Year = 365 * Day
)
```
