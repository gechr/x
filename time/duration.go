// Package time provides time helpers.
package time

import "time"

// Calendar-scaled durations beyond what the standard time package names. A
// day is 24 hours, a week is 7 days, and a year is 365 days, matching
// [github.com/gechr/x/human.FormatDuration] and
// [github.com/gechr/x/human.ParseDuration].
const (
	Day  = 24 * time.Hour
	Week = 7 * Day
	Year = 365 * Day
)
