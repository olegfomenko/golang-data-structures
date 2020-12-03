package segment_tree

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func get(t []Value, v int, tl int, tr int, l int, r int) Value {
	if l > r {
		return nil
	}

	if l == tl && r == tr {
		return t[v]
	}

	var tm int = (tl + tr) / 2

	left := get(t, v*2, tl, tm, l, min(r, tm))
	right := get(t, v*2+1, tm+1, tr, max(l, tm+1), r)

	if left == nil {
		return right
	} else {
		return left.Add(right)
	}
}

func update(t []Value, v int, tl int, tr int, pos int, value Value) {
	if tl == tr {
		t[v] = value
	} else {
		var tm int = (tl + tr) / 2

		if pos <= tm {
			update(t, v*2, tl, tm, pos, value)
		} else {
			update(t, v*2+1, tm+1, tr, pos, value)
		}

		t[v] = t[v*2].Add(t[v*2+1])
	}
}

func getEmpty(size int, def Value) []Value {
	values := make([]Value, 4*size+4)

	for i := 0; i < 4*size+4; i++ {
		values = append(values, def)
	}

	return values
}
