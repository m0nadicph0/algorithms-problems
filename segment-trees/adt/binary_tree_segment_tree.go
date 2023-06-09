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
	updateHelper(b.root, index, value)
}

func updateHelper(node *SegmentTreeNode, index int, value int) {
	if node == nil {
		return
	}

	if node.start == index && node.end == index {
		node.value = value
		return
	}

	mid := (node.start + node.end) / 2

	if index <= mid {
		updateHelper(node.left, index, value)
	} else {
		updateHelper(node.right, index, value)
	}

	if node.left != nil && node.right != nil {
		node.value = node.left.value + node.right.value
	}
}

func (b *BinaryTreeSegmentTree) Query(start int, end int) int {
	return queryHelper(b.root, start, end)
}

func queryHelper(node *SegmentTreeNode, start int, end int) int {
	if start <= node.start && end >= node.end {
		return node.value
	}

	if start > node.end || end < node.start {
		return 0
	}

	return queryHelper(node.left, start, end) + queryHelper(node.right, start, end)
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
