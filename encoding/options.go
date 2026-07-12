package encoding

// RenderOption configures [Path.Render].
type RenderOption func(*renderConfig)

type renderConfig struct {
	marker rune
	root   bool
}

// WithRoot prefixes the rendered path with a root marker: `$` by default
// (`$.items[0]`), or `marker` if given (`WithRoot('@')` renders `@.items[0]`).
func WithRoot(marker ...rune) RenderOption {
	return func(cfg *renderConfig) {
		cfg.root = true
		cfg.marker = '$'
		if len(marker) > 0 {
			cfg.marker = marker[0]
		}
	}
}
