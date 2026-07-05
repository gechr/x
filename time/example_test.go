package time_test

import (
	"fmt"

	xtime "github.com/gechr/x/time"
)

// The calendar-scaled durations compose with the standard time package.
func Example() {
	fmt.Println(xtime.Day)
	fmt.Println(xtime.Week)
	fmt.Println(xtime.Year)
	// Output:
	// 24h0m0s
	// 168h0m0s
	// 8760h0m0s
}
