package human_test

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gechr/x/human"
)

func ExampleFormatDuration() {
	fmt.Println(human.FormatDuration(90 * time.Second))
	fmt.Println(human.FormatDuration(2*time.Hour + 15*time.Minute))
	fmt.Println(human.FormatDuration(400 * 24 * time.Hour))
	fmt.Println(human.FormatDuration(50 * time.Millisecond))
	fmt.Println(human.FormatDuration(5200 * time.Millisecond))
	// Output:
	// 1m30s
	// 2h15m
	// 1y5w
	// 50ms
	// 5.2s
}

// ParseDuration is the inverse of [human.FormatDuration], accepting the y, w,
// d, h, m, s, ms, µs, and ns units that function emits.
func ExampleParseDuration() {
	d, _ := human.ParseDuration("2h15m")
	fmt.Println(d)
	d, _ = human.ParseDuration("1w2d")
	fmt.Println(d)
	d, _ = human.ParseDuration("-1m30s")
	fmt.Println(d)
	// Output:
	// 2h15m0s
	// 216h0m0s
	// -1m30s
}

// ContractHome leaves paths outside the home directory untouched.
func ExampleContractHome() {
	home, _ := os.UserHomeDir()
	fmt.Println(human.ContractHome(filepath.Join(home, "projects", "x")))
	fmt.Println(human.ContractHome("/etc/hosts"))
	// Output:
	// ~/projects/x
	// /etc/hosts
}

// FormatTimeAgo formats a time relative to the current time; see
// [human.FormatTimeAgoFrom] for the full range of outputs.
func ExampleFormatTimeAgo() {
	fmt.Println(human.FormatTimeAgo(time.Now().Add(-90 * time.Minute)))
	// Output:
	// 1 hour ago
}

// FormatTimeAgoCompact formats a time relative to the current time; see
// [human.FormatTimeAgoCompactFrom] for the full range of outputs.
func ExampleFormatTimeAgoCompact() {
	fmt.Println(human.FormatTimeAgoCompact(time.Now().Add(-90 * time.Minute)))
	// Output:
	// 1h ago
}

func ExampleFormatIECBytes() {
	fmt.Println(human.FormatIECBytes(512))
	fmt.Println(human.FormatIECBytes(1536))
	fmt.Println(human.FormatIECBytes(5 * human.GiB))
	// Output:
	// 512 B
	// 1.50 KiB
	// 5.00 GiB
}

func ExampleParseByteSize() {
	fmt.Printf("%.0f\n", human.ParseByteSize("1.5 GB"))
	fmt.Printf("%.0f\n", human.ParseByteSize("27.61 MiB"))
	// Output:
	// 1500000000
	// 28951183
}

func ExamplePluralize() {
	fmt.Println(human.Pluralize(1, "file", "files"))
	fmt.Println(human.Pluralize(3, "file", "files"))
	// Output:
	// 1 file
	// 3 files
}

func ExampleFormatNumberCompact() {
	fmt.Println(human.FormatNumberCompact(950))
	fmt.Println(human.FormatNumberCompact(1500))
	fmt.Println(human.FormatNumberCompact(999999))
	fmt.Println(human.FormatNumberCompact(3400000000))
	// Output:
	// 950
	// 1.5K
	// 1M
	// 3.4B
}

func ExampleFormatOrdinal() {
	fmt.Println(human.FormatOrdinal(1))
	fmt.Println(human.FormatOrdinal(22))
	fmt.Println(human.FormatOrdinal(113))
	// Output:
	// 1st
	// 22nd
	// 113th
}

func ExampleFormatTimeAgoFrom() {
	now := time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(human.FormatTimeAgoFrom(now.Add(-30*time.Second), now))
	fmt.Println(human.FormatTimeAgoFrom(now.Add(-90*time.Minute), now))
	fmt.Println(human.FormatTimeAgoFrom(now.Add(-3*24*time.Hour), now))
	fmt.Println(human.FormatTimeAgoFrom(now.Add(2*time.Hour), now))
	// Output:
	// now
	// 1 hour ago
	// 3 days ago
	// in 2 hours
}

func ExampleFormatNumber() {
	fmt.Println(human.FormatNumber(1234567, ","))
	fmt.Println(human.FormatNumber(1234567, "."))
	fmt.Println(human.FormatNumber(1234567, " "))
	fmt.Println(human.FormatNumber(-42, ","))
	// Output:
	// 1,234,567
	// 1.234.567
	// 1 234 567
	// -42
}

func ExampleFormatSIBytes() {
	fmt.Println(human.FormatSIBytes(512))
	fmt.Println(human.FormatSIBytes(1500))
	fmt.Println(human.FormatSIBytes(5 * human.GB))
	// Output:
	// 512 B
	// 1.50 KB
	// 5.00 GB
}

func ExamplePlural() {
	fmt.Println(human.Plural(1, "file", "files"))
	fmt.Println(human.Plural(3, "file", "files"))
	// Output:
	// file
	// files
}

func ExampleFormatTimeAgoCompactFrom() {
	now := time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-30*time.Second), now))
	fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-90*time.Minute), now))
	fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(-3*24*time.Hour), now))
	fmt.Println(human.FormatTimeAgoCompactFrom(now.Add(2*time.Hour), now))
	// Output:
	// now
	// 1h ago
	// 3d ago
	// in 2h
}
