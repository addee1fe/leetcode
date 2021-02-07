package solution

// Simple solution
func arrangeCoins(n int) int {
	var i int
	for i < n {
		i++
		n -= i
	}
	return i
}

// Math oneliner
// func arrangeCoins(n int) int {
// 	return int(math.Sqrt(2*float64(n)+0.25) - 0.5)
// }

// Binary search
// func arrangeCoins(n int) int {
// 	left, right := 0, n
// 	for left <= right {
// 		k := left + (right-left)/2
// 		cur := k * (k + 1) / 2
// 		if cur == n {
// 			return k
// 		}
// 		if cur < n {
// 			left = k + 1
// 		} else {
// 			right = k - 1
// 		}
// 	}
// 	return right
// }
