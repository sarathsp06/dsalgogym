package linkedlists

// O(1)
func removeMiddle(input []int, target int) string {
	ll := InitLinkedList(input)
	for i := ll.Head; i != nil && i.Next != nil; i = i.Next {
		if i.Value == target {
			i.Value = i.Next.Value
			i.Next = i.Next.Next
		}
	}
	return ll.String()
}
