package solution

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow, fast = nums[0], nums[fast]
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
