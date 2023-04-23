package adt

import (
	"fmt"
)

type SegmentTreeNode struct {
	start int
	end   int
	value int
	left  *SegmentTreeNode
	right *SegmentTreeNode
}

type BinaryTreeSegmentTree struct {
	root *SegmentTreeNode
}

func NewBinaryTreeSegmentTree(arr []int) SegmentTree {
	root := buildTreeHelper(arr, 0, len(arr)-1)
	return &BinaryTreeSegmentTree{root: root}
}

func buildTreeHelper(arr []int, start int, end int) *SegmentTreeNode {
	if start > end {
		return nil
	}

	node := &SegmentTreeNode{start: start, end: end}

	if start == end {
		node.value = arr[start]
		return node
	} else {
		mid := (start + end) / 2
		node.left = buildTreeHelper(arr, start, mid)
		node.right = buildTreeHelper(arr, mid+1, end)
		node.value = node.left.value + node.right.value
		return node
	}
}

func (b *BinaryTreeSegmentTree) Update(index int, value int) {
	//TODO implement me
	panic("implement me")
}

func (b *BinaryTreeSegmentTree) Query(start int, end int) int {
	//TODO implement me
	panic("implement me")
}

func (b *BinaryTreeSegmentTree) Print() {
	printHelper(b.root, "", false)
}

func printHelper(node *SegmentTreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}
	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Printf("[%d,%d] -> %d\n", node.start, node.end, node.value)
	printHelper(node.left, prefix+getPrefix(isLeft), true)
	printHelper(node.right, prefix+getPrefix(isLeft), false)
}

func getPrefix(isLeft bool) string {
	if isLeft {
		return "│   "
	} else {
		return "    "
	}
}
