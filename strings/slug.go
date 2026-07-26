package strings

import (
	"slices"
	"strings"
)

// defaultSlugSeparator is the only separator [IsSlug] admits, and so the one
// [Slug] joins with unless [WithSlugSeparator] overrides it.
const defaultSlugSeparator = '-'

// SlugOption shapes what [Slug] produces. The options compose in any order, and
// all but one are bounds - they decide how much of the slug is kept, so the
// result still satisfies [IsSlug] (or is empty, when a bound leaves nothing).
// [WithSlugSeparator] is the exception, changing the grammar itself.
type SlugOption func(*slugConfig)

// slugConfig is the bounds a Slug call was given. A zero value is the unbounded
// default, and each field is inert until set to a positive value - the standard
// option semantics, so a caller computing a bound that comes out 0 gets the
// unbounded slug rather than an empty one.
type slugConfig struct {
	maxLength     int
	maxWords      int
	midWord       bool
	minWordLength int
	separator     rune
}

// joiner returns the separator the words are joined with: the override, or
// [defaultSlugSeparator] when none was given (a zero rune is no override, since
// NUL is not a separator anyone means).
func (c slugConfig) joiner() string {
	if c.separator == 0 {
		return string(defaultSlugSeparator)
	}
	return string(c.separator)
}

// WithSlugMaxLength caps the slug at `n` characters, cut at a word boundary: a
// cut landing inside a word takes that whole word with it, so the result reads as
// the words that were kept rather than one that was chopped.
//
//	Slug("My Long Service Name", WithSlugMaxLength(12)) // "my-long", not "my-long-serv"
//
// A first word longer than `n` is the one exception, having no boundary to fall
// back to: it is cut mid-word rather than yielding nothing.
//
//	Slug("verylongword", WithSlugMaxLength(4)) // "very"
//
// `n` counts bytes, which a word makes identical to runes - words are ASCII -
// unless a multi-byte [WithSlugSeparator] joins them. Bytes are what a cap is
// usually spent against (a DNS label, a filename, a 63-character field), so a
// rune count would be the one thing a cap must never do: overrun.
func WithSlugMaxLength(n int) SlugOption {
	return func(c *slugConfig) { c.maxLength = n }
}

// WithSlugSeparator joins the words with `sep` instead of '-'.
//
//	Slug("My Service", WithSlugSeparator('_')) // "my_service"
//
// It is the one option that changes what the result is rather than how much of it
// there is, so it is also the one that forfeits the [IsSlug] guarantee: only '-'
// is a slug separator, and '_' yields a string [IsSlugLenient] accepts and
// [IsSlug] rejects. Anything else - a '.', a '/', a non-ASCII rune - satisfies
// neither, and the collapsing, trimming, and bounds still hold. Pick it when the
// identifier's grammar is the caller's, not this package's.
func WithSlugSeparator(sep rune) SlugOption {
	return func(c *slugConfig) { c.separator = sep }
}

// WithSlugCutMidWord lets [WithSlugMaxLength] cut inside a word, filling the cap
// instead of falling back to the last word boundary. The separator is still never
// left dangling, so the result stays a valid slug - it just ends in a fragment.
//
//	Slug("My Long Service Name", WithSlugMaxLength(12))                       // "my-long"
//	Slug("My Long Service Name", WithSlugMaxLength(12), WithSlugCutMidWord()) // "my-long-serv"
//
// Use it where the slug is an identifier that should use the space it has (a
// generated name, a truncated key), not where a human reads it as words. It is
// inert without a length cap, which is the only bound that cuts.
func WithSlugCutMidWord() SlugOption {
	return func(c *slugConfig) { c.midWord = true }
}

// WithSlugMaxWords keeps at most the first `n` words, dropping the rest.
//
//	Slug("My Long Service Name", WithSlugMaxWords(2)) // "my-long"
func WithSlugMaxWords(n int) SlugOption {
	return func(c *slugConfig) { c.maxWords = n }
}

// WithSlugMinWordLength drops every word shorter than `n` characters, which is
// how the noise a punctuation split leaves behind is kept out of the slug -
// initials, stray digits, the `s` of a possessive.
//
//	Slug("Bob's API v2 Service", WithSlugMinWordLength(2)) // "bob-api-v2-service"
//
// It is applied before the other bounds, so a dropped word does not consume a
// [WithSlugMaxWords] slot or [WithSlugMaxLength] characters.
func WithSlugMinWordLength(n int) SlugOption {
	return func(c *slugConfig) { c.minWordLength = n }
}

// Slug converts `s` to a slug: lowercased, with every run of characters that is
// not an ASCII alphanumeric collapsed to a single '-', and no leading or
// trailing separator. `My Service`, `my_service`, and `  my.service!!  ` all
// slugify to `my-service`.
//
// The result satisfies [IsSlug] except for the one input that has no slug: a
// string carrying no ASCII alphanumeric at all (`""`, `"___"`, `"日本"`) returns
// the empty string, which [IsSlug] rejects. Test the result rather than the
// input when that case must be caught. [WithSlugSeparator] is the only option
// that forfeits the guarantee; the bounds all preserve it.
//
// Non-ASCII characters are separators rather than letters, since a slug admits
// only a-z, 0-9, and '-'; transliterating them is the caller's job. Case is the
// only thing folded, so `myService` slugifies to `myservice` - a slug reflects
// the characters it was given, not the word boundaries a reader infers from
// them.
//
// Nothing is bounded by default - no length cap, no word count, matching
// [IsSlug], which admits a slug of any shape - since a bound is a policy the
// caller cannot undo. Pass [WithSlugMaxLength], [WithSlugMaxWords],
// [WithSlugMinWordLength], or [WithSlugCutMidWord] to impose one, and
// [WithSlugSeparator] to join with something other than '-'.
//
// Slugification is idempotent: every valid slug is its own slug, and stays its
// own slug under any bound it already satisfies.
func Slug(s string, opts ...SlugOption) string {
	var cfg slugConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	words := slugWords(s)
	if cfg.minWordLength > 0 {
		words = slices.DeleteFunc(words, func(w string) bool { return len(w) < cfg.minWordLength })
	}
	if cfg.maxWords > 0 && len(words) > cfg.maxWords {
		words = words[:cfg.maxWords]
	}
	separator := cfg.joiner()
	slug := strings.Join(words, separator)
	if cfg.maxLength <= 0 || len(slug) <= cfg.maxLength {
		return slug
	}
	return joinSlugWithin(words, cfg.maxLength, separator, cfg.midWord)
}

// slugWords splits `s` into the lowercase alphanumeric runs a slug is built from,
// which is where every separator question is settled: a run of non-alphanumerics
// is one boundary however it is spelled, and a leading or trailing one yields no
// word at all rather than an empty first or last element.
func slugWords(s string) []string {
	var words []string
	var b strings.Builder
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
		case c >= 'A' && c <= 'Z':
			c = c - 'A' + 'a'
		default:
			if b.Len() > 0 {
				words = append(words, b.String())
				b.Reset()
			}
			continue
		}
		b.WriteRune(c)
	}
	if b.Len() > 0 {
		words = append(words, b.String())
	}
	return words
}

// joinSlugWithin joins as many words as fit in `limit` bytes, separators
// included. A word that would overrun the cap ends the slug, so the cut lands on
// a word boundary; midWord instead spends what is left of the cap on the front of
// that word. A first word that cannot fit at all is cut either way, in preference
// to returning nothing.
//
// The cut is made while joining rather than on the joined string, which is what
// keeps it structural: the separator is written whole or not at all, so no partial
// rune can be emitted, and no trailing separator can appear that would have to be
// trimmed back off. Trimming would be wrong as well as redundant, since a word
// character may be the separator rune itself (`Slug("my tax b",
// WithSlugSeparator('x'), …)`) and trimming cannot tell the two apart.
func joinSlugWithin(words []string, limit int, separator string, midWord bool) string {
	var b strings.Builder
	b.Grow(limit)
	for i, word := range words {
		length := len(word)
		if i > 0 {
			length += len(separator) // the separator this word is joined with
		}
		if b.Len()+length <= limit {
			if i > 0 {
				b.WriteString(separator)
			}
			b.WriteString(word)
			continue
		}

		// The word overruns the cap, so this is the last one either way.
		if !midWord {
			if b.Len() == 0 {
				b.WriteString(word[:limit]) // the first word alone cannot fit
			}
			break
		}
		// A word is ASCII, so what remains of the cap is a rune-safe cut of it -
		// but only once the separator has been paid for in full.
		remaining := limit - b.Len()
		if i > 0 {
			remaining -= len(separator)
			if remaining <= 0 {
				break
			}
			b.WriteString(separator)
		}
		b.WriteString(word[:remaining])
		break
	}
	return b.String()
}

// IsSlug reports whether `s` is a valid slug: a non-empty, URL-friendly
// identifier of lowercase alphanumerics and '-', starting and ending with an
// alphanumeric (e.g. `my-service`). Underscores are not permitted; `-` is the
// only allowed separator, and it may not appear consecutively. Every valid slug
// is therefore a fixed point of slugification. An empty string is not a slug.
func IsSlug(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			// Always allowed.
		case c == '-':
			// Not allowed at the start, end, or after another '-'.
			if i == 0 || i == len(s)-1 || s[i-1] == '-' {
				return false
			}
		default:
			return false
		}
	}
	return true
}

// IsSlugLenient reports whether `s` is a valid lenient slug: a non-empty
// identifier of lowercase alphanumerics, '-', and '_', starting and ending with
// an alphanumeric (e.g. `my-service`, `my_service`, `a--b__c`). Unlike
// [IsSlug], underscores are permitted and separators may appear consecutively
// or mixed; only leading and trailing separators are rejected. An empty string
// is not a slug.
func IsSlugLenient(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		switch {
		case c >= 'a' && c <= 'z', c >= '0' && c <= '9':
			// Always allowed.
		case c == '-', c == '_':
			// Allowed internally, including consecutively and mixed, but not at
			// the start or end.
			if i == 0 || i == len(s)-1 {
				return false
			}
		default:
			return false
		}
	}
	return true
}
