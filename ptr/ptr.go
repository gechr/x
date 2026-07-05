// Package ptr provides pointer helpers.
package ptr

// Deref returns the value `p` points to, or the zero value when `p` is nil.
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}
