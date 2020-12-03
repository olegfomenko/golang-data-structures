package main

import (
	"../segment-tree"
	"fmt"
)

func main() {
	var tree segment_tree.Tree = segment_tree.GetTree(5, segment_tree.IntValue(0))
	tree.Assign(segment_tree.IntValue(1), 0)
	tree.Assign(segment_tree.IntValue(14), 3)
	tree.Assign(segment_tree.IntValue(5), 4)

	fmt.Println(tree.Get(0, 4), tree.Get(0, 0), tree.Get(3, 4))
}
