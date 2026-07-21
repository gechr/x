package bytes

import "bytes"

// SplitLines splits `s` into non-empty trimmed lines.
func SplitLines(s []byte) [][]byte {
	return SplitBy(s, []byte("\n"))
}

// SplitLinesRaw splits `s` into lines losslessly, normalizing CRLF to LF: every
// line is kept verbatim - empty lines and the trailing empty element included -
// so the result joins back with `"\n"` without losing content or line numbers.
func SplitLinesRaw(s []byte) [][]byte {
	return bytes.Split(bytes.ReplaceAll(s, []byte("\r\n"), []byte("\n")), []byte("\n"))
}
