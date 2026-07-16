package slices

import (
	"fmt"
	"reflect"
)

// Format returns the result of applying fmt.Sprintf(format, ...) once per
// element of the shortest slice in args, substituting each slice argument
// with its i'th element while repeating non-slice arguments unchanged. If
// args contains no slices, it returns a single formatted string.
//
// Byte slices ([]byte and named types with that underlying type) are treated
// as scalars rather than iterated, so they format as a single value.
func Format(format string, args ...any) []string {
	n := -1
	for _, arg := range args {
		if v := reflect.ValueOf(arg); isIterable(v) && (n == -1 || v.Len() < n) {
			n = v.Len()
		}
	}
	if n == -1 {
		return []string{fmt.Sprintf(format, args...)}
	}

	out := make([]string, n)
	row := make([]any, len(args))
	for i := range n {
		for j, arg := range args {
			if v := reflect.ValueOf(arg); isIterable(v) {
				row[j] = v.Index(i).Interface()
			} else {
				row[j] = arg
			}
		}
		out[i] = fmt.Sprintf(format, row...)
	}
	return out
}

// isIterable reports whether v is a slice that should be iterated per element,
// excluding byte slices which are formatted as a single scalar value.
func isIterable(v reflect.Value) bool {
	return v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8
}
