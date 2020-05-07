package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromePermutation(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
	}{
		{"Empty String", "", true},
		{"Single Character", "a", true},
		{"Even Palindrome", "aa", true},
		{"Odd Palindrome", "aba", true},
		{"Even Palindrom with spaces", "a a", true},
		{"Odd Palindrome with spaces", "a ba ", true},
		{"Shuffled Even Palindrome", "abab", true},
		{"Shuffled Odd Palindrome", "baa", true},
		{"Shuffled Even Palindrome with spaces", "             a bab", true},
		{"Shuffled Odd Palindrome with spaces", "ba a", true},
		{"Shuffled Odd Palindrome with spaces", "Taco cat", true},
		{"Not palindrome", "abasdf", false},
		{"Not palindrome with spaces", "abasdf        lasldxf", false},
		{"Not palindrome with spaces", "aba      x", false},
	}
	testFuncs := []func(string) bool{isPalindromePermutation}
	for _, fn := range testFuncs {
		for _, tc := range testCases {
			t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
