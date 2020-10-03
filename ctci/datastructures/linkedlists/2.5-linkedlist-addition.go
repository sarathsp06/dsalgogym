// Package linkedlists has function to add two linked lists where : 216 = 2->1->6,12 = 1->2
// link : https://play.golang.org/p/EFmTmAtzxu
package linkedlists

import "strconv"

type Digit struct {
	Next  *Digit
	Digit int
}

type List struct {
	Head *Digit
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

// String implements the stringify interface
func (l List) String() string {
	var result []byte
	node := l.Head
	for node != nil {
		result = append(result, byte(node.Digit)+'0')
		node = node.Next
	}
	lastIndex := len(result) - 1
	for i := 0; i <= lastIndex/2; i++ {
		result[i], result[lastIndex-i] = result[lastIndex-i], result[i]
	}
	return string(result)
}

// Add wrap digits add function
func (l *List) Add(v *List) {
	l.Head = l.Head.Add(v.Head)
}

// Add adds two numbers recursively
func (n *Digit) Add(v *Digit) *Digit {
	if n == nil {
		return v
	}
	if v == nil {
		return n
	}
	n.Digit = n.Digit + v.Digit
	carry := n.Digit / 10
	n.Digit = n.Digit % 10
	n.Next = n.Next.Add(v.Next)
	if n.Next != nil {
		n.Next.Digit += carry
	}
	return n
}

// O(N)
func sumListBackward(lst1 []int, lst2 []int) string {
	p1 := InitLinkedList(lst1).Head
	p2 := InitLinkedList(lst2).Head
	var result LinkedList
	carry := 0
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		positionalSum := (p1.Value + p2.Value + carry) % 10
		carry = (p1.Value + p2.Value + carry) / 10
		result.Add(positionalSum)
	}
	ptr := p1
	if ptr == nil {
		ptr = p2
	}
	for ; ptr != nil; ptr = ptr.Next {
		positionalSum := (ptr.Value + carry) % 10
		carry = (ptr.Value + carry) / 10
		result.Add(positionalSum)
	}
	return result.String()
}

// O(N)
func sumListForward(lst1 []int, lst2 []int) string {
	p1 := InitLinkedList(lst1).Head
	p2 := InitLinkedList(lst2).Head
	sum := 0
	for _, ptr := range []*Node{p1, p2} {
		num := 0
		for p := ptr; p != nil; p = p.Next {
			num = num*10 + p.Value
		}
		sum = sum + num
	}
	var result LinkedList
	for _, e := range strconv.Itoa(sum) {
		digit, _ := strconv.Atoi(string(e))
		result.Add(digit)
	}
	return result.String()
}
