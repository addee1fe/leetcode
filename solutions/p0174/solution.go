package solution

// Please note, this is not mine solution
// TODO: back to this solution later

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	lr, lc := len(dungeon), len(dungeon[0])
	dp := make([][]int, lr)
	for i := range dungeon {
		dp[i] = make([]int, lc)
	}

	if dungeon[lr-1][lc-1] > 0 {
		dp[lr-1][lc-1] = 1
	} else {
		dp[lr-1][lc-1] = -1*dungeon[lr-1][lc-1] + 1
	}

	for i := lr - 2; i >= 0; i-- {
		if dungeon[i][lc-1] > 0 {
			if dp[i+1][lc-1]-dungeon[i][lc-1] > 0 {
				dp[i][lc-1] = dp[i+1][lc-1] - dungeon[i][lc-1]
			} else {
				dp[i][lc-1] = 0
			}
		} else {
			if dp[i+1][lc-1] == 0 {
				dp[i][lc-1] = 1 + -1*dungeon[i][lc-1]
			} else {
				dp[i][lc-1] = dp[i+1][lc-1] + -1*dungeon[i][lc-1]
			}
		}
	}

	for j := lc - 2; j >= 0; j-- {
		if dungeon[lr-1][j] > 0 {
			if dp[lr-1][j+1]-dungeon[lr-1][j] > 0 {
				dp[lr-1][j] = dp[lr-1][j+1] - dungeon[lr-1][j]
			} else {
				dp[lr-1][j] = 0
			}
		} else {
			if dp[lr-1][j+1] == 0 {
				dp[lr-1][j] = 1 + -1*dungeon[lr-1][j]
			} else {
				dp[lr-1][j] = dp[lr-1][j+1] + -1*dungeon[lr-1][j]
			}
		}
	}

	cm := 0
	for i := lr - 2; i >= 0; i-- {
		for j := lc - 2; j >= 0; j-- {
			cm = min(dp[i+1][j], dp[i][j+1])
			if dungeon[i][j] > 0 {
				if cm-dungeon[i][j] > 0 {
					dp[i][j] = cm - dungeon[i][j]
				} else {
					dp[i][j] = 0
				}
			} else {
				if cm == 0 {
					dp[i][j] = 1 + -1*dungeon[i][j]
				} else {
					dp[i][j] = cm + -1*dungeon[i][j]
				}
			}
		}
	}
	if dp[0][0] == 0 {
		return 1
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
