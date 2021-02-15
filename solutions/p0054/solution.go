package solution

// TODO: review
// Note: this solution fully copy-pasted from solution sections, coz I'm too dump to solve this in my own
func spiralOrder(matrix [][]int) []int {
	var ans []int
	if len(matrix) == 0 {
		return ans
	}
	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			ans = append(ans, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			ans = append(ans, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				ans = append(ans, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				ans = append(ans, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return ans
}
