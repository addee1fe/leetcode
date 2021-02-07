package solution

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 || n <= 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}

// func isPowerOfTwo(n int) bool {
//     return (n & (n -1)) == 0 && n > 0
// }

// func isPowerOfTwo(n int) bool {
// 	count1Bits := 0
// 	for n > 0 {
// 			if n&1 != 0 {
// 					count1Bits++
// 			}
// 			n = n >> 1
// 	}
// 	return count1Bits == 1
// }
