package arrays

import "sort"

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
