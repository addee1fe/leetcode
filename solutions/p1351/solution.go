package solution

func countNegatives(grid [][]int) int {
	var sum int
	for _, v := range grid {
		sum += countRow(v)
	}
	return sum
}

func countRow(row []int) int {
	l, r := 0, len(row)
	for l < r {
		m := (l + r) >> 1
		if row[m] >= 0 {
			l = m + 1
		} else {
			r = m
		}
	}
	return len(row) - l
}
