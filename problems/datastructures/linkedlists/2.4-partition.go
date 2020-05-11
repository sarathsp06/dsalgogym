package linkedlists

// O(N)
func partition(input []int, target int) string {
	var ll LinkedList
	ll.Init(input)
	var targetLoc []*Node
	for ptr := ll.Head; ptr != nil; ptr = ptr.Next {
		if ptr.Value == target {
			targetLoc = append(targetLoc, ptr)
		}

		if len(targetLoc) >= 1 && ptr.Value < target {
			targetPtr := targetLoc[0]
			targetLoc = targetLoc[1:]
			targetPtr.Value = ptr.Value
			ptr.Value = target
			targetLoc = append(targetLoc, ptr)
		}
	}
	return ll.String()
}
