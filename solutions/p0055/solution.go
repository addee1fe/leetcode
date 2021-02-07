package solution

func canJump(nums []int) bool {
	i := 0
	for j := 0; i < len(nums) && i <= j; {
		j = max(i+nums[i], j)
		i++
	}
	return i == len(nums)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
