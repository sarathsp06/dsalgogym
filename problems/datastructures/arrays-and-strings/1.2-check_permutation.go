package arrays

import "sort"

func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	counter := make(map[rune]int)
	for _, letter := range []rune(s1) {
		counter[letter] = counter[letter] + 1
	}
	for _, letter := range []rune(s2) {
		count := counter[letter]
		if count == 0 {
			return false
		}
		counter[letter] = count - 1
	}
	return true
}

// O(nLgn) : Sort each string and then compare
// 2 * O(nLgn) + O(n) = O(nLgn)
func isPermutationSort(s1, s2 string) bool {
	if !isSameLength(s1, s2) {
		return false
	}
	return sortString(s1) == sortString(s2)
}

// O(n) : Use maps to keep a character count
// Technically
// O(n) + O(n) + constant number of map keys = O(n)
// Alternatively, you can use an array instead of a map since there
// are only a fixed number of characters in the alphabet.
func isPermutationMap(s1, s2 string) bool {
	if !isSameLength(s1, s2) {
		return false
	}
	m := make(map[rune]int)
	// Get a count of each character in s1
	for _, c := range []rune(s1) {
		m[c]++
	}
	// Check if the counts match with s2
	for _, c := range []rune(s2) {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}
	// PS: I later learnt that this bit is redundant since both
	// strings have the same number of characters i.e
	// 1. We'd hit a non-zero value and deduct it
	// 2. We'd hit a zero value and trip the condition in the above loop
	// Leaving it here for posterity
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

// HELPERS
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func isSameLength(s1, s2 string) bool {
	return len(s1) == len(s2)
}
