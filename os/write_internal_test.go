package os

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRawDir(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"file":            ".",
		"/file":           "/",
		"dir/file":        "dir",
		"/a/b/file":       "/a/b",
		"a/link/../file":  "a/link/..",
		"dir/įfile":       "dir",
		"dir/sub/":        "dir/sub",
		"./file":          ".",
		"../file":         "..",
		"/a/../../b/file": "/a/../../b",
	}
	for path, want := range tests {
		require.Equal(t, want, rawDir(path), path)
	}
}
