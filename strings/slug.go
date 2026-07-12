package strings

// IsSlug reports whether `s` is a valid slug: a non-empty, URL-friendly
// identifier of lowercase alphanumerics and '-', starting and ending with an
// alphanumeric (e.g. `my-service`). Underscores are not permitted; `-` is the
// only allowed separator, and it may not appear consecutively. Every valid slug
// is therefore a fixed point of slugification. An empty string is not a slug.
func IsSlug(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			// Always allowed.
		case c == '-':
			// Not allowed at the start, end, or after another '-'.
			if i == 0 || i == len(s)-1 || s[i-1] == '-' {
				return false
			}
		default:
			return false
		}
	}
	return true
}

// IsSlugLenient reports whether `s` is a valid lenient slug: a non-empty
// identifier of lowercase alphanumerics, '-', and '_', starting and ending with
// an alphanumeric (e.g. `my-service`, `my_service`, `a--b__c`). Unlike
// [IsSlug], underscores are permitted and separators may appear consecutively
// or mixed; only leading and trailing separators are rejected. An empty string
// is not a slug.
func IsSlugLenient(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			// Always allowed.
		case c == '-', c == '_':
			// Allowed internally, including consecutively and mixed, but not at
			// the start or end.
			if i == 0 || i == len(s)-1 {
				return false
			}
		default:
			return false
		}
	}
	return true
}
