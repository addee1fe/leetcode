package solution

import "sort"

func change(amount int, coins []int) int {
	if len(coins) == 1 {
		if coins[0] == amount {
			return 1
		}
		if coins[0] < amount {
			return 0
		}
	}
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, j := range coins {
		for i := j; i <= amount; i++ {
			dp[i] += dp[i-j]
		}
	}
	return dp[amount]
}
