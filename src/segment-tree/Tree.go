package segment_tree

type Tree interface {
	// Get sum on section [l; r]
	// Attention: this struct uses zero-indexing
	Get(l int, r int) Value

	// Change value on certain pos
	Assign(value Value, pos int)

	// Inner get function that contains main segment tree get-on-segment logic
	get(v int, tl int, tr int, l int, r int) Value

	// Inner update function that contains main segment tree re-assign logic
	update(v int, tl int, tr int, pos int, value Value)
}

type segmentTree struct {
	t    []Value
	size int
}

func (tree segmentTree) Get(l int, r int) Value {
	return tree.get(1, 0, tree.size-1, l, r)
}

func (tree segmentTree) Assign(value Value, pos int) {
	tree.update(1, 0, tree.size-1, pos, value)
}

func (tree segmentTree) get(v int, tl int, tr int, l int, r int) Value {
	if l > r {
		return nil
	}

	if l == tl && r == tr {
		return tree.t[v]
	}

	var tm int = (tl + tr) / 2

	left := tree.get(v*2, tl, tm, l, min(r, tm))
	right := tree.get(v*2+1, tm+1, tr, max(l, tm+1), r)

	if left == nil {
		return right
	} else {
		return left.Add(right)
	}
}

func (tree segmentTree) update(v int, tl int, tr int, pos int, value Value) {
	if tl == tr {
		tree.t[v] = value
	} else {
		var tm int = (tl + tr) / 2

		if pos <= tm {
			tree.update(v*2, tl, tm, pos, value)
		} else {
			tree.update(v*2+1, tm+1, tr, pos, value)
		}

		tree.t[v] = tree.t[v*2].Add(tree.t[v*2+1])
	}
}

func GetTree(size int, defaultValue Value) Tree {
	return segmentTree{getEmpty(size, defaultValue), size}
}
