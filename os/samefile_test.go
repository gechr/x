package os_test

import (
	stdos "os"
	stdpath "path/filepath"
	"testing"

	xos "github.com/gechr/x/os"
	"github.com/stretchr/testify/require"
)

func TestSameFile(t *testing.T) {
	t.Parallel()

	dir, err := stdpath.EvalSymlinks(t.TempDir())
	require.NoError(t, err)

	file := stdpath.Join(dir, "file")
	require.NoError(t, stdos.WriteFile(file, []byte("payload"), 0o600))

	link := stdpath.Join(dir, "link")
	require.NoError(t, stdos.Symlink(file, link))

	realDir := stdpath.Join(dir, "real")
	linkDir := stdpath.Join(dir, "link-dir")
	require.NoError(t, stdos.Mkdir(realDir, 0o755))
	require.NoError(t, stdos.Symlink(realDir, linkDir))

	hardlink := stdpath.Join(dir, "hardlink")
	require.NoError(t, stdos.Link(file, hardlink))

	other := stdpath.Join(dir, "other")
	require.NoError(t, stdos.WriteFile(other, []byte("payload"), 0o600))

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
			a:    stdpath.Join(linkDir, "missing"),
			b:    stdpath.Join(realDir, "missing"),
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
