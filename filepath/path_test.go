package filepath_test

import (
	ifs "io/fs"
	"os"
	"path/filepath"
	"testing"

	xfilepath "github.com/gechr/x/filepath"
	"github.com/stretchr/testify/require"
)

func TestExpand_Tilde(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)

	require.Equal(t, filepath.Join(home, "config.yaml"), xfilepath.Expand("~/config.yaml"))
}

func TestExpand_BareTilde(t *testing.T) {
	home, err := os.UserHomeDir()
	require.NoError(t, err)

	require.Equal(t, home, xfilepath.Expand("~"))
}

func TestExpand_EnvVar(t *testing.T) {
	t.Setenv("TEST_EXPAND_DIR", "/opt/data")
	require.Equal(t, "/opt/data/file.txt", xfilepath.Expand("$TEST_EXPAND_DIR/file.txt"))
}

func TestExpand_Empty(t *testing.T) {
	require.Empty(t, xfilepath.Expand(""))
}

func TestExpand_NoExpansion(t *testing.T) {
	require.Equal(t, "/absolute/path", xfilepath.Expand("/absolute/path"))
}

func TestResolve(t *testing.T) {
	t.Parallel()

	dir, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	file := filepath.Join(dir, "f")
	link1 := filepath.Join(dir, "l1")
	link2 := filepath.Join(dir, "l2")
	require.NoError(t, os.WriteFile(file, []byte("x"), 0o600))
	require.NoError(t, os.Symlink(file, link1))
	require.NoError(t, os.Symlink(link1, link2))

	got, err := xfilepath.Resolve(link2)
	require.NoError(t, err)
	require.Equal(t, file, got)

	got, err = xfilepath.Resolve(link1)
	require.NoError(t, err)
	require.Equal(t, file, got)

	got, err = xfilepath.Resolve(file)
	require.NoError(t, err)
	require.Equal(t, file, got)

	missing := filepath.Join(dir, "missing")
	got, err = xfilepath.Resolve(missing)
	require.ErrorIs(t, err, ifs.ErrNotExist)
	require.Equal(t, missing, got)
}

func TestResolveLenient(t *testing.T) {
	t.Parallel()

	dir, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	realDir := filepath.Join(dir, "real")
	linkDir := filepath.Join(dir, "link")
	require.NoError(t, os.Mkdir(realDir, 0o755))
	require.NoError(t, os.Symlink(realDir, linkDir))

	got, err := xfilepath.ResolveLenient(filepath.Join(linkDir, "missing"))
	require.NoError(t, err)
	require.Equal(t, filepath.Join(realDir, "missing"), got)
}

func TestIsWithin(t *testing.T) {
	t.Parallel()

	require.True(t, xfilepath.IsWithin(".", "README.md"))
	require.True(t, xfilepath.IsWithin(".", "a/b.go", "c/d.go"))
	require.False(t, xfilepath.IsWithin("src", "lib/foo.go"))
	require.False(t, xfilepath.IsWithin("src"))

	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	require.NoError(t, os.Mkdir(sub, 0o755))

	require.True(t, xfilepath.IsWithin(dir, sub))
	require.True(t, xfilepath.IsWithin(dir, dir))
	require.False(t, xfilepath.IsWithin(sub, dir))
}

func TestMerge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"exact duplicates collapse", []string{"a", "a"}, []string{"a"}},
		{"nested under ancestor drops", []string{".", "./sub"}, []string{"."}},
		{"later ancestor subsumes earlier child", []string{"./sub", "."}, []string{"."}},
		{"child listed before its ancestor", []string{"a/b", "a"}, []string{"a"}},
		{"disjoint paths are kept", []string{"a", "b"}, []string{"a", "b"}},
		{
			"spellings of one path collapse to the first",
			[]string{"./foo", "foo"},
			[]string{"./foo"},
		},
		{"a sibling sharing a prefix is not a parent", []string{"a", "ab"}, []string{"a", "ab"}},
		{"empty input", nil, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, xfilepath.Merge(tt.in))
		})
	}
}

func TestMergeResolveSymlinks(t *testing.T) {
	t.Parallel()

	dir, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)
	target := filepath.Join(dir, "target")
	link := filepath.Join(dir, "link")
	require.NoError(t, os.Mkdir(target, 0o755))
	require.NoError(t, os.Symlink(target, link))

	// Lexically the symlink and its target are distinct, so the default keeps both.
	require.Equal(t, []string{target, link}, xfilepath.Merge([]string{target, link}))

	// Resolved, they are one location; the first-seen spelling survives.
	require.Equal(t, []string{target},
		xfilepath.Merge([]string{target, link}, xfilepath.WithResolveSymlinks()))
}
