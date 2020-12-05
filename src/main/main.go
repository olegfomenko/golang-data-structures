package main

import (
	"../segment-tree"
	"../treap"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var tree segment_tree.Tree = segment_tree.GetTree(5, segment_tree.IntValue(0))
	tree.Assign(segment_tree.IntValue(1), 0)
	tree.Assign(segment_tree.IntValue(14), 3)
	tree.Assign(segment_tree.IntValue(5), 4)

	fmt.Println(tree.Get(0, 4), tree.Get(0, 0), tree.Get(3, 4))

	rand.Seed(time.Now().UnixNano())
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var set = treap.GetInstance()

	for i := 0; i < n; i++ {
		var t, val int
		fmt.Fscan(in, &t, &val)

		switch t {
		case 1:
			set.Insert(treap.IntValue(val))
		case -1:
			set.Delete(treap.IntValue(val))
		case 0:
			fmt.Println(set.Get(val))
		}
	}
}
