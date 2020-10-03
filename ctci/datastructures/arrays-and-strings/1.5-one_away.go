package arrays

import "math"

// Complexity: O(n)
func isOneAway(str1, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	foundEdit := false
	if math.Abs(float64(l1)-float64(l2)) > 1 {
		// If the lenghts differ by more than one, then don't even bother
		return false
	} else if l1 == l2 {
		// Replacement is the only option.
		// Both strings can differ by utmost one character
		for i := 0; i < l1; i++ {
			if str1[i] != str2[i] {
				if foundEdit {
					return false
				}
				foundEdit = true
			}
		}
	} else {
		// Insert or Delete if strings differ in length by 1 character
		var i, j int
		for ; i < l1 && j < l2; i, j = i+1, j+1 {
			if str1[i] == str2[j] {
				continue
			}
			if str1[i] != str2[j] {
				if l1 > l2 && str1[i+1] == str2[j] && !foundEdit {
					// If str1 is longer, then delete is the only option
					foundEdit = true
					i++
				} else if l1 < l2 && str1[i] == str2[j+1] && !foundEdit {
					// If str1 is smaller, then insert is the only option
					j++
					foundEdit = true
				} else {
					return false
				}
			}
		}
	}
	return true
}
