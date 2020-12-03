package main

import (
	"../segment-tree"
	"fmt"
)

func main() {
	value := segment_tree.GetInt(1)
	value = value.Add(segment_tree.GetInt(2))
	fmt.Println(value)

	str := segment_tree.GetString("abacaba")
	str = str.Add(segment_tree.GetString("xxx"))
	fmt.Println(str)

	var tree segment_tree.Tree = segment_tree.GetIntTree(5)
	tree.Assign(segment_tree.IntValue(1), 0)
	tree.Assign(segment_tree.IntValue(14), 3)
	tree.Assign(segment_tree.IntValue(5), 4)

	fmt.Println(tree.Get(0, 4), tree.Get(0, 0), tree.Get(3, 4))
}
