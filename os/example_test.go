package os_test

import (
	"fmt"
	"os"
	"path/filepath"

	xos "github.com/gechr/x/os"
)

func ExampleExists() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)

	exists, _ := xos.Exists(path)
	missing, _ := xos.Exists(filepath.Join(dir, "missing.txt"))
	fmt.Println(exists)
	fmt.Println(missing)
	// Output:
	// true
	// false
}

func ExampleIsFile() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)

	file, _ := xos.IsFile(path)
	notFile, _ := xos.IsFile(dir)
	fmt.Println(file)
	fmt.Println(notFile)
	// Output:
	// true
	// false
}

func ExampleIsDir() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)

	isDir, _ := xos.IsDir(dir)
	notDir, _ := xos.IsDir(path)
	fmt.Println(isDir)
	fmt.Println(notDir)
	// Output:
	// true
	// false
}

// EnsureDir creates missing parent directories, like mkdir -p.
func ExampleEnsureDir() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	nested := filepath.Join(dir, "a", "b", "c")
	if err := xos.EnsureDir(nested, 0o755); err != nil {
		fmt.Println(err)
		return
	}

	isDir, _ := xos.IsDir(nested)
	fmt.Println(isDir)
	// Output:
	// true
}

// EnsureFile creates the file and any missing parent directories.
func ExampleEnsureFile() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "a", "b", "config.txt")
	if err := xos.EnsureFile(path, 0o600); err != nil {
		fmt.Println(err)
		return
	}

	isFile, _ := xos.IsFile(path)
	fmt.Println(isFile)
	// Output:
	// true
}

func ExampleIsSymlink() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)
	link := filepath.Join(dir, "link.txt")
	_ = os.Symlink(path, link)

	isLink, _ := xos.IsSymlink(link)
	notLink, _ := xos.IsSymlink(path)
	fmt.Println(isLink)
	fmt.Println(notLink)
	// Output:
	// true
	// false
}

// A missing path is not a writable directory.
func ExampleIsWritableDir() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	fmt.Println(xos.IsWritableDir(dir))
	fmt.Println(xos.IsWritableDir(filepath.Join(dir, "missing")))
	// Output:
	// true
	// false
}

func ExampleSameFile() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)
	link := filepath.Join(dir, "link.txt")
	_ = os.Link(path, link)
	other := filepath.Join(dir, "other.txt")
	_ = os.WriteFile(other, []byte("hello"), 0o600)

	same, _ := xos.SameFile(path, link)
	different, _ := xos.SameFile(path, other)
	fmt.Println(same)
	fmt.Println(different)
	// Output:
	// true
	// false
}

func ExampleRemoveIfExists() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "file.txt")
	_ = os.WriteFile(path, []byte("hello"), 0o600)

	fmt.Println(xos.RemoveIfExists(path))
	fmt.Println(xos.RemoveIfExists(path))
	// Output:
	// <nil>
	// <nil>
}

func ExampleAtomicWrite() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "config.txt")
	if err := xos.AtomicWrite(path, []byte("hello\n"), 0o600); err != nil {
		fmt.Println(err)
		return
	}

	data, _ := os.ReadFile(path)
	fmt.Printf("%s", data)
	// Output:
	// hello
}

func ExampleCopyFile() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "dst.txt")
	_ = os.WriteFile(src, []byte("hello\n"), 0o600)

	if err := xos.CopyFile(src, dst); err != nil {
		fmt.Println(err)
		return
	}

	data, _ := os.ReadFile(dst)
	fmt.Printf("%s", data)
	// Output:
	// hello
}

// ReadLines drops blank lines and trims surrounding whitespace.
func ExampleReadLines() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "lines.txt")
	_ = os.WriteFile(path, []byte("  alpha  \n\n\tbeta\n\ngamma\n"), 0o600)

	lines, _ := xos.ReadLines(path)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// alpha
	// beta
	// gamma
}

func ExampleWriteLines() {
	dir, _ := os.MkdirTemp("", "example")
	defer func() { _ = os.RemoveAll(dir) }()

	path := filepath.Join(dir, "lines.txt")
	if err := xos.WriteLines(path, []string{"alpha", "beta"}, 0o600); err != nil {
		fmt.Println(err)
		return
	}

	data, _ := os.ReadFile(path)
	fmt.Printf("%q\n", data)
	// Output:
	// "alpha\nbeta\n"
}
