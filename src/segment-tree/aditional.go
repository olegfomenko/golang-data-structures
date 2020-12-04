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

func getEmpty(size int, def value) []value {
	values := make([]value, 4*size+4)

	for i := 0; i < 4*size+4; i++ {
		values = append(values, def)
	}

	return values
}
