package solution

func reverse(x int) int {
	var r int
	for x != 0 {
		r = 10*r + x%10
		x /= 10
	}
	// only added to pass leetcode tests
	if r > 1<<31-1 || r < -1<<31 {
		return 0
	}
	return r
}

// func reverse(x int) int {
// 	var res int
// 	for x != 0 {
// 		res = 10*res + x%10
// 		x /= 10
// 	}
// 	return res
// }
