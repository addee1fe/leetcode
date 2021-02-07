package solution

func singleNumber(nums []int) int {
	unpaired := nums[0]
	for i := 1; i < len(nums); i++ {
		unpaired ^= nums[i]
	}
	return unpaired
}
