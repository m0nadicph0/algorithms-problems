package middle_linked_list

func MiddleNodeRecursive(head *Node) *Node {
	return middleNodeHelper(head, head)
}

func middleNodeHelper(slow *Node, fast *Node) *Node {
	if fast == nil || fast.Next == nil {
		return slow
	}
	return middleNodeHelper(slow.Next, fast.Next.Next)
}
