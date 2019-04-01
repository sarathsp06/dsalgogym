//Code to check for palindrom using iterative approach
// url : https://play.golang.org/p/unJ2YWYHeq
package main

import (
	"fmt"
)

type String struct {
	Head *Node
}

func (s String) String() string {
	if s.Head == nil {
		return ""
	}
	tmp := s.Head
	var str []rune
	for tmp != nil {
		str = append(str, tmp.Value)
		tmp = tmp.Next
	}
	return string(str)
}

func NewString(s string) String {
	if len(s) == 0 {
		return String{}
	}
	sRunes := []rune(s)
	str := String{}
	str.Head = &Node{Value: sRunes[0]}
	tmp := str.Head
	for _, r := range sRunes[1:] {
		tmp.Next = &Node{Value: r}
		tmp = tmp.Next
	}
	return str
}

type Stack struct {
	Head *Node
}

func (s *Stack) Push(n Node) {
	n.Next = s.Head
	s.Head = &n
}

type Node struct {
	Next  *Node
	Value rune
}

func IterativePalindrom(a String) bool {
	s := &Stack{}
	p1 := a.Head
	if a.Head == nil {
		return true
	}
	s.Push(*p1)
	p2 := a.Head.Next
	if p2 == nil {
		return true
	}

	for p2.Next != nil && p2.Next.Next != nil {
		p2 = p2.Next.Next
		p1 = p1.Next
		s.Push(*p1)
	}

	if p2.Next == nil {
		p1 = p1.Next
	} else {
		p1 = p1.Next.Next
	}
	p2 = s.Head
	for p2 != nil {
		if p2.Value != p1.Value {
			return false
		}
		p2 = p2.Next
		p1 = p1.Next
	}
	return true
}

func main() {
	malayalam := NewString("MalayalaM")
	paap := NewString("PAAP")
	paapo := NewString("PAAPo")
	p := NewString("P")
	pa := NewString("PA")
	empty := NewString("")

	fmt.Println("malayalam", IterativePalindrom(malayalam))
	fmt.Println("paap", IterativePalindrom(paap))
	fmt.Println("paapo", IterativePalindrom(paapo))
	fmt.Println("p", IterativePalindrom(p))
	fmt.Println("pa", IterativePalindrom(pa))
	fmt.Println("", IterativePalindrom(empty))

}
