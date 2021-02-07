package solution

func removeElement(nums []int, val int) int {
	var w int
	for _, v := range nums {
		if v != val {
			nums[w] = v
			w++
		}
	}
	return w
}
