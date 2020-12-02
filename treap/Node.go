package treap

type Node struct {
	// value, priority, subtree size
	x, y, sz int

	// son's nodes
	left, right *Node
}

type INode interface {
	size() int

	update()

	merge(t *Node) *Node

	split(k int) (*Node, *Node)

	get(k int) int

	delete(x int) *Node
}

// Get size of Node
func (node *Node) size() int {
	if node == nil {
		return 0
	} else {
		return node.sz
	}
}

// Update node size after sons reassignment
func (node *Node) update() {
	if node != nil {
		node.sz = 1 + node.left.size() + node.right.size()
	}
}

// Merge current tree and tree with bigger values according to nodes priority
func (t1 *Node) merge(t2 *Node) *Node {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	if t1.y < t2.y {
		t2.left = t1.merge(t2.left)
		t2.update()
		return t2
	} else {
		t1.right = t1.right.merge(t2)
		t1.update()
		return t1
	}
}

// Splitting tree by key.
// Values that less then key will be stored int left answer
func (t *Node) split(k int) (*Node, *Node) {
	if t == nil {
		return nil, nil
	}

	if k <= t.x {
		t1, t2 := t.left.split(k)
		t.left = t2
		t.update()
		return t1, t
	} else {
		t1, t2 := t.right.split(k)
		t.right = t1
		t.update()
		return t, t2
	}
}

// Getting k-th maximum value in treap
func (t *Node) get(k int) int {
	sz := t.right.size() + 1

	if sz == k {
		return t.x
	}

	if sz < k {
		return t.left.get(k - sz)
	} else {
		return t.right.get(k)
	}
}

// Deleting value x from treap
func (t *Node) delete(x int) *Node {
	if t == nil {
		return nil
	}

	if t.x == x {
		return t.left.merge(t.right)
	}

	if x < t.x {
		t.left = t.left.delete(x)
	} else {
		t.right = t.right.delete(x)
	}

	t.update()
	return t
}
