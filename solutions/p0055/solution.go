package solution

func canJump(nums []int) bool {
	left := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= left {
			left = i
		}
	}
	return left == 0
}
