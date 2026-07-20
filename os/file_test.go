package os_test

import (
	"os"
	"path/filepath"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestSameFile(t *testing.T) {
	t.Parallel()

	dir, err := filepath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)

	file := filepath.Join(dir, "file")
	require.NoError(t, os.WriteFile(file, []byte("payload"), 0o600))

	link := filepath.Join(dir, "link")
	require.NoError(t, os.Symlink(file, link))

	realDir := filepath.Join(dir, "real")
	linkDir := filepath.Join(dir, "link-dir")
	require.NoError(t, os.Mkdir(realDir, 0o755))
	require.NoError(t, os.Symlink(realDir, linkDir))

	hardlink := filepath.Join(dir, "hardlink")
	require.NoError(t, os.Link(file, hardlink))

	other := filepath.Join(dir, "other")
	require.NoError(t, os.WriteFile(other, []byte("payload"), 0o600))

	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "identical path",
			a:    file,
			b:    file,
			want: true,
		},
		{
			name: "symlink to same target",
			a:    link,
			b:    file,
			want: true,
		},
		{
			name: "symlinked ancestor with missing leaf",
			a:    filepath.Join(linkDir, "missing"),
			b:    filepath.Join(realDir, "missing"),
			want: true,
		},
		{
			name: "hardlink",
			a:    hardlink,
			b:    file,
			want: true,
		},
		{
			name: "different files",
			a:    file,
			b:    other,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := xos.SameFile(tt.a, tt.b)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestRemoveIfExists(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	file := filepath.Join(dir, "file")
	require.NoError(t, os.WriteFile(file, []byte("payload"), 0o600))

	require.NoError(t, xos.RemoveIfExists(file))
	require.NoFileExists(t, file)
	require.NoError(t, xos.RemoveIfExists(file))

	// A path below a regular file does not exist (ENOTDIR).
	require.NoError(t, os.WriteFile(file, []byte("payload"), 0o600))
	require.NoError(t, xos.RemoveIfExists(filepath.Join(file, "child")))

	nonEmptyDir := filepath.Join(dir, "non-empty")
	require.NoError(t, os.Mkdir(nonEmptyDir, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(nonEmptyDir, "child"), nil, 0o600))
	require.Error(t, xos.RemoveIfExists(nonEmptyDir))
}
