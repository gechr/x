package strings_test

import (
	"strings"
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestIsSlug(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":                        false,
		"a":                       true,
		"api":                     true,
		"a1":                      true,
		"my-service":              true,
		"a-b-c":                   true,
		"1abc":                    true, // may start with a digit
		"-abc":                    false,
		"abc-":                    false,
		"ABC":                     false, // uppercase not allowed
		"my.service":              false, // dots not allowed
		"a_b_c":                   false, // underscores not allowed
		"my--service":             false, // consecutive hyphens not allowed
		strings.Repeat("a", 1000): true,  // no length cap
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsSlug(in), "IsSlug(%q)", in)
	}
}

func TestIsSlugLenient(t *testing.T) {
	t.Parallel()

	cases := map[string]bool{
		"":            false,
		"a":           true,
		"my-service":  true,
		"my_service":  true,
		"my--service": true, // consecutive hyphens allowed
		"my__service": true, // consecutive underscores allowed
		"a-_-b":       true, // mixed separators allowed
		"1abc":        true, // may start with a digit
		"-abc":        false,
		"abc-":        false,
		"_abc":        false, // leading underscore not allowed
		"abc_":        false, // trailing underscore not allowed
		"ABC":         false, // uppercase not allowed
		"my.service":  false, // dots not allowed
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.IsSlugLenient(in), "IsSlugLenient(%q)", in)
	}
}
