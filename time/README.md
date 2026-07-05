# time

```go
import "github.com/gechr/x/time"
```

Package time provides time helpers.

<details><summary><b>Example</b></summary>

The calendar-scaled durations compose with the standard time package.

```go
fmt.Println(xtime.Day)
fmt.Println(xtime.Week)
fmt.Println(xtime.Year)
```

Output:

```text
24h0m0s
168h0m0s
8760h0m0s
```

</details>

## Index

- [Constants](<#constants>)

## Constants

<a name="Day"></a>Calendar-scaled durations beyond what the standard time package names. A day is 24 hours, a week is 7 days, and a year is 365 days, matching [human.FormatDuration](<../human/README.md#FormatDuration>) and [human.ParseDuration](<../human/README.md#ParseDuration>).

```go
const (
    Day  = 24 * time.Hour
    Week = 7 * Day
    Year = 365 * Day
)
```
