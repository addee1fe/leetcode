package solution

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	count := make(map[int]int)
	count[0] = 1
	var cur, sum int
	for _, v := range nums {
		cur += v
		if val, ok := count[cur-k]; ok {
			sum += val
		}
		count[cur]++
	}
	return sum
}
