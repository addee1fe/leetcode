package solution

import "math"

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j := 0, n-1
	for k := n - 1; k >= 0; k-- {
		if abs(nums[i]) > abs(nums[j]) {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
	}
	return res
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
