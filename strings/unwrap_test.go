package strings_test

import (
	"testing"

	xstrings "github.com/gechr/x/strings"
	"github.com/stretchr/testify/require"
)

func TestUnwrap(t *testing.T) {
	t.Parallel()

	unwrap := func(s, prefix, suffix, want string, wantOK bool) {
		t.Helper()
		got, ok := xstrings.Unwrap(s, prefix, suffix)
		require.Equal(t, want, got)
		require.Equal(t, wantOK, ok)
	}

	// Both present.
	unwrap("<https://example.com>", "<", ">", "https://example.com", true)
	unwrap(`"quoted"`, `"`, `"`, "quoted", true)
	unwrap("[[link]]", "[[", "]]", "link", true)
	unwrap("<>", "<", ">", "", true)
	unwrap("s", "", "", "s", true)

	// One side missing - returned unchanged.
	unwrap("<half", "<", ">", "<half", false)
	unwrap("half>", "<", ">", "half>", false)
	unwrap("neither", "<", ">", "neither", false)
	unwrap("", "<", ">", "", false)

	// Prefix and suffix must not overlap.
	unwrap(`"`, `"`, `"`, `"`, false)
	unwrap("<", "<", "<", "<", false)
}
