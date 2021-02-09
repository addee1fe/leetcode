package solution

func findNumbers(nums []int) int {
	var r int
	for _, v := range nums {
		if countDigits(v)%2 == 0 {
			r++
		}
	}
	return r
}

func countDigits(a int) int {
	var c int
	for a > 0 {
		a /= 10
		c++
	}
	return c
}
