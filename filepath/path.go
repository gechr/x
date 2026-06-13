package filepath

import (
	stdpath "path/filepath"
	"strings"
)

// Resolve recursively follows every symlink along path and returns the fully
// resolved absolute path. On any error (missing component, cycle, permission)
// the input path is returned alongside the error so callers can choose whether
// to handle it or fall back.
func Resolve(path string) (string, error) {
	resolved, err := stdpath.EvalSymlinks(path)
	if err != nil {
		return path, err
	}
	abs, err := stdpath.Abs(resolved)
	if err != nil {
		return resolved, err
	}
	return abs, nil
}

// ResolveLenient returns an absolute path with symlinks resolved where
// possible. If path itself cannot be resolved, it resolves the parent directory
// and rejoins the original base name. If neither can be resolved, it returns
// the absolute path.
func ResolveLenient(path string) (string, error) {
	abs, err := stdpath.Abs(path)
	if err != nil {
		return path, err
	}
	resolved, err := Resolve(abs)
	if err == nil {
		return resolved, nil
	}
	parent, err := Resolve(stdpath.Dir(abs))
	if err == nil {
		return stdpath.Join(parent, stdpath.Base(abs)), nil
	}
	return abs, nil
}

// IsWithin reports whether all target paths are equal to or contained within
// base. Returns false when no targets are provided.
//
// Example:
//
//	IsWithin("src", "src/foo.go")             // true
//	IsWithin(".", "src/foo.go", "lib/bar.go") // true
//	IsWithin("src", "lib/foo.go")             // false
func IsWithin(base string, targets ...string) bool {
	if len(targets) == 0 {
		return false
	}
	absBase, err := stdpath.Abs(base)
	if err != nil {
		return false
	}
	prefix := absBase
	if !strings.HasSuffix(prefix, string(stdpath.Separator)) {
		prefix += string(stdpath.Separator)
	}
	for _, target := range targets {
		absTarget, err := stdpath.Abs(target)
		if err != nil {
			return false
		}
		if !equalPath(absTarget, absBase) && !hasPathPrefix(absTarget, prefix) {
			return false
		}
	}
	return true
}
