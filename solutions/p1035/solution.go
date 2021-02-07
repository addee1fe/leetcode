package solution

func maxUncrossedLines(A []int, B []int) int {
	m, n := len(A), len(B)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j] < dp[i][j-1] { // can be replaced with max(a,b)
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
