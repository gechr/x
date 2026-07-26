package strings_test

import (
	"strings"
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestSlug(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"":                        "",
		"api":                     "api",
		"API":                     "api",
		"My Service":              "my-service",
		"my_service":              "my-service",
		"my.service":              "my-service",
		"my--service":             "my-service", // a run collapses to one separator
		"my _ - . service":        "my-service",
		"  api  ":                 "api", // leading and trailing separators are dropped
		"___":                     "",    // nothing to slugify
		"αβγ":                     "",    // non-ASCII is a separator, not a letter
		"café au lait":            "caf-au-lait",
		"myService":               "myservice", // case is folded, not read as a boundary
		"v1.2.3":                  "v1-2-3",
		"1abc":                    "1abc",                    // may start with a digit
		strings.Repeat("a", 1000): strings.Repeat("a", 1000), // no length cap
	}
	for in, want := range cases {
		require.Equal(t, want, xstrings.Slug(in), "Slug(%q)", in)
	}
}

func TestSlugMaxLength(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   string
		max  int
		want string
	}{
		{"My Long Service Name", 12, "my-long"}, // the split word goes with the cut
		{"my-service", 99, "my-service"},        // a slug that fits is untouched
		{"my-service", 10, "my-service"},        // exactly the cap
		{"my-service", 9, "my"},                 // the cut would split "service"
		{"my-service", 3, "my"},
		{"my-service", 2, "my"},
		{"my-service", 1, "m"}, // a first word over the cap is cut mid-word
		{"verylongword", 4, "very"},
		{"my-service", 0, "my-service"},  // a non-positive cap is inert
		{"my-service", -1, "my-service"}, //
		{"My Service", 4, "my"},
		{"___", 5, ""}, // nothing to slugify
		{"", 5, ""},
	}
	for _, tt := range tests {
		got := xstrings.Slug(tt.in, xstrings.WithSlugMaxLength(tt.max))
		require.Equal(t, tt.want, got, "Slug(%q, WithSlugMaxLength(%d))", tt.in, tt.max)
		if tt.max > 0 {
			require.LessOrEqual(t, len(got), tt.max, "Slug(%q, max %d)", tt.in, tt.max)
		}
		if got != "" {
			require.True(t, xstrings.IsSlug(got), "IsSlug of Slug(%q, max %d)", tt.in, tt.max)
		}
	}
}

func TestSlugCutMidWord(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   string
		max  int
		want string
	}{
		{"My Long Service Name", 12, "my-long-serv"}, // the cap is filled
		{"my-service", 3, "my"},                      // the exposed separator still goes
		{"my-service", 4, "my-s"},
		{"my-service", 99, "my-service"}, // a slug that fits is untouched
		{"verylongword", 4, "very"},
		{"___", 5, ""},
	}
	for _, tt := range tests {
		got := xstrings.Slug(
			tt.in,
			xstrings.WithSlugMaxLength(tt.max),
			xstrings.WithSlugCutMidWord(),
		)
		require.Equal(t, tt.want, got, "Slug(%q, max %d, mid-word)", tt.in, tt.max)
		if got != "" {
			require.True(
				t,
				xstrings.IsSlug(got),
				"IsSlug of Slug(%q, max %d, mid-word)",
				tt.in,
				tt.max,
			)
		}
	}

	// Without a length cap there is nothing to cut, so the option is inert.
	require.Equal(t,
		"my-long-service-name",
		xstrings.Slug("My Long Service Name", xstrings.WithSlugCutMidWord()),
	)
}

func TestSlugWordBounds(t *testing.T) {
	t.Parallel()

	require.Equal(t, "my-long", xstrings.Slug("My Long Service Name", xstrings.WithSlugMaxWords(2)))
	require.Equal(t,
		"my-long-service-name",
		xstrings.Slug("My Long Service Name", xstrings.WithSlugMaxWords(99)),
	)
	require.Equal(t,
		"my-long-service-name",
		xstrings.Slug("My Long Service Name", xstrings.WithSlugMaxWords(0)), // inert
	)

	require.Equal(t,
		"bob-api-v2-service",
		xstrings.Slug("Bob's API v2 Service", xstrings.WithSlugMinWordLength(2)),
	)
	require.Empty(t, xstrings.Slug("a b c", xstrings.WithSlugMinWordLength(2)))
	require.Equal(t,
		"bob-s-api-v2-service",
		xstrings.Slug("Bob's API v2 Service", xstrings.WithSlugMinWordLength(0)), // inert
	)

	// A word the minimum drops does not consume a max-words slot: "s" is gone
	// before the count is applied, so two words means bob and api.
	require.Equal(t, "bob-api", xstrings.Slug(
		"Bob's API v2 Service",
		xstrings.WithSlugMinWordLength(2),
		xstrings.WithSlugMaxWords(2),
	))
}

func TestSlugSeparator(t *testing.T) {
	t.Parallel()

	require.Equal(t, "my_service", xstrings.Slug("My Service", xstrings.WithSlugSeparator('_')))
	require.Equal(t, "my.service", xstrings.Slug("My Service", xstrings.WithSlugSeparator('.')))
	require.Equal(t, "my-service", xstrings.Slug("My Service", xstrings.WithSlugSeparator('-')))

	// A separator is only written between words, so a run still collapses to one
	// and neither end carries it.
	require.Equal(t, "a_b", xstrings.Slug("  a...b  ", xstrings.WithSlugSeparator('_')))

	// The override forfeits IsSlug - '_' is lenient-only - while the bounds it
	// composes with still hold.
	underscored := xstrings.Slug("My Service", xstrings.WithSlugSeparator('_'))
	require.False(t, xstrings.IsSlug(underscored))
	require.True(t, xstrings.IsSlugLenient(underscored))
	require.Equal(t, "my", xstrings.Slug(
		"My Service",
		xstrings.WithSlugSeparator('_'),
		xstrings.WithSlugMaxLength(9),
	))

	// A multi-byte separator is counted in bytes and written whole or not at all,
	// so a cap that cannot afford one emits neither it nor half of its rune.
	require.Equal(t, "a×b", xstrings.Slug("a b", xstrings.WithSlugSeparator('×')))
	require.Equal(t, "a", xstrings.Slug(
		"a b",
		xstrings.WithSlugSeparator('×'),
		xstrings.WithSlugMaxLength(2),
		xstrings.WithSlugCutMidWord(),
	))

	// A word character may be the separator rune itself, and the cut must not
	// mistake one for the other: "my"+"x"+"tax" fills the cap exactly, so the 'x'
	// ending "tax" stays. Cutting the joined string and trimming the separator off
	// its tail returned "myxta" here, which is why the cut is made while joining.
	require.Equal(t, "myxtax", xstrings.Slug(
		"my tax b",
		xstrings.WithSlugSeparator('x'),
		xstrings.WithSlugMaxLength(6),
		xstrings.WithSlugCutMidWord(),
	))
}

// Slug's contract is that it agrees with [xstrings.IsSlug]: every non-empty
// result is a valid slug, and a valid slug is its own slug. The two are read
// together, so the agreement is pinned here rather than left to a doc comment.
func TestSlugAgreesWithIsSlug(t *testing.T) {
	t.Parallel()

	inputs := []string{
		"", "api", "API", "My Service", "my_service", "my.service", "my--service",
		"  api  ", "___", "αβγ", "café au lait", "v1.2.3", "-abc", "abc-", "a-_-b",
		"!!!", "a1", "GO_VERSION", "rust-analyzer", strings.Repeat("a-", 500),
	}
	for _, in := range inputs {
		slug := xstrings.Slug(in)
		if slug != "" {
			require.True(t, xstrings.IsSlug(slug), "IsSlug(Slug(%q)) with slug %q", in, slug)
		}
		require.Equal(t, slug, xstrings.Slug(slug), "Slug is idempotent for %q", in)
	}
}

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
