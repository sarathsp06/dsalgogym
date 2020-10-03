package arrays

import (
	"sort"
	"strings"
)

/**
* Question:
* Implement an algorithm to determine if a string has all unique characters.  What if you can not use additional data struyctures.
 */

// Unique checks if the string has all unique characters
func Unique(str string) bool {
	charsFound := make(map[rune]bool)
	for _, char := range []rune(str) {
		if found := charsFound[char]; found {
			return false
		}
		charsFound[char] = true
	}
	return true
}

// UniqueNoExtra checks if the string has all unique characters
// without using an extra data structure
func UniqueNoExtra(str string) bool {
	if len(str) == 0 {
		return true
	}
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	prev := bytes[0]
	for _, j := range bytes[1:] {
		if prev == j {
			return false
		}
		prev = j
	}
	return true
}

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
