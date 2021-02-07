package solution

func isPalindrome(x int) bool {
	tmp, rev := x, 0
	for tmp > 0 {
		rev = 10*rev + tmp%10
		tmp /= 10
	}
	return x == rev
}
