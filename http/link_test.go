package http_test

import (
	"net/http"
	"testing"

	xhttp "github.com/gechr/x/http"
	"github.com/stretchr/testify/require"
)

func TestNextLink(t *testing.T) {
	t.Parallel()

	header := func(values ...string) http.Header {
		h := http.Header{}
		for _, v := range values {
			h.Add("Link", v)
		}
		return h
	}

	// A single quoted rel="next" link, as GitHub and GitLab emit.
	require.Equal(t,
		"https://api.github.com/repos/o/r/tags?page=2",
		xhttp.NextLink(header(`<https://api.github.com/repos/o/r/tags?page=2>; rel="next"`)),
	)

	// rel="next" among other links, in any position.
	require.Equal(t,
		"https://example.com/?page=3",
		xhttp.NextLink(header(
			`<https://example.com/?page=1>; rel="prev", <https://example.com/?page=3>; rel="next", <https://example.com/?page=9>; rel="last"`,
		)),
	)

	// An unquoted rel token, as some OCI registries emit.
	require.Equal(t,
		"/v2/library/alpine/tags/list?last=3.19",
		xhttp.NextLink(header(`</v2/library/alpine/tags/list?last=3.19>; rel=next`)),
	)

	// A quoted rel list matches on any member, case-insensitively.
	require.Equal(t,
		"https://example.com/?page=2",
		xhttp.NextLink(header(`<https://example.com/?page=2>; rel="Next last"`)),
	)

	// Extra params before rel do not confuse the parse.
	require.Equal(t,
		"https://example.com/?page=2",
		xhttp.NextLink(header(`<https://example.com/?page=2>; title="two"; rel="next"`)),
	)

	// A comma inside the <target> is not a link separator.
	require.Equal(t,
		"https://example.com/?fields=a,b&page=2",
		xhttp.NextLink(header(`<https://example.com/?fields=a,b&page=2>; rel="next"`)),
	)

	// All Link header lines are searched, not just the first.
	require.Equal(t,
		"https://example.com/?page=2",
		xhttp.NextLink(header(
			`<https://example.com/?page=1>; rel="prev"`,
			`<https://example.com/?page=2>; rel="next"`,
		)),
	)

	// No rel="next": other rels, a bare target, an empty header.
	require.Empty(t, xhttp.NextLink(header(`<https://example.com/?page=9>; rel="last"`)))
	require.Empty(t, xhttp.NextLink(header(`<https://example.com/?page=2>`)))
	require.Empty(t, xhttp.NextLink(header(`rel="next"`)))
	require.Empty(t, xhttp.NextLink(http.Header{}))
}
