package solution

// also can be solved with Gauss formula(arithmetic progression)
func missingNumber(nums []int) int {
	m := len(nums)
	for i, v := range nums {
		m ^= i ^ v
	}
	return m
}
