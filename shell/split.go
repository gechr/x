package shell

import (
	"fmt"
	"strings"
)

type splitState int

const (
	splitStart splitState = iota
	splitStartEscape
	splitWord
	splitEscape
	splitSingleQuote
	splitDoubleQuote
	splitDoubleQuoteEscape
	splitComment
)

// Split partitions `s` into shell-style words. Whitespace separates words,
// quotes preserve whitespace, backslashes escape the following rune, a
// backslash-newline pair is removed as a line continuation, and a "#" starts a
// comment when it appears where a new word could start. Inside double quotes,
// a backslash is special only before '$', '`', '"', '\', or a newline; before
// any other rune it is kept literally, following POSIX.
func Split(s string) ([]string, error) {
	var words []string
	var word strings.Builder
	state := splitStart

	emit := func() {
		words = append(words, word.String())
		word.Reset()
	}

	for _, r := range s {
		switch state {
		case splitStart:
			switch {
			case isSplitSpace(r):
			case r == '#':
				state = splitComment
			case r == '\\':
				state = splitStartEscape
			case r == '\'':
				state = splitSingleQuote
			case r == '"':
				state = splitDoubleQuote
			default:
				word.WriteRune(r)
				state = splitWord
			}
		case splitWord:
			switch {
			case isSplitSpace(r):
				emit()
				state = splitStart
			case r == '\\':
				state = splitEscape
			case r == '\'':
				state = splitSingleQuote
			case r == '"':
				state = splitDoubleQuote
			default:
				word.WriteRune(r)
			}
		case splitStartEscape:
			if r == '\n' {
				state = splitStart
			} else {
				word.WriteRune(r)
				state = splitWord
			}
		case splitEscape:
			if r != '\n' {
				word.WriteRune(r)
			}
			state = splitWord
		case splitSingleQuote:
			if r == '\'' {
				state = splitWord
			} else {
				word.WriteRune(r)
			}
		case splitDoubleQuote:
			switch r {
			case '\\':
				state = splitDoubleQuoteEscape
			case '"':
				state = splitWord
			default:
				word.WriteRune(r)
			}
		case splitDoubleQuoteEscape:
			switch r {
			case '$', '`', '"', '\\':
				word.WriteRune(r)
			case '\n': // line continuation: both runes are dropped
			default:
				word.WriteRune('\\')
				word.WriteRune(r)
			}
			state = splitDoubleQuote
		case splitComment:
			if r == '\n' {
				state = splitStart
			}
		default:
			return words, fmt.Errorf("unexpected shell split state: %d", state)
		}
	}

	switch state {
	case splitStart, splitComment:
		return words, nil
	case splitWord:
		emit()
		return words, nil
	case splitStartEscape, splitEscape, splitDoubleQuoteEscape:
		return words, fmt.Errorf("EOF found after escape character")
	case splitSingleQuote, splitDoubleQuote:
		return words, fmt.Errorf("EOF found when expecting closing quote")
	default:
		return words, fmt.Errorf("unexpected shell split state: %d", state)
	}
}

func isSplitSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}
