# http

```go
import "github.com/gechr/x/http"
```

Package `http` provides HTTP helpers: retryable status codes, status text, and Link headers.

## Index

- [func IsRetryableStatus(code int) bool](<#IsRetryableStatus>)
- [func NextLink(h http.Header) string](<#NextLink>)
- [func Status(code int) string](<#Status>)

<a name="IsRetryableStatus"></a>

## func [IsRetryableStatus](<https://github.com/gechr/x/blob/main/http/status.go#L20>)

```go
func IsRetryableStatus(code int) bool
```

**IsRetryableStatus** reports whether an HTTP status code represents a transient failure worth retrying: a request timeout (408), rate limiting (429), or any server error (5xx).

<details><summary><b>Example</b></summary>

```go
fmt.Println(xhttp.IsRetryableStatus(http.StatusTooManyRequests))
fmt.Println(xhttp.IsRetryableStatus(http.StatusBadGateway))
fmt.Println(xhttp.IsRetryableStatus(http.StatusNotFound))
```

Output:

```text
true
true
false
```

</details>

<a name="NextLink"></a>

## func [NextLink](<https://github.com/gechr/x/blob/main/http/link.go#L16>)

```go
func NextLink(h http.Header) string
```

**NextLink** returns the `rel="next"` target from an RFC 8288 Link header, or `""` when none. The target is returned as written - possibly relative - so a caller that needs an absolute URL resolves it against the request URL. All Link header lines are searched, an unquoted rel token is tolerated, and a quoted rel list (e.g. `rel="next last"`) matches on any member.

<details><summary><b>Example</b></summary>

```go
h := http.Header{}
h.Add(
    "Link",
    `<https://api.github.com/repos/o/r/tags?page=2>; rel="next", <https://api.github.com/repos/o/r/tags?page=5>; rel="last"`,
)

fmt.Println(xhttp.NextLink(h))
```

Output:

```text
https://api.github.com/repos/o/r/tags?page=2
```

</details>

<details><summary><b>Example (RelList)</b></summary>

A quoted rel list matches on any member, and the empty string is returned when no link carries rel="next".

```go
h := http.Header{}
h.Add("Link", `<https://example.com/?page=3>; rel="next last"`)

fmt.Println(xhttp.NextLink(h))
fmt.Println(xhttp.NextLink(http.Header{}) == "")
```

Output:

```text
https://example.com/?page=3
true
```

</details>

<a name="Status"></a>

## func [Status](<https://github.com/gechr/x/blob/main/http/status.go#L13>)

```go
func Status(code int) string
```

**Status** returns a human-readable form of an HTTP status code, pairing the numeric code with its canonical reason phrase, e.g. `404 Not Found`.

<details><summary><b>Example</b></summary>

```go
fmt.Println(xhttp.Status(http.StatusNotFound))
fmt.Println(xhttp.Status(http.StatusTeapot))
```

Output:

```text
404 Not Found
418 I'm a teapot
```

</details>
