package filepath_test

import (
	"fmt"
	"os"
	"path/filepath"

	xfilepath "github.com/gechr/x/filepath"
)

// Expand also expands a leading ~ to the user's home directory.
func ExampleExpand() {
	_ = os.Setenv("PROJECT_ROOT", "/srv/app")
	fmt.Println(xfilepath.Expand("$PROJECT_ROOT/config.toml"))
	// Output:
	// /srv/app/config.toml
}

// SplitPath splits a PATH-style list and drops the empty entry left by the
// trailing separator.
func ExampleSplitPath() {
	sep := string(os.PathListSeparator)
	fmt.Println(xfilepath.SplitPath("/usr/bin" + sep + "/bin" + sep))
	// Output:
	// [/usr/bin /bin]
}

// Resolve follows symlinks, so a link and its target resolve to the same
// path.
func ExampleResolve() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	target := filepath.Join(dir, "target.txt")
	_ = os.WriteFile(target, []byte("hello"), 0o600)
	link := filepath.Join(dir, "link.txt")
	_ = os.Symlink(target, link)

	resolvedLink, _ := xfilepath.Resolve(link)
	resolvedTarget, _ := xfilepath.Resolve(target)
	fmt.Println(resolvedLink == resolvedTarget)
	// Output:
	// true
}

// ResolveLenient succeeds where Resolve fails: a missing file resolves via
// its parent directory, keeping the original base name.
func ExampleResolveLenient() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	missing := filepath.Join(dir, "missing.txt")
	_, err := xfilepath.Resolve(missing)
	fmt.Println(err != nil)

	resolved, _ := xfilepath.ResolveLenient(missing)
	fmt.Println(filepath.Base(resolved))
	// Output:
	// true
	// missing.txt
}

// Rebase anchors a relative path to a base directory. (ToSlash keeps the
// output identical across operating systems.)
func ExampleRebase() {
	base := filepath.Join("etc", "app")
	fmt.Println(filepath.ToSlash(xfilepath.Rebase(base, "conf.d")))
	// Output:
	// etc/app/conf.d
}

func ExampleIsWithin() {
	fmt.Println(xfilepath.IsWithin("src", "src/foo.go"))
	fmt.Println(xfilepath.IsWithin("src", "src"))
	fmt.Println(xfilepath.IsWithin("src", "lib/foo.go"))
	fmt.Println(xfilepath.IsWithin("src"))
	// Output:
	// true
	// true
	// false
	// false
}

// IsWithin only reports `true` when every target is contained within the base.
func ExampleIsWithin_multipleTargets() {
	fmt.Println(xfilepath.IsWithin(".", "src/foo.go", "lib/bar.go"))
	fmt.Println(xfilepath.IsWithin("src", "src/foo.go", "lib/bar.go"))
	// Output:
	// true
	// false
}

func ExampleMerge() {
	fmt.Println(xfilepath.Merge([]string{".", "./sub"}))
	fmt.Println(xfilepath.Merge([]string{"a/b", "a"}))
	fmt.Println(xfilepath.Merge([]string{"a", "b"}))
	// Output:
	// [.]
	// [a]
	// [a b]
}

// Exact duplicates are merged; the first occurrence survives in its
// original spelling.
func ExampleMerge_duplicates() {
	fmt.Println(xfilepath.Merge([]string{"a", "./a", "a/"}))
	// Output:
	// [a]
}
