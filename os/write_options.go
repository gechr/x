package os

// AtomicWriteOption configures [AtomicWrite].
type AtomicWriteOption func(*atomicWriteConfig)

type atomicWriteConfig struct {
	noFollow bool
	noSync   bool
}

// WithNoFollow makes [AtomicWrite] replace a symlink at `path` instead of
// writing through it.
func WithNoFollow() AtomicWriteOption {
	return func(c *atomicWriteConfig) { c.noFollow = true }
}

// WithNoSync makes [AtomicWrite] skip both syncs: the write stays atomic but may
// not survive a power loss. Syncing is slow on macOS (F_FULLFSYNC), so this suits
// bulk rewrites such as formatting a tree.
func WithNoSync() AtomicWriteOption {
	return func(c *atomicWriteConfig) { c.noSync = true }
}
