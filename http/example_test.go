package http_test

import (
	"fmt"
	"net/http"

	xhttp "github.com/gechr/x/http"
)

func ExampleNextLink() {
	h := http.Header{}
	h.Add(
		"Link",
		`<https://api.github.com/repos/o/r/tags?page=2>; rel="next", <https://api.github.com/repos/o/r/tags?page=5>; rel="last"`,
	)

	fmt.Println(xhttp.NextLink(h))
	// Output:
	// https://api.github.com/repos/o/r/tags?page=2
}

// A quoted rel list matches on any member, and the empty string is returned
// when no link carries rel="next".
func ExampleNextLink_relList() {
	h := http.Header{}
	h.Add("Link", `<https://example.com/?page=3>; rel="next last"`)

	fmt.Println(xhttp.NextLink(h))
	fmt.Println(xhttp.NextLink(http.Header{}) == "")
	// Output:
	// https://example.com/?page=3
	// true
}

func ExampleStatus() {
	fmt.Println(xhttp.Status(http.StatusNotFound))
	fmt.Println(xhttp.Status(http.StatusTeapot))
	// Output:
	// 404 Not Found
	// 418 I'm a teapot
}

func ExampleIsRetryableStatus() {
	fmt.Println(xhttp.IsRetryableStatus(http.StatusTooManyRequests))
	fmt.Println(xhttp.IsRetryableStatus(http.StatusBadGateway))
	fmt.Println(xhttp.IsRetryableStatus(http.StatusNotFound))
	// Output:
	// true
	// true
	// false
}
