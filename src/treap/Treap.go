package treap

type Treap struct {
	// Treap root node (with highest priority)
	root *node
}

type treap interface {
	insert(x Value, y Value)

	delete(x Value)

	get(k int) Value
}

// Inserting pair (x, y) in treap,
// where x is a value and y is a priority
func (t *Treap) Insert(x Value, y Value) {
	node := &node{x, y, 1, nil, nil}

	if t.root == nil {
		t.root = node
	} else {
		t1, t2 := t.root.split(x)
		t1 = t1.merge(node)
		t.root = t1.merge(t2)
	}
}

// Deleting x from treap
func (t *Treap) Delete(x Value) {
	t.root = t.root.delete(x)
}

// Get k-th max value
func (t *Treap) Get(k int) Value {
	return t.root.get(k)
}
