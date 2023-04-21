package linked_list_cycle

type Node struct {
	Value int
	Next  *Node
}

func hasCycle(head *Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}
	return false
}
