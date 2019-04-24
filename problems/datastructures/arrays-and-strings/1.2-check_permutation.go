package arrays

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
