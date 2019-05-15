package substr

import "fmt"

// Substr stores a substring
type Substr struct {
	i, j           int
	lastOccurrence map[rune]int
}

func newSubstr() *Substr {
	substr := new(Substr)
	substr.lastOccurrence = make(map[rune]int)
	return substr
}

func (s Substr) nextSubStr() *Substr {
	substr := newSubstr()
	fmt.Println(s)
	defer fmt.Println("sub", substr)
	//find lowest last occurrence
	low := int(^uint(0) >> 1)
	for key := range s.lastOccurrence {
		if s.lastOccurrence[key] < low {
			low = s.lastOccurrence[key]
		}
	}

	substr.i = low + 1
	substr.j = s.j
	for key := range s.lastOccurrence {
		if s.lastOccurrence[key] != low {
			substr.lastOccurrence[key] = s.lastOccurrence[key]
		}
	}
	return substr
}

func (s Substr) len() int {
	return s.j - s.i + 1
}

// GetLargestSubstr finds the first largest substring with
//at-most  k distinct characters
func GetLargestSubstr(s string, k int) string {
	if s == "" || k == 0 {
		return ""
	}
	maxSubstr := newSubstr()
	substr := newSubstr()

	for idx, r := range []rune(s) {
		substr.lastOccurrence[r] = idx
		if len(substr.lastOccurrence) <= k {
			substr.j = idx
			continue
		}
		delete(substr.lastOccurrence, r)
		if maxSubstr.len() < substr.len() {
			maxSubstr = substr
		}
		substr = substr.nextSubStr()
		substr.j = idx
		substr.lastOccurrence[r] = idx
	}
	if maxSubstr.len() < substr.len() {
		maxSubstr = substr
	}
	return s[maxSubstr.i : maxSubstr.j+1]
}
