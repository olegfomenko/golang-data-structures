package treap

type orderedNode struct {
	// value and priority
	x, y value

	// node subtree size
	sz int

	left, right *orderedNode
}

type orderedTreapNode interface {
	// Get size of Node
	size() int

	// Update node size after sons reassignment
	update()

	// Merge current tree and tree with bigger values according to nodes priority
	merge(t *orderedNode) *orderedNode

	// Split first k elements from tree
	split(k int) (*orderedNode, *orderedNode)

	// Deleting value on id position from array
	deleteById(id int) *orderedNode

	// Get value by id
	getById(id int) *orderedNode
}

func (node *orderedNode) size() int {
	if node == nil {
		return 0
	} else {
		return node.sz
	}
}

func (node *orderedNode) update() {
	if node != nil {
		node.sz = 1 + node.left.size() + node.right.size()
	}
}

func (t1 *orderedNode) merge(t2 *orderedNode) *orderedNode {
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

func (t *orderedNode) split(k int) (*orderedNode, *orderedNode) {
	if t == nil {
		return nil, nil
	}

	sz := t.left.size() + 1

	if sz == k {
		node := t.left
		t.left = nil
		t.update()
		return t, node
	}

	if sz > k {
		var t1, t2 = t.left.split(k)
		t.left = t2
		t.update()
		return t1, t
	} else {
		var t1, t2 = t.right.split(k - sz)
		t.right = t1
		t.update()
		return t, t2
	}
}

func (t *orderedNode) deleteById(id int) *orderedNode {
	if t == nil {
		return nil
	}

	sz := t.left.size() + 1

	if sz == id {
		return t.left.merge(t.right)
	}

	if sz > id {
		t.left = t.left.deleteById(id)
		t.update()
		return t
	} else {
		t.right = t.right.deleteById(id - sz)
		t.update()
		return t
	}
}

func (t *orderedNode) getById(id int) *orderedNode {
	if t == nil {
		return nil
	}

	sz := t.left.size() + 1

	if sz == id {
		return t
	}

	if sz > id {
		return t.left.getById(id)
	} else {
		return t.right.getById(id - sz)
	}
}
