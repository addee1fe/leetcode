package solution

func moveZeroes(nums []int) {
	p := 0
	for i := range nums {
		if nums[i] != 0 {
			hold := nums[i]
			nums[i] = nums[p]
			nums[p] = hold
			p++
		}
	}
}
