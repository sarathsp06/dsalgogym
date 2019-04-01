//Program to add two linked lists where : 216 = 2->1->6,12 = 1->2
//link : https://play.golang.org/p/EFmTmAtzxu
package main

import (
	"fmt"
)

type Node struct {
	Next  *Node
	Digit int
}

type List struct {
	Head *Node
}

func (l List) Len() int {
	n := l.Head
	len := 0
	for n != nil {
		len += 1
		n = n.Next
	}
	return len
}

func (l List) String() string {
	var result []byte
	node := l.Head
	for node != nil {
		result = append(result, byte(node.Digit)+'0')
		node = node.Next
	}
	return string(result)
}

func (l *List) Add(v *List) {
	carry := l.Head.Add(v.Head)
	if carry != 0 {
		node := &Node{Digit: carry, Next: l.Head}
		l.Head = node
	}
}

func (n *Node) Add(v *Node) int {
	if n == nil {
		return 0
	}
	carry := n.Next.Add(v.Next)
	n.Digit = n.Digit + v.Digit + carry
	carry = n.Digit / 10
	n.Digit = n.Digit % 10
	return carry
}

func AdjustLength(a, b *List) {
	lenA := a.Len()
	lenB := b.Len()
	if lenA == lenB {
		return
	}
	if lenA > lenB {
		addPadding(b, lenA-lenB)
		return
	}
	addPadding(a, lenB-lenA)
}

func addPadding(l *List, len int) {
	for i := 0; i < len; i++ {
		node := new(Node)
		node.Next = l.Head
		l.Head = node
	}
}

func main() {
	L1 := new(List)
	L1.Head = &Node{Digit: 8}
	L1.Head.Next = &Node{Digit: 7}
	L1.Head.Next.Next = &Node{Digit: 8}
	L2 := new(List)
	L2.Head = &Node{Digit: 8}
	L2.Head.Next = &Node{Digit: 2}
	L2.Head.Next.Next = &Node{Digit: 5}
	L2.Head.Next.Next.Next = &Node{Digit: 3}

	//Before adjesting length
	fmt.Println("Before adjesting length")
	fmt.Println(L1)
	fmt.Println(L2)
	AdjustLength(L1, L2)

	//After adjesting length
	fmt.Println("After adjesting length")
	fmt.Println(L1)
	fmt.Println(L2)

	//Sum
	L1.Add(L2)
	fmt.Println("A + B")
	fmt.Println(L1)

}
