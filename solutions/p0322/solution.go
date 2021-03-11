package solution

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 1; i < len(dp); i++ {
		dp[i] = max
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
