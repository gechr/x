package http

import (
	"fmt"
	"net/http"
)

// maxServerErrorStatus is the highest server-error (5xx) status code.
const maxServerErrorStatus = 599

// Status returns a human-readable form of an HTTP status code, pairing the
// numeric code with its canonical reason phrase, e.g. `404 Not Found`.
func Status(code int) string {
	return fmt.Sprintf("%d %s", code, http.StatusText(code))
}

// IsRetryableStatus reports whether an HTTP status code represents a transient
// failure worth retrying: a request timeout (408), rate limiting (429), or any
// server error (5xx).
func IsRetryableStatus(code int) bool {
	return code == http.StatusRequestTimeout ||
		code == http.StatusTooManyRequests ||
		(code >= http.StatusInternalServerError && code <= maxServerErrorStatus)
}
