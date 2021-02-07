package solution

func productExceptSelf(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	prods := make([]int, len(nums))
	p := 1
	for i, n := range nums {
		p *= n
		prods[i] = p
	}
	prods[len(nums)-1] = prods[len(nums)-2]
	p = 1
	for i := len(nums) - 2; i > 0; i-- {
		p *= nums[i+1]
		prods[i] = prods[i-1] * p
	}
	prods[0] = p * nums[1]
	return prods
}
