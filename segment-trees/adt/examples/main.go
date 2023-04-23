package main

import (
	"algorithms-problems/segment-trees/adt"
	"fmt"
)

func main() {
	tree := adt.NewBinaryTreeSegmentTree([]int{1, 2, 3, 3, 2, 1})
	tree.Print()
	fmt.Println(tree.Query(0, 22))
	tree.Update(0, 5)
	tree.Print()
	fmt.Println(tree.Query(0, 22))
}
