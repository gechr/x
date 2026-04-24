package strings

// AppendCSV splits raw on commas, trims whitespace, drops empty values, and
// appends the remaining values to dst.
func AppendCSV(dst []string, raw string) []string {
	return append(dst, SplitCSV(raw)...)
}

// SplitCSV splits s on commas, trims whitespace, and drops empty values.
func SplitCSV(s string) []string {
	return SplitBy(s, ",")
}
