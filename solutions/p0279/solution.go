package solution

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		min := i
		for j := 1; j*j <= i; j++ {
			if curr := 1 + dp[i-j*j]; min > curr {
				min = curr
			}
		}
		dp[i] = min
	}
	return dp[n]
}

// func numSquares(n int) int {
//     if math.Floor(math.Sqrt(float64(n))) == math.Ceil(math.Sqrt(float64(n))) { return 1 }
//     for n % 4 == 0 { n /= 4 }
//     if n % 8 == 7 { return 4 }
//     for i := 1; i * i <= n; i++ {
//         tmp := int(math.Sqrt(float64(n - i * i)))
//         if tmp * tmp == n - i * i { return 2 }
//     }
//     return 3
// }
