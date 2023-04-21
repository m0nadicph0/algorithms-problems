package linked_list_cycle

func hasCycleRec(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	return hasCycleHelper(head, head.Next)
}

func hasCycleHelper(slow *Node, fast *Node) bool {
	if slow == nil || fast == nil || fast.Next == nil {
		return false
	}

	if slow == fast {
		return true
	}

	return hasCycleHelper(slow.Next, fast.Next.Next)
}
