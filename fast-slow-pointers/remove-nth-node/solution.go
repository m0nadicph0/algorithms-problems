package remove_nth_node

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) ToList() []int {
	result := make([]int, 0)
	current := n
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func removeNthFromEnd(head *Node, n int) *Node {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
