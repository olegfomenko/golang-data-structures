package segment_tree

type Tree interface {
	// Get sum on section [l; r]
	// Attention: this struct uses zero-indexing
	Get(l int, r int) Value

	// Change value on certain pos
	Assign(value Value, pos int)
}

type intTree struct {
	t    []Value
	size int
}

type floatTree struct {
	t    []Value
	size int
}

type stringTree struct {
	t    []Value
	size int
}

func (tree intTree) Get(l int, r int) Value    { return get(tree.t, 1, 0, tree.size-1, l, r) }
func (tree floatTree) Get(l int, r int) Value  { return get(tree.t, 1, 0, tree.size-1, l, r) }
func (tree stringTree) Get(l int, r int) Value { return get(tree.t, 1, 0, tree.size-1, l, r) }

func (tree intTree) Assign(value Value, pos int)    { update(tree.t, 1, 0, tree.size-1, pos, value) }
func (tree floatTree) Assign(value Value, pos int)  { update(tree.t, 1, 0, tree.size-1, pos, value) }
func (tree stringTree) Assign(value Value, pos int) { update(tree.t, 1, 0, tree.size-1, pos, value) }

func GetIntTree(size int) Tree    { return intTree{getEmpty(size, IntValue(0)), size} }
func GetFloatTree(size int) Tree  { return floatTree{getEmpty(size, FloatValue(0)), size} }
func GetStringTree(size int) Tree { return stringTree{getEmpty(size, StringValue("")), size} }
