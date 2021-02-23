package solution

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		n := matrix[i][j]
		if n == target {
			return true
		}
		if n > target {
			j--
		} else {
			i++
		}
	}
	return false
}
