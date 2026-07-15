# human

```go
import "github.com/gechr/x/human"
```

Package `human` formats bytes, durations, counts, numbers, and ordinals for human-readable output.

## Index

- [Constants](<#constants>)
- [func ContractHome(path string) string](<#ContractHome>)
- [func FormatDuration(d time.Duration, options ...DurationFormatOptions) string](<#FormatDuration>)
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
- [type DurationFormatOptions](<#DurationFormatOptions>)

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

<details><summary><b>Example</b></summary>

**ContractHome** leaves paths outside the home directory untouched.

```go
home, _ := os.UserHomeDir()
fmt.Println(human.ContractHome(filepath.Join(home, "projects", "x")))
fmt.Println(human.ContractHome("/etc/hosts"))
```

Output:

```text
~/projects/x
/etc/hosts
```

</details>

<a name="FormatDuration"></a>

## func [FormatDuration](<https://github.com/gechr/x/blob/main/human/duration.go#L43>)

```go
func FormatDuration(d time.Duration, options ...DurationFormatOptions) string
```

**FormatDuration** formats d as up to two adjacent units with no separator (e.g. "5.2s", "2h15m", "1w2d"). Years are 365 days, weeks are 7 days. With no options it uses magnitude-based resolution: the largest exact sub-second unit, one decimal place below ten seconds, and whole seconds thereafter. When multiple options values are supplied, the last applies.

```go
FormatDuration(90 * time.Second)             // "1m30s"
FormatDuration(2*time.Hour + 15*time.Minute) // "2h15m"
FormatDuration(8 * 24 * time.Hour)           // "1w1d"
FormatDuration(400 * 24 * time.Hour)         // "1y5w"
FormatDuration(50 * time.Millisecond)        // "50ms"
FormatDuration(5200 * time.Millisecond)      // "5.2s"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatDuration(90 * time.Second))
fmt.Println(human.FormatDuration(2*time.Hour + 15*time.Minute))
fmt.Println(human.FormatDuration(400 * 24 * time.Hour))
fmt.Println(human.FormatDuration(50 * time.Millisecond))
fmt.Println(human.FormatDuration(5200 * time.Millisecond))
```

Output:

```text
1m30s
2h15m
1y5w
50ms
5.2s
```

</details>

<a name="FormatIECBytes"></a>

## func [FormatIECBytes](<https://github.com/gechr/x/blob/main/human/bytes.go#L108>)

```go
func FormatIECBytes(b float64) string
```

**FormatIECBytes** formats a byte count using IEC binary units (KiB, MiB, GiB, TiB, PiB, EiB).

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatIECBytes(512))
fmt.Println(human.FormatIECBytes(1536))
fmt.Println(human.FormatIECBytes(5 * human.GiB))
```

Output:

```text
512 B
1.50 KiB
5.00 GiB
```

</details>

<a name="FormatNumber"></a>

## func [FormatNumber](<https://github.com/gechr/x/blob/main/human/count.go#L37>)

```go
func FormatNumber(n int64, sep string) string
```

**FormatNumber** groups `n`'s digits in threes from the right, joined with `sep`. Not locale-aware: pick a separator suited to your output.

```go
FormatNumber(1234567, ",") // "1,234,567"
FormatNumber(1234567, ".") // "1.234.567"
FormatNumber(1234567, " ") // "1 234 567"
FormatNumber(-42, ",")     // "-42"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatNumber(1234567, ","))
fmt.Println(human.FormatNumber(1234567, "."))
fmt.Println(human.FormatNumber(1234567, " "))
fmt.Println(human.FormatNumber(-42, ","))
```

Output:

```text
1,234,567
1.234.567
1 234 567
-42
```

</details>

<a name="FormatNumberCompact"></a>

## func [FormatNumberCompact](<https://github.com/gechr/x/blob/main/human/count.go#L87>)

```go
func FormatNumberCompact(n int64) string
```

**FormatNumberCompact** renders `n` in a compact, abbreviated form using K, M, B, and T suffixes (powers of 1000), with up to one decimal place and a trailing `.0` trimmed. Values whose magnitude is below 1000 are returned verbatim. Values that round up to the next unit are promoted (e.g. 999999 → `1M`), and magnitudes beyond a trillion stay in `T`.

```go
FormatNumberCompact(950)      // "950"
FormatNumberCompact(1234)     // "1.2K"
FormatNumberCompact(1000000)  // "1M"
FormatNumberCompact(9999999)  // "10M"
FormatNumberCompact(-1500000) // "-1.5M"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatNumberCompact(950))
fmt.Println(human.FormatNumberCompact(1500))
fmt.Println(human.FormatNumberCompact(999999))
fmt.Println(human.FormatNumberCompact(3400000000))
```

Output:

```text
950
1.5K
1M
3.4B
```

</details>

<a name="FormatOrdinal"></a>

## func [FormatOrdinal](<https://github.com/gechr/x/blob/main/human/count.go#L127>)

```go
func FormatOrdinal(n int) string
```

**FormatOrdinal** returns `n` with its English ordinal suffix.

```go
FormatOrdinal(1)   // "1st"
FormatOrdinal(22)  // "22nd"
FormatOrdinal(113) // "113th"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatOrdinal(1))
fmt.Println(human.FormatOrdinal(22))
fmt.Println(human.FormatOrdinal(113))
```

Output:

```text
1st
22nd
113th
```

</details>

<a name="FormatSIBytes"></a>

## func [FormatSIBytes](<https://github.com/gechr/x/blob/main/human/bytes.go#L103>)

```go
func FormatSIBytes(b float64) string
```

**FormatSIBytes** formats a byte count using SI decimal units (KB, MB, GB, TB, PB, EB).

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.FormatSIBytes(512))
fmt.Println(human.FormatSIBytes(1500))
fmt.Println(human.FormatSIBytes(5 * human.GB))
```

Output:

```text
512 B
1.50 KB
5.00 GB
```

</details>

<a name="FormatTimeAgo"></a>

## func [FormatTimeAgo](<https://github.com/gechr/x/blob/main/human/time.go#L27>)

```go
func FormatTimeAgo(t time.Time) string
```

**FormatTimeAgo** formats a time as a human-readable relative string (plain text).

<details><summary><b>Example</b></summary>

**FormatTimeAgo** formats a time relative to the current time; see [human.FormatTimeAgoFrom](<#FormatTimeAgoFrom>) for the full range of outputs.

```go
fmt.Println(human.FormatTimeAgo(time.Now().Add(-90 * time.Minute)))
```

Output:

```text
1 hour ago
```

</details>

<a name="FormatTimeAgoCompact"></a>

## func [FormatTimeAgoCompact](<https://github.com/gechr/x/blob/main/human/time.go#L47>)

```go
func FormatTimeAgoCompact(t time.Time) string
```

**FormatTimeAgoCompact** formats a time as a compact relative string (e.g. `15m ago`).

<details><summary><b>Example</b></summary>

**FormatTimeAgoCompact** formats a time relative to the current time; see [human.FormatTimeAgoCompactFrom](<#FormatTimeAgoCompactFrom>) for the full range of outputs.

```go
fmt.Println(human.FormatTimeAgoCompact(time.Now().Add(-90 * time.Minute)))
```

Output:

```text
1h ago
```

</details>

<a name="FormatTimeAgoCompactFrom"></a>

## func [FormatTimeAgoCompactFrom](<https://github.com/gechr/x/blob/main/human/time.go#L52>)

```go
func FormatTimeAgoCompactFrom(t, now time.Time) string
```

**FormatTimeAgoCompactFrom** formats a time as a compact relative string relative to `now`.

<details><summary><b>Example</b></summary>

```go
now := time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)
fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-30*time.Second), now))
fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-90*time.Minute), now))
fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-3*24*time.Hour), now))
fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(2*time.Hour), now))
```

Output:

```text
now
1h ago
3d ago
in 2h
```

</details>

<a name="FormatTimeAgoFrom"></a>

## func [FormatTimeAgoFrom](<https://github.com/gechr/x/blob/main/human/time.go#L32>)

```go
func FormatTimeAgoFrom(t, now time.Time) string
```

**FormatTimeAgoFrom** formats a time relative to the given reference time `now`.

<details><summary><b>Example</b></summary>

```go
now := time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)
fmt.Println(human.FormatTimeAgoFrom(now.Add(-30*time.Second), now))
fmt.Println(human.FormatTimeAgoFrom(now.Add(-90*time.Minute), now))
fmt.Println(human.FormatTimeAgoFrom(now.Add(-3*24*time.Hour), now))
fmt.Println(human.FormatTimeAgoFrom(now.Add(2*time.Hour), now))
```

Output:

```text
now
1 hour ago
3 days ago
in 2 hours
```

</details>

<a name="ParseByteSize"></a>

## func [ParseByteSize](<https://github.com/gechr/x/blob/main/human/bytes.go#L50>)

```go
func ParseByteSize(s string) float64
```

**ParseByteSize** parses a human-readable byte size string like `27.61 MiB` or `1.5 GB` into a byte count. Supports both IEC (KiB, MiB, GiB, TiB, PiB, EiB) and SI (KB, MB, GB, TB, PB, EB) units. Returns 0 for empty or unparseable input.

<details><summary><b>Example</b></summary>

```go
fmt.Printf("%.0f\n", human.ParseByteSize("1.5 GB"))
fmt.Printf("%.0f\n", human.ParseByteSize("27.61 MiB"))
```

Output:

```text
1500000000
28951183
```

</details>

<a name="ParseDuration"></a>

## func [ParseDuration](<https://github.com/gechr/x/blob/main/human/duration.go#L179>)

```go
func ParseDuration(s string) (time.Duration, error)
```

**ParseDuration** parses a human duration string into a [time.Duration](<https://pkg.go.dev/time#Duration>). It is the inverse of [FormatDuration](<#FormatDuration>), accepting decimal values and the units y, w, d, h, m, s, ms, µs (or us), and ns, where a year is 365 days and a week is 7 days. Units may be combined but each may appear at most once and must run in descending order of size. A decimal value must be the final component. An optional leading - negates the result, and `0` parses to zero.

```go
ParseDuration("2h15m")  // 2*time.Hour + 15*time.Minute
ParseDuration("1w2d")   // 9 * 24 * time.Hour
ParseDuration("5.2s")   // 5200 * time.Millisecond
ParseDuration("-1m30s") // -90 * time.Second
```

<details><summary><b>Example</b></summary>

**ParseDuration** is the inverse of [human.FormatDuration](<#FormatDuration>), accepting the y, w, d, h, m, s, ms, µs, and ns units that function emits.

```go
d, _ := human.ParseDuration("2h15m")
fmt.Println(d)
d, _ = human.ParseDuration("1w2d")
fmt.Println(d)
d, _ = human.ParseDuration("-1m30s")
fmt.Println(d)
```

Output:

```text
2h15m0s
216h0m0s
-1m30s
```

</details>

<a name="Plural"></a>

## func [Plural](<https://github.com/gechr/x/blob/main/human/count.go#L15>)

```go
func Plural(n int, singular, plural string) string
```

**Plural** returns `singular` when `n == 1`, otherwise `plural`. Unlike [Pluralize](<#Pluralize>), it omits the count.

```go
Plural(1, "file", "files") // "file"
Plural(3, "file", "files") // "files"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.Plural(1, "file", "files"))
fmt.Println(human.Plural(3, "file", "files"))
```

Output:

```text
file
files
```

</details>

<a name="Pluralize"></a>

## func [Pluralize](<https://github.com/gechr/x/blob/main/human/count.go#L26>)

```go
func Pluralize(n int, singular, plural string) string
```

**Pluralize** returns "1 singular" or "n plural".

```go
Pluralize(1, "file", "files") // "1 file"
Pluralize(3, "file", "files") // "3 files"
```

<details><summary><b>Example</b></summary>

```go
fmt.Println(human.Pluralize(1, "file", "files"))
fmt.Println(human.Pluralize(3, "file", "files"))
```

Output:

```text
1 file
3 files
```

</details>

<a name="DurationFormatOptions"></a>

## type [DurationFormatOptions](<https://github.com/gechr/x/blob/main/human/duration.go#L21-L29>)

**DurationFormatOptions** controls duration rounding and fractional display.

```go
type DurationFormatOptions struct {
    // Precision is the number of decimal places used for durations below one
    // minute.
    Precision int
    // Round is the rounding granularity. Zero disables pre-format rounding.
    Round time.Duration
    // TrimTrailingZeros removes trailing zeroes from the fractional part.
    TrimTrailingZeros bool
}
```
