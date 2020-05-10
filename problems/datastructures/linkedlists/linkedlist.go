package linkedlists

import (
	"strconv"
	"strings"
)

// Node represents a singular element in the linkedList
type Node struct {
	Value int
	Next  *Node
}

// LinkedList represents a conveniene wrapper to a linkedlist of integers
type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

// Init initalizes the linkedlist with the given list of integers
func (l *LinkedList) Init(input []int) {
	for _, val := range input {
		l.Add(val)
	}
}

// Add adds a single element to the linkedlist
func (l *LinkedList) Add(val int) {
	n := Node{Value: val, Next: nil}
	// If the list is empty then create a new node
	// and point head & tail to it
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
	} else { // Append to the end otherwise
		l.Tail.Next = &n
		l.Tail = &n
	}
	l.Len++
}

func (l LinkedList) String() string {
	var output strings.Builder
	for i := l.Head; i != nil; i = i.Next {
		output.WriteString(strconv.Itoa(i.Value))
		if i.Next != nil {
			output.WriteString("-")
		}
	}
	return output.String()
}
