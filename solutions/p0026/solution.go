package solution

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		nums[i], nums[j] = nums[j], nums[i]
	}
	return len(nums[:j+1])
}
