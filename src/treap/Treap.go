package treap

type Treap interface {
	// Inserting pair (x, y) in treap,
	// where x is a value and y is a priority
	Insert(x Value, y Value)

	// Delete value from set
	Delete(x Value)

	// Get k-th maximum value
	Get(k int) Value
}

type treap struct {
	// Treap root node (with highest priority)
	root *node
}

func (t *treap) Insert(x Value, y Value) {
	node := &node{x, y, 1, nil, nil}

	if t.root == nil {
		t.root = node
	} else {
		t1, t2 := t.root.split(x)
		t1 = t1.merge(node)
		t.root = t1.merge(t2)
	}
}

func (t *treap) Delete(x Value) {
	t.root = t.root.delete(x)
}

func (t *treap) Get(k int) Value {
	return t.root.get(k)
}

func GetInstance() Treap {
	return &treap{nil}
}
