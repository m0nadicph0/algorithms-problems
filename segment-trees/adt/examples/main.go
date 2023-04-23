package main

import "algorithms-problems/segment-trees/adt"

func main() {
	tree := adt.NewBinaryTreeSegmentTree([]int{1, 2, 3, 3, 2, 1})
	tree.Print()
}
