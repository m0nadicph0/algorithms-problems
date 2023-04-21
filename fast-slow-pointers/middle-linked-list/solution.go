package middle_linked_list

type Node struct {
	Value int
	Next  *Node
}

func MiddleNode(head *Node) *Node {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
