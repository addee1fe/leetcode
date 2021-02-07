package solution

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{1}
	}
	var carry bool
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
		carry = true
	}
	if carry {
		return append([]int{1}, digits...)
	}
	return digits
}
