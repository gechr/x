// Package filepath provides path helpers: symlink resolution and containment checks.
package filepath

import (
	"os"
	"path/filepath"
	"strings"
)

// Expand expands a leading ~ to the user's home directory and resolves
// environment variables via [os.ExpandEnv]. It is purely lexical: the result is
// not checked for existence or resolved against the filesystem (use [Resolve]
// or [ResolveLenient] for that).
func Expand(path string) string {
	if path == "" {
		return path
	}
	if path == "~" {
		if home, err := os.UserHomeDir(); err == nil {
			return home
		}
	}
	if rest, ok := strings.CutPrefix(path, "~/"); ok {
		if home, err := os.UserHomeDir(); err == nil {
			path = filepath.Join(home, rest)
		}
	}
	return os.ExpandEnv(path)
}

// LooksLikePath reports whether s begins with a marker that identifies it as a
// filesystem path rather than a bare name (a command, an "owner/repo" slug, or
// any other identifier). It is a purely lexical heuristic on the leading
// characters: it neither cleans s nor touches the filesystem, and an unrooted
// relative path such as "foo/bar" returns false because it is indistinguishable
// from a bare identifier.
//
// A path is recognised by a leading ".", "/", or "~" on every platform, plus a
// leading "\" (including a UNC "\\") or a drive-letter prefix ("C:") on Windows.
//
// Example:
//
//	LooksLikePath("./foo")      // true
//	LooksLikePath("~/cfg")      // true
//	LooksLikePath("/etc/app")   // true
//	LooksLikePath("owner/repo") // false
//	LooksLikePath("build")      // false
func LooksLikePath(s string) bool {
	if s == "" {
		return false
	}
	switch s[0] {
	case '.', '/', '~':
		return true
	}
	return looksLikePathOS(s)
}

// SplitPath splits a PATH-style list (such as $PATH or $GOPATH) on the
// OS-specific list separator ([os.PathListSeparator]), dropping the empty
// entries produced by leading, trailing, or doubled separators - an empty
// entry otherwise resolves to the current directory when joined. On Windows it
// honours the same quoting rules as [filepath.SplitList].
//
// Example:
//
//	SplitPath("/usr/bin:/bin:")   // ["/usr/bin", "/bin"]  (Unix)
//	SplitPath("")                 // []
func SplitPath(list string) []string {
	raw := filepath.SplitList(list)
	paths := make([]string, 0, len(raw))
	for _, path := range raw {
		if path != "" {
			paths = append(paths, path)
		}
	}
	return paths
}

// Rebase anchors a relative `path` to the base directory `dir`, returning
// `filepath.Join(dir, path)`. An empty or already-absolute `path` is returned
// unchanged, so a caller-supplied override always wins over the default base.
// It is purely lexical and touches neither the filesystem nor the current
// working directory; the result is absolute only when `dir` is.
//
// Example:
//
//	Rebase("/etc/app", "conf.d")   // "/etc/app/conf.d"
//	Rebase("/etc/app", "/abs.cfg") // "/abs.cfg"
//	Rebase("/etc/app", "")         // ""
func Rebase(dir, path string) string {
	if path == "" || filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(dir, path)
}

// Resolve recursively follows every symlink along `path` and returns the fully
// resolved absolute path. On any error (missing component, cycle, permission)
// the input path is returned alongside the error so callers can choose whether
// to handle it or fall back.
func Resolve(path string) (string, error) {
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path, err
	}
	abs, err := filepath.Abs(resolved)
	if err != nil {
		return resolved, err
	}
	return abs, nil
}

// ResolveLenient returns an absolute path with symlinks resolved where
// possible. If `path` itself cannot be resolved, it resolves the parent directory
// and rejoins the original base name. If neither can be resolved, it returns
// the absolute path.
func ResolveLenient(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return path, err
	}
	resolved, err := Resolve(abs)
	if err == nil {
		return resolved, nil
	}
	parent, err := Resolve(filepath.Dir(abs))
	if err == nil {
		return filepath.Join(parent, filepath.Base(abs)), nil
	}
	return abs, nil
}

// IsWithin reports whether all target paths are equal to or contained within
// `base`. Returns false when no `targets` are provided.
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
	absBase, err := filepath.Abs(base)
	if err != nil {
		return false
	}
	for _, target := range targets {
		absTarget, err := filepath.Abs(target)
		if err != nil {
			return false
		}
		if !contains(absBase, absTarget) {
			return false
		}
	}
	return true
}

// contains reports whether `inner` is equal to or nested under `outer`. Both must be
// absolute, cleaned paths. The separator appended to the prefix stops `a` from
// matching a sibling `ab`.
func contains(outer, inner string) bool {
	if equalPath(inner, outer) {
		return true
	}
	prefix := outer
	if !strings.HasSuffix(prefix, string(filepath.Separator)) {
		prefix += string(filepath.Separator)
	}
	return hasPathPrefix(inner, prefix)
}

// MergeOption configures [Merge].
type MergeOption func(*mergeConfig)

type mergeConfig struct {
	resolveSymlinks bool
}

// WithResolveSymlinks makes [Merge] compare paths by their resolved physical
// location (via [ResolveLenient]) rather than lexically, so two spellings that
// reach the same target through a symlink are merged. It touches the filesystem;
// without it [Merge] is pure and lexical.
func WithResolveSymlinks() MergeOption {
	return func(c *mergeConfig) { c.resolveSymlinks = true }
}

// Merge reduces `paths` to the minimal set covering the same locations: comparing
// them as cleaned absolute paths, it drops any that duplicate or are nested
// within another, so a later walk visits each file once. Survivors keep their
// original form and first-seen order; a path whose absolute form cannot be
// computed is compared by its cleaned form.
//
// The comparison is lexical by default; pass [WithResolveSymlinks] to compare
// resolved physical locations instead.
//
// Example:
//
//	Merge([]string{"a", "a"})     // ["a"]
//	Merge([]string{".", "./sub"}) // ["."]
//	Merge([]string{"a/b", "a"})   // ["a"]
//	Merge([]string{"a", "b"})     // ["a", "b"]
func Merge(paths []string, opts ...MergeOption) []string {
	var cfg mergeConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	keys := make([]string, len(paths))
	for i, path := range paths {
		keys[i] = cfg.mergeKey(path)
	}

	merged := make([]string, 0, len(paths))
	for i := range paths {
		if !subsumed(keys, i) {
			merged = append(merged, paths[i])
		}
	}
	return merged
}

// mergeKey returns the comparison key for a path: its resolved physical location
// when [WithResolveSymlinks] is set, otherwise its cleaned absolute (lexical) form,
// falling back to the cleaned form when an absolute path cannot be computed.
func (c mergeConfig) mergeKey(path string) string {
	if c.resolveSymlinks {
		resolved, _ := ResolveLenient(path)
		return resolved
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return filepath.Clean(path)
	}
	return abs
}

// subsumed reports whether `keys[i]` is covered by another entry: a strict
// ancestor, or - for an exact duplicate - an earlier occurrence (so the first of
// a set of equal paths survives).
func subsumed(keys []string, i int) bool {
	for j := range keys {
		if i == j || !contains(keys[j], keys[i]) {
			continue
		}
		if equalPath(keys[i], keys[j]) {
			if j < i {
				return true
			}
			continue
		}
		return true
	}
	return false
}
