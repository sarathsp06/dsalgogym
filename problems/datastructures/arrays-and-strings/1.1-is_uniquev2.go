package arrays

import (
	"sort"
	"strings"
)

// Complexity: O(n)
// Using maps
func isUniqueSimple(input string) bool {
	m := make(map[rune]bool)
	for _, c := range strings.ToLower(input) {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}

// Complexity: O(n^2)
func isUniqueBrute(input string) bool {
	for i, c1 := range strings.ToLower(input) {
		for _, c2 := range strings.ToLower(input[i+1:]) {
			if c1 == c2 {
				return false
			}
		}
	}
	return true
}

// Complexity: O(nLogn)
func isUniqueSorted(input string) bool {
	s := []rune(strings.ToLower(input))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	return true
}

// Assumption: input consists of english alphabets
// Complexity: O(n)
func isUniqueOptimized(input string) bool {
	var bitMap int64
	// We create a bitmap with 26 slots, each corresponding to a letter in the
	// alphabet
	bitMap = 1 << 26
	for _, c := range strings.ToLower(input) {
		var mask int64
		bitIndex := uint(rune(c) - rune('a'))
		// A mask is created for each letter and then OR-ed with the
		// original bitmask to see if it has been encountered before
		mask = 1 << bitIndex
		if bitMap&mask == mask {
			return false
		}
		bitMap = bitMap | mask
	}
	return true
}
