package linkedlists

// O(n)
// Since we're removing elements, even with two loops we'll only effectively
// touch N elements. Notice that the same loop variable is being manipulated by
// both the loops.
func removeDups(input []int) string {
	var ll LinkedList
	ll.Init(input)
	seen := make(map[int]bool)
	for i := ll.Head; i != nil; i = i.Next {
		seen[i.Value] = true
		for i.Next != nil && seen[i.Next.Value] {
			i.Next = i.Next.Next
		}
	}
	return ll.String()
}

// O(n^2)
// Like the previous case, even though there are 3 loops, we'll only touch n^2
// elements in the worst case. Notice that the same loop variable is being
// manipulated by both the inner loops.
func removeDupsNoBuffer(input []int) string {
	var ll LinkedList
	ll.Init(input)
	for i := ll.Head; i != nil; i = i.Next {
		for j := i; j != nil; j = j.Next {
			for j.Next != nil && j.Next.Value == i.Value {
				j.Next = j.Next.Next
			}
		}
	}
	return ll.String()
}
