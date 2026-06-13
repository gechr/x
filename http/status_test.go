package http_test

import (
	"net/http"
	"testing"

	xhttp "github.com/gechr/x/http"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	t.Parallel()

	require.Equal(t, "200 OK", xhttp.Status(http.StatusOK))
	require.Equal(t, "404 Not Found", xhttp.Status(http.StatusNotFound))
	require.Equal(t, "503 Service Unavailable", xhttp.Status(http.StatusServiceUnavailable))
	// Unknown codes still render, with an empty reason phrase.
	require.Equal(t, "799 ", xhttp.Status(799))
}

func TestIsRetryableStatus(t *testing.T) {
	t.Parallel()

	require.True(t, xhttp.IsRetryableStatus(http.StatusRequestTimeout))
	require.True(t, xhttp.IsRetryableStatus(http.StatusInternalServerError))
	require.True(t, xhttp.IsRetryableStatus(http.StatusBadGateway))
	require.True(t, xhttp.IsRetryableStatus(http.StatusGatewayTimeout))

	require.False(t, xhttp.IsRetryableStatus(http.StatusOK))
	require.False(t, xhttp.IsRetryableStatus(http.StatusNotFound))
	require.False(t, xhttp.IsRetryableStatus(http.StatusBadRequest))
	require.False(t, xhttp.IsRetryableStatus(http.StatusTooManyRequests))
}
