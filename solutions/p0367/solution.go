package solution

// Newton's method
func isPerfectSquare(num int) bool {
	x := num
	for x*x > num {
		x = (x + num/x) >> 1
	}
	return x*x == num
}
