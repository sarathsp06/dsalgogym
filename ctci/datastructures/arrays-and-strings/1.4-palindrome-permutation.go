package arrays

import (
	"unicode"
)

// Complexity: O(n)
func isPalindromePermutation(input string) bool {
	// Even Palindrome: count(c) % 2 == 0 for all c in input
	// Odd Palindrome: count(c) % 2 == 0 for all c in input except for one character
	// Ignore spaces
	count := make(map[rune]int)
	for _, c := range []rune(input) {
		count[unicode.ToLower(c)]++
	}
	// We can have only one character with odd count (excluding space) in an odd palindrome
	foundOddRune := false
	for k, v := range count {
		if v%2 != 0 && !unicode.IsSpace(k) {
			if foundOddRune {
				return false
			}
			foundOddRune = true
		}
	}
	return true
}
