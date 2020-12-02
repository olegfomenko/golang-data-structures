/**
	E-Olymp problem:
	https://www.e-olymp.com/uk/problems/687
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	x, y, sz int
	left, right *Node
}

type INode interface {
	size() int
	update()
}

func (node * Node) size() int {
	if node == nil {
		return 0
	} else {
		return node.sz
	}
}

func (node * Node) update() {
	if node != nil {
		node.sz = 1 + node.left.size() + node.right.size()
	}
}

type Treap struct {
	root * Node
}

func split(t *Node, k int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}

	if k <= t.x {
		t1, t2 := split(t.left, k)
		t.left = t2
		t.update()
		return t1, t
	} else {
		t1, t2 := split(t.right, k)
		t.right = t1
		t.update()
		return t, t2
	}
}

func merge(t1 *Node, t2 *Node) *Node {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	if t1.y < t2.y {
		t2.left = merge(t1, t2.left)
		t2.update()
		return t2
	} else {
		t1.right = merge(t1.right, t2)
		t1.update()
		return t1
	}
}

func get(t *Node, k int) int {
	sz := t.right.size() + 1

	if sz == k {
		return t.x
	}

	if sz < k {
		return get(t.left, k - sz)
	} else {
		return get(t.right, k)
	}
}

func delete(t * Node, x int) *Node {
	if t == nil {
		return nil
	}

	if t.x == x {
		return merge(t.left, t.right)
	}

	if x < t.x {
		t.left = delete(t.left, x)
	} else {
		t.right = delete(t.right, x)
	}

	t.update()
	return t
}

type ITreap interface {
	insert(x int, y int)

	delete(x int)

	get(k int) int
}

func (t * Treap) insert(x int, y int) {
	node := &Node{x, y, 1, nil, nil}

	if t.root == nil {
		t.root = node
	} else {
		t1, t2 := split(t.root, x)
		t1 = merge(t1, node)
		t.root = merge(t1, t2)
	}
}

func (t * Treap) delete(x int) {
	t.root = delete(t.root, x)
}

func (t * Treap) get(k int) int {
	return get(t.root, k)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var treap Treap

	for i := 0; i < n; i++ {
		var t, val int
		fmt.Fscan(in, &t, &val)

		switch t {
		case 1:
			treap.insert(val, rand.Int())
		case -1:
			treap.delete(val)
		case 0:
			fmt.Println(treap.get(val))
		}
	}

}