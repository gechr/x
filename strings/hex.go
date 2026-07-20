package strings

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const sha256HexLength = 64

// HexEqual reports whether `a` and `b` denote the same hexadecimal value,
// ignoring surrounding whitespace, an optional `0x` (or `0X`) prefix, and case.
// Two blank strings are equal; a blank string never equals a non-blank one.
func HexEqual(a, b string) bool {
	return strings.EqualFold(
		trimHexPrefix(strings.TrimSpace(a)),
		trimHexPrefix(strings.TrimSpace(b)),
	)
}

// trimHexPrefix removes a leading `0x` or `0X` prefix from `s`, if present.
func trimHexPrefix(s string) string {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		return s[2:]
	}
	return s
}

// IsHex reports whether `s` is non-empty and consists entirely of hexadecimal
// digits. An empty string is not hex.
func IsHex(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if !IsHexChar(c) {
			return false
		}
	}
	return true
}

// IsHexChar reports whether `c` is a valid hexadecimal digit (0-9, a-f, A-F).
func IsHexChar(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// IsGitCommit reports whether `s` is 40 hexadecimal digits (a Git commit hash).
func IsGitCommit(s string) bool {
	return len(s) == 40 && IsHex(s)
}

// IsSHA256 reports whether `s` is 64 hexadecimal digits (a sha256 digest).
func IsSHA256(s string) bool {
	return len(s) == sha256HexLength && IsHex(s)
}

// DecodeSHA256 decodes a 64-digit hexadecimal sha256 digest.
// It returns the zero digest and an error if `s` has the wrong length or
// contains a non-hexadecimal character.
func DecodeSHA256(s string) ([32]byte, error) {
	var digest [32]byte
	if len(s) != sha256HexLength {
		return digest, fmt.Errorf(
			"invalid sha256 digest length: got %d, want %d",
			len(s), sha256HexLength,
		)
	}

	if _, err := hex.Decode(digest[:], []byte(s)); err != nil {
		return [32]byte{}, fmt.Errorf("invalid sha256 digest: %w", err)
	}
	return digest, nil
}
