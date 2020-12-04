package treap

type node struct {
	// Node value and priority
	x, y value

	// size of subtree with current node as a root
	sz int

	// son's nodes
	left, right *node
}

type treapNode interface {
	// Get size of Node
	size() int

	// Update node size after sons reassignment
	update()

	// Merge current tree and tree with bigger values according to nodes priority
	merge(t *node) *node

	// Splitting tree by key.
	// Values that less then key will be stored int left answer
	split(k int) (*node, *node)

	// Getting k-th maximum value in treap
	get(k int) int

	// Deleting value x from treap
	delete(x int) *node
}

func (node *node) size() int {
	if node == nil {
		return 0
	} else {
		return node.sz
	}
}

func (node *node) update() {
	if node != nil {
		node.sz = 1 + node.left.size() + node.right.size()
	}
}

func (t1 *node) merge(t2 *node) *node {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	if t1.y.Less(t2.y) {
		t2.left = t1.merge(t2.left)
		t2.update()
		return t2
	} else {
		t1.right = t1.right.merge(t2)
		t1.update()
		return t1
	}
}

func (t *node) split(k value) (*node, *node) {
	if t == nil {
		return nil, nil
	}

	if t.x.Less(k) {
		t1, t2 := t.right.split(k)
		t.right = t1
		t.update()
		return t, t2
	} else {
		t1, t2 := t.left.split(k)
		t.left = t2
		t.update()
		return t1, t
	}
}

func (t *node) get(k int) value {
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

func (t *node) delete(x value) *node {
	if t == nil {
		return nil
	}

	if t.x.Equals(x) {
		return t.left.merge(t.right)
	}

	if x.Less(t.x) {
		t.left = t.left.delete(x)
	} else {
		t.right = t.right.delete(x)
	}

	t.update()
	return t
}
