package http

import (
	"net/http"
	"strings"

	xstrings "github.com/gechr/x/strings"
)

// NextLink returns the rel="next" target from an RFC 8288 Link header, or ""
// when none. The target is returned as written - possibly relative - so a
// caller that needs an absolute URL resolves it against the request URL. All
// Link header lines are searched, an unquoted rel token is tolerated, and a
// quoted rel list (e.g. rel="next last") matches on any member.
func NextLink(h http.Header) string {
	for _, value := range h.Values("Link") {
		for _, link := range splitLinks(value) {
			target, params, ok := strings.Cut(link, ";")
			target, wrapped := xstrings.Unwrap(strings.TrimSpace(target), "<", ">")
			if ok && wrapped && hasRel(params, "next") {
				return target
			}
		}
	}
	return ""
}

// splitLinks splits a Link header value on the commas that separate one
// link-value from the next, ignoring commas inside a <target>.
func splitLinks(value string) []string {
	var links []string
	inTarget := false
	start := 0
	for i := range len(value) {
		switch value[i] {
		case '<':
			inTarget = true
		case '>':
			inTarget = false
		case ',':
			if !inTarget {
				links = append(links, value[start:i])
				start = i + 1
			}
		}
	}
	return append(links, value[start:])
}

// hasRel reports whether the semicolon-separated link params carry a rel
// whose (possibly quoted, space-separated) value list contains rel.
func hasRel(params, rel string) bool {
	for param := range strings.SplitSeq(params, ";") {
		name, value, ok := strings.Cut(strings.TrimSpace(param), "=")
		if !ok || !strings.EqualFold(strings.TrimSpace(name), "rel") {
			continue
		}
		for candidate := range strings.SplitSeq(strings.Trim(strings.TrimSpace(value), `"`), " ") {
			if strings.EqualFold(candidate, rel) {
				return true
			}
		}
	}
	return false
}
