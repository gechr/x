# human

```go
import "github.com/gechr/x/human"
```

Package human formats bytes, durations, counts, numbers, and ordinals for human-readable output.

## Index

- [Constants](<#constants>)
- [func ContractHome(path string) string](<#ContractHome>)
- [func FormatDuration(d time.Duration) string](<#FormatDuration>)
- [func FormatIECBytes(b float64) string](<#FormatIECBytes>)
- [func FormatNumber(n int64, sep string) string](<#FormatNumber>)
- [func FormatNumberCompact(n int64) string](<#FormatNumberCompact>)
- [func FormatOrdinal(n int) string](<#FormatOrdinal>)
- [func FormatSIBytes(b float64) string](<#FormatSIBytes>)
- [func FormatTimeAgo(t time.Time) string](<#FormatTimeAgo>)
- [func FormatTimeAgoCompact(t time.Time) string](<#FormatTimeAgoCompact>)
- [func FormatTimeAgoCompactFrom(t, now time.Time) string](<#FormatTimeAgoCompactFrom>)
- [func FormatTimeAgoFrom(t, now time.Time) string](<#FormatTimeAgoFrom>)
- [func ParseByteSize(s string) float64](<#ParseByteSize>)
- [func ParseDuration(s string) (time.Duration, error)](<#ParseDuration>)
- [func Plural(n int, singular, plural string) string](<#Plural>)
- [func Pluralize(n int, singular, plural string) string](<#Pluralize>)

## Constants

<a name="KB"></a>SI byte size constants (powers of 1000).

```go
const (
    KB  = 1000
    MB  = 1000 * KB
    GB  = 1000 * MB
    TB  = 1000 * GB
    PB  = 1000 * TB
    EB  = 1000 * PB
)
```

<a name="KiB"></a>IEC byte size constants (powers of 1024).

```go
const (
    KiB = 1024
    MiB = 1024 * KiB
    GiB = 1024 * MiB
    TiB = 1024 * GiB
    PiB = 1024 * TiB
    EiB = 1024 * PiB
)
```

<a name="UnitB"></a>Unit label constants.

```go
const (
    UnitB   = "B"
    UnitKB  = "KB"
    UnitMB  = "MB"
    UnitGB  = "GB"
    UnitTB  = "TB"
    UnitPB  = "PB"
    UnitEB  = "EB"
    UnitKiB = "KiB"
    UnitMiB = "MiB"
    UnitGiB = "GiB"
    UnitTiB = "TiB"
    UnitPiB = "PiB"
    UnitEiB = "EiB"
)
```

<a name="SecondsPerMinute"></a>Time arithmetic constants.

```go
const (
    SecondsPerMinute = 60

    MinutesPerHour = 60

    HoursPerDay = 24

    DaysPerWeek  = 7
    DaysPerMonth = 30
    DaysPerYear  = 365

    WeeksPerMonth = 4
    WeeksPerYear  = 52

    MonthsPerYear = 12
)
```

<a name="ContractHome"></a>

## func [ContractHome](<https://github.com/gechr/x/blob/main/human/path.go#L10>)

```go
func ContractHome(path string) string
```

**ContractHome** replaces the user's home directory prefix with ~.

<a name="FormatDuration"></a>

## func [FormatDuration](<https://github.com/gechr/x/blob/main/human/duration.go#L23>)

```go
func FormatDuration(d time.Duration) string
```

**FormatDuration** formats `d` as up to two adjacent units with no separator (e.g. "2h15m", "1w2d", "1y5w"). Years are 365 days, weeks are 7 days. Durations >= 1s are rounded to the nearest second.

```text
FormatDuration(90 * time.Second)             // "1m30s"
FormatDuration(2*time.Hour + 15*time.Minute) // "2h15m"
FormatDuration(8 * 24 * time.Hour)           // "1w1d"
FormatDuration(400 * 24 * time.Hour)         // "1y5w"
FormatDuration(50 * time.Millisecond)        // "50ms"
```

<a name="FormatIECBytes"></a>

## func [FormatIECBytes](<https://github.com/gechr/x/blob/main/human/bytes.go#L108>)

```go
func FormatIECBytes(b float64) string
```

**FormatIECBytes** formats a byte count using IEC binary units (KiB, MiB, GiB, TiB, PiB, EiB).

<a name="FormatNumber"></a>

## func [FormatNumber](<https://github.com/gechr/x/blob/main/human/count.go#L37>)

```go
func FormatNumber(n int64, sep string) string
```

**FormatNumber** groups `n`'s digits in threes from the right, joined with `sep`. Not locale-aware: pick a separator suited to your output.

```text
FormatNumber(1234567, ",") // "1,234,567"
FormatNumber(1234567, ".") // "1.234.567"
FormatNumber(1234567, " ") // "1 234 567"
FormatNumber(-42, ",")     // "-42"
```

<a name="FormatNumberCompact"></a>

## func [FormatNumberCompact](<https://github.com/gechr/x/blob/main/human/count.go#L87>)

```go
func FormatNumberCompact(n int64) string
```

**FormatNumberCompact** renders `n` in a compact, abbreviated form using K, M, B, and T suffixes (powers of 1000), with up to one decimal place and a trailing ".0" trimmed. Values whose magnitude is below 1000 are returned verbatim. Values that round up to the next unit are promoted (e.g. 999999 → "1M"), and magnitudes beyond a trillion stay in "T".

```text
FormatNumberCompact(950)      // "950"
FormatNumberCompact(1234)     // "1.2K"
FormatNumberCompact(1000000)  // "1M"
FormatNumberCompact(9999999)  // "10M"
FormatNumberCompact(-1500000) // "-1.5M"
```

<a name="FormatOrdinal"></a>

## func [FormatOrdinal](<https://github.com/gechr/x/blob/main/human/count.go#L127>)

```go
func FormatOrdinal(n int) string
```

**FormatOrdinal** returns `n` with its English ordinal suffix.

```text
FormatOrdinal(1)   // "1st"
FormatOrdinal(22)  // "22nd"
FormatOrdinal(113) // "113th"
```

<a name="FormatSIBytes"></a>

## func [FormatSIBytes](<https://github.com/gechr/x/blob/main/human/bytes.go#L103>)

```go
func FormatSIBytes(b float64) string
```

**FormatSIBytes** formats a byte count using SI decimal units (KB, MB, GB, TB, PB, EB).

<a name="FormatTimeAgo"></a>

## func [FormatTimeAgo](<https://github.com/gechr/x/blob/main/human/time.go#L27>)

```go
func FormatTimeAgo(t time.Time) string
```

**FormatTimeAgo** formats a time as a human-readable relative string (plain text).

<a name="FormatTimeAgoCompact"></a>

## func [FormatTimeAgoCompact](<https://github.com/gechr/x/blob/main/human/time.go#L47>)

```go
func FormatTimeAgoCompact(t time.Time) string
```

**FormatTimeAgoCompact** formats a time as a compact relative string (e.g. "15m ago").

<a name="FormatTimeAgoCompactFrom"></a>

## func [FormatTimeAgoCompactFrom](<https://github.com/gechr/x/blob/main/human/time.go#L52>)

```go
func FormatTimeAgoCompactFrom(t, now time.Time) string
```

**FormatTimeAgoCompactFrom** formats a time as a compact relative string relative to `now`.

<a name="FormatTimeAgoFrom"></a>

## func [FormatTimeAgoFrom](<https://github.com/gechr/x/blob/main/human/time.go#L32>)

```go
func FormatTimeAgoFrom(t, now time.Time) string
```

**FormatTimeAgoFrom** formats a time relative to the given reference time `now`.

<a name="ParseByteSize"></a>

## func [ParseByteSize](<https://github.com/gechr/x/blob/main/human/bytes.go#L50>)

```go
func ParseByteSize(s string) float64
```

**ParseByteSize** parses a human-readable byte size string like "27.61 MiB" or "1.5 GB" into a byte count. Supports both IEC (KiB, MiB, GiB, TiB, PiB, EiB) and SI (KB, MB, GB, TB, PB, EB) units. Returns 0 for empty or unparseable input.

<a name="ParseDuration"></a>

## func [ParseDuration](<https://github.com/gechr/x/blob/main/human/duration.go#L106>)

```go
func ParseDuration(s string) (time.Duration, error)
```

**ParseDuration** parses a human duration string into a [time.Duration](<https://pkg.go.dev/time#Duration>). It is the inverse of [FormatDuration](<#FormatDuration>), accepting the units that function emits: y, w, d, h, m, s, ms, µs (or us), and ns, where a year is 365 days and a week is 7 days. Units may be combined but each may appear at most once and must run in descending order of size, so "1y2w", "2h15m", and "90s" are valid while a repeated ("5w5w") or out-of-order ("1w1y") unit is an error. An optional leading - negates the result, and "0" parses to zero.

```text
ParseDuration("2h15m")  // 2*time.Hour + 15*time.Minute
ParseDuration("1w2d")   // 9 * 24 * time.Hour
ParseDuration("-1m30s") // -90 * time.Second
```

<a name="Plural"></a>

## func [Plural](<https://github.com/gechr/x/blob/main/human/count.go#L15>)

```go
func Plural(n int, singular, plural string) string
```

**Plural** returns `singular` when `n` == 1, otherwise `plural`. Unlike [Pluralize](<#Pluralize>), it omits the count.

```text
Plural(1, "file", "files") // "file"
Plural(3, "file", "files") // "files"
```

<a name="Pluralize"></a>

## func [Pluralize](<https://github.com/gechr/x/blob/main/human/count.go#L26>)

```go
func Pluralize(n int, singular, plural string) string
```

**Pluralize** returns "1 singular" or "n plural".

```text
Pluralize(1, "file", "files") // "1 file"
Pluralize(3, "file", "files") // "3 files"
```
