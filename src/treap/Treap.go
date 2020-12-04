package treap

type Treap interface {
	// Inserting pair (x, y) in treap,
	// where x is a value and y is a priority
	Insert(x value, y value)

	// Delete value from set
	Delete(x value)

	// Get k-th maximum value
	Get(k int) value
}

type treap struct {
	// Treap root node (with highest priority)
	root *node
}

func (t *treap) Insert(x value, y value) {
	node := &node{x, y, 1, nil, nil}

	if t.root == nil {
		t.root = node
	} else {
		t1, t2 := t.root.split(x)
		t1 = t1.merge(node)
		t.root = t1.merge(t2)
	}
}

func (t *treap) Delete(x value) {
	t.root = t.root.delete(x)
}

func (t *treap) Get(k int) value {
	return t.root.get(k)
}

func GetInstance() Treap {
	return &treap{nil}
}
