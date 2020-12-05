package treap

import "math/rand"

type OrderedTreap interface {
	// Inserting x value on pos position
	Insert(x value, pos int)

	// Delete value from set
	DeleteById(id int)

	// Get value by id
	GetById(id int) value

	// Swap two subarray
	Swap(l1 int, r1 int, l2 int, r2 int)

	// Delete subarray
	DeleteSubarray(l1 int, l2 int)
}

type orderedTreap struct {
	// Treap root node (with highest priority)
	root *orderedNode
}

func (t *orderedTreap) Insert(x value, pos int) {
	var t1, t2 = t.root.split(pos - 1)
	var node = &orderedNode{x, IntValue(rand.Int()), 1, nil, nil}

	t1 = t1.merge(node)
	t.root = t1.merge(t2)
}

func (t *orderedTreap) DeleteById(id int) {
	t.root = t.root.deleteById(id)
}

func (t *orderedTreap) GetById(id int) value {
	if t.root.sz < id {
		return nil
	}

	return t.root.getById(id).x
}

func (t *orderedTreap) Swap(l1 int, r1 int, l2 int, r2 int) {
	if r1 < l2 || l1 > r2 {
		// TODO
	}
}

func (t *orderedTreap) DeleteSubarray(l1 int, r1 int) {
	if l1 < r1 {
		// TODO
	}
}
