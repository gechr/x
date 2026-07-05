package strings

// closestMaxDistanceDivisor bounds how far a candidate may sit from the target
// and still be offered: at most a third of the target's length in edits, so a
// long word tolerates more than a short one.
const closestMaxDistanceDivisor = 3

// Closest returns the candidate nearest to `target`, suitable for a "did you
// mean?" suggestion. Distance is the Damerau-Levenshtein (optimal string
// alignment) edit distance, so an adjacent transposition like "verfiy" counts
// as one edit, not two - the common typo plain Levenshtein over-penalizes. It
// returns "" when the nearest candidate is further than a third of `target`'s
// length in edits, so an unrelated word is never suggested. An empty `target`
// carries no signal and suggests nothing. Ties resolve to the first candidate.
//
//	Closest("verfiy", []string{"verify", "deep"}) // "verify"
//	Closest("xyzzy", []string{"verify", "deep"})  // ""
func Closest(target string, candidates []string) string {
	a := []rune(target)
	if len(a) == 0 {
		return ""
	}
	best, bestDist := "", 0
	for _, candidate := range candidates {
		d := damerauLevenshtein(a, []rune(candidate))
		if best == "" || d < bestDist {
			best, bestDist = candidate, d
		}
	}
	if best == "" || bestDist > max(len(a)/closestMaxDistanceDivisor, 1) {
		return ""
	}
	return best
}

// damerauLevenshtein returns the optimal-string-alignment distance between `a`
// and `b`: the fewest single-character insertions, deletions, substitutions, or
// adjacent transpositions to turn one into the other. It keeps three rows so a
// transposition can reach two rows back, for O(`len(a)*len(b)`) time and
// O(`len(b)`) space.
func damerauLevenshtein(a, b []rune) int {
	// Three rolling rows of the matrix: `prev1` is the previous row (`i-1`),
	// seeded as the base case - the cost of building `b`'s prefixes from an
	// empty `a`; `prev2` is the row before it (`i-2`), read only by a
	// transposition; `curr` is the row being filled.
	prev1 := make([]int, len(b)+1)
	prev2 := make([]int, len(b)+1)
	curr := make([]int, len(b)+1)
	for j := range prev1 {
		prev1[j] = j
	}
	for i := 1; i <= len(a); i++ {
		curr[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			curr[j] = min(prev1[j]+1, curr[j-1]+1, prev1[j-1]+cost)
			if i > 1 && j > 1 && a[i-1] == b[j-2] && a[i-2] == b[j-1] {
				curr[j] = min(curr[j], prev2[j-2]+1)
			}
		}
		prev2, prev1, curr = prev1, curr, prev2
	}
	return prev1[len(b)]
}
