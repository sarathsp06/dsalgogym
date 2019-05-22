package listintersection

import "strings"

const linkStr = "->"

// Node represents one node in a linked list
type Node struct {
	val  string
	next *Node
}

// List represents a singly linked list
type List struct {
	head *Node
	tail *Node
	len  int
}

// Len returns the length
func (l *List) Len() int {
	return l.len
}

// CheckIntersection checks
func (l *List) CheckIntersection(ll *List) *Node {
	diff := l.Len() - ll.Len()
	m, n := l.head, ll.head
	if diff < 0 {
		diff = 0 - diff
		m, n = ll.head, l.head
	}
	for i := 0; i < diff; i++ {
		m = m.next
	}
	for m != nil && n != nil && m.val != n.val {
		m, n = m.next, n.next
	}
	if m != nil && n != nil {
		return m
	}
	return nil
}

//ListFromString creates a list given a string
//eg:
//	ListFromString("3->7->8->10")
func ListFromString(list string) *List {
	if list == "" {
		return new(List)

	}
	items := strings.Split(list, linkStr)
	l := &List{}
	l.head = &Node{val: items[0]}
	l.tail = l.head
	for _, item := range items[1:] {
		l.tail.next = &Node{val: item}
		l.tail = l.tail.next
	}
	l.len = len(items)
	return l
}

func (l *List) String() string {
	var vals []string
	m := l.head
	for m != nil {
		vals = append(vals, m.val)
		m = m.next
	}
	return strings.Join(vals, linkStr)
}
