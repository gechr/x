package shell

import (
	"fmt"
	"strings"
)

type splitState int

const (
	splitStart splitState = iota
	splitWord
	splitEscape
	splitSingleQuote
	splitDoubleQuote
	splitDoubleQuoteEscape
	splitComment
)

// Split partitions s into shell-style words. Whitespace separates words,
// quotes preserve whitespace, backslashes escape the following rune, and a "#"
// starts a comment when it appears where a new word could start.
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
				state = splitEscape
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
		case splitEscape:
			word.WriteRune(r)
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
			word.WriteRune(r)
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
	case splitEscape, splitDoubleQuoteEscape:
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
