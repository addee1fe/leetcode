package solution

func singleNumber(nums []int) int {
	var a, b int
	for _, v := range nums {
		a ^= v &^ b
		b ^= v &^ a
	}
	return a
}
