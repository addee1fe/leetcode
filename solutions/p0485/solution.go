package solution

func findMaxConsecutiveOnes(nums []int) int {
	var max, sum int
	for _, v := range nums {
		sum = (sum + v) * v
		if max < sum {
			max = sum
		}
	}
	return max
}
