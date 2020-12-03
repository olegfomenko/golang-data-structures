package segment_tree

type Tree interface {
	Get(l int, r int) Value

	Assign(value Value, pos int)
}

type IntTree struct {
	t    []Value
	size int
}

func (tree IntTree) Init(size int) {

}

func (tree IntTree) Get(l int, r int) Value {
	return get(tree.t, 1, 0, tree.size-1, l, r)
}

func (tree IntTree) Assign(value Value, pos int) {
	update(tree.t, 1, 0, tree.size-1, pos, value)
}

func GetIntTree(size int) Tree {
	return IntTree{getEmpty(size), size}
}
