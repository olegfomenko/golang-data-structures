package segment_tree

type Tree interface {
	// Get sum on section [l; r]
	// Attention: this struct uses zero-indexing
	Get(l int, r int) Value

	// Change value on certain pos
	Assign(value Value, pos int)
}

type segmentTree struct {
	t    []Value
	size int
}

func (tree segmentTree) Get(l int, r int) Value {
	return get(tree.t, 1, 0, tree.size-1, l, r)
}

func (tree segmentTree) Assign(value Value, pos int) {
	update(tree.t, 1, 0, tree.size-1, pos, value)
}

func GetIntTree(size int) Tree    { return segmentTree{getEmpty(size, IntValue(0)), size} }
func GetStringTree(size int) Tree { return segmentTree{getEmpty(size, StringValue("")), size} }
func GetFloatTree(size int) Tree  { return segmentTree{getEmpty(size, FloatValue(0)), size} }
