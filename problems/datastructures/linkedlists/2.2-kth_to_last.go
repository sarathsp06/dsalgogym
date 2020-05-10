package linkedlists

// O(n)
func kthToLast(k int, input []int) int {
	var ll LinkedList
	ll.Init(input)
	if k >= ll.Len || k < 0 {
		return -1
	}
	// Head is at a distance of Len-1 from Last
	dist := ll.Len - 1
	for i := ll.Head; i != nil; i = i.Next {
		if dist == k {
			return i.Value
		}
		dist--
	}
	return -1
}
