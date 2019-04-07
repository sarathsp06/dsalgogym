// Package linkedlists contains problems from cracking the coding interview from linkedlists chapter
// Code to check for palindrom using iterative approach
// url : https://play.golang.org/p/unJ2YWYHeq
package linkedlists

type String struct {
	Head *Char
}

func (s *String) Push(r rune) {
	n := &Char{Value: r, Next: s.Head}
	s.Head = n
}

// assumption string is not empty
func (s *String) Pop() rune {
	val := s.Head.Value
	s.Head = s.Head.Next
	return val
}

type Char struct {
	Next  *Char
	Value rune
}

func IterativePalindrom(str string) bool {
	s := new(String)
	strLen := len(str)
	if strLen == 0 {
		return true
	}

	if strLen == 1 {
		return true
	}

	mid := int(strLen / 2)

	for _, char := range []rune(str[:mid]) {
		s.Push(char)
	}
	if strLen%2 == 1 {
		mid++
	}
	for _, char := range str[mid:] {
		r := s.Pop()
		if char != r {
			return false
		}
	}
	return true
}
