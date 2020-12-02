package treap

type Treap struct {
	root *Node
}

type ITreap interface {
	insert(x int, y int)

	delete(x int)

	get(k int) int
}

// Inserting pair (x, y) in treap, where x is a value and y is a priority
func (t *Treap) Insert(x int, y int) {
	node := &Node{x, y, 1, nil, nil}

	if t.root == nil {
		t.root = node
	} else {
		t1, t2 := t.root.split(x)
		t1 = t1.merge(node)
		t.root = t1.merge(t2)
	}
}

// Deleting x from treap
func (t *Treap) Delete(x int) {
	t.root = t.root.delete(x)
}

// Get k-th max value
func (t *Treap) Get(k int) int {
	return t.root.get(k)
}
