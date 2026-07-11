//go:build !unix

package os

// isUnix is false when built outside Go's `unix` build constraint (e.g. Windows,
// Plan 9, WASM).
const isUnix = false
