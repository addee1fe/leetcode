package solution

import "math/big"

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	s := factorial(big.NewInt(int64(m + n - 2)))
	t := factorial(big.NewInt(int64(m - 1)))
	u := factorial(big.NewInt(int64(n - 1)))
	s = s.Div(s, t)
	u = u.Div(s, u)
	return int(u.Int64())
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

// DP Solution
// func uniquePaths(m int, n int) int {
//     dp := make([]int, n)
//     dp[0] = 1
//     for i := 0; i < m; i++ {
//         for j := 1; j < n; j++ {
//             dp[j] += dp[j - 1]
//         }
//     }
//     return dp[n - 1]
// }
