package solution

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, 0)
	for i, v := range nums {
		if val, ok := index[target-v]; ok {
			return []int{val, i}
		}
		index[v] = i
	}
	return []int{}
}
