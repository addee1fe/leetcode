package solution

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for j := 2; j < n+1; j++ {
		for i := 1; i < j+1; i++ {
			dp[j] = dp[j] + dp[i-1]*dp[j-i]
		}
	}
	return dp[n]
}
