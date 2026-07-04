# http

```go
import "github.com/gechr/x/http"
```

Package http provides HTTP helpers: retryable status codes, status text, and Link headers.

## Index

- [func IsRetryableStatus\(code int\) bool](<#IsRetryableStatus>)
- [func NextLink\(h http.Header\) string](<#NextLink>)
- [func Status\(code int\) string](<#Status>)

<a name="IsRetryableStatus"></a>

## func [IsRetryableStatus](<https://github.com/gechr/x/blob/main/http/status.go#L20>)

```go
func IsRetryableStatus(code int) bool
```

IsRetryableStatus reports whether an HTTP status code represents a transient failure worth retrying: a request timeout \(408\), rate limiting \(429\), or any server error \(5xx\).

<a name="NextLink"></a>

## func [NextLink](<https://github.com/gechr/x/blob/main/http/link.go#L16>)

```go
func NextLink(h http.Header) string
```

NextLink returns the rel="next" target from an RFC 8288 Link header, or "" when none. The target is returned as written \- possibly relative \- so a caller that needs an absolute URL resolves it against the request URL. All Link header lines are searched, an unquoted rel token is tolerated, and a quoted rel list \(e.g. rel="next last"\) matches on any member.

<a name="Status"></a>

## func [Status](<https://github.com/gechr/x/blob/main/http/status.go#L13>)

```go
func Status(code int) string
```

Status returns a human\-readable form of an HTTP status code, pairing the numeric code with its canonical reason phrase, e.g. "404 Not Found".
