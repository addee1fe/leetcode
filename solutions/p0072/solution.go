package solution

// TODO: refactor
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i < l1+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < l2+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[l1][l2]

}
func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	var m int
	m = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < m {
			m = a[i]
		}
	}
	return m
}

// func minDistance(word1 string, word2 string) int {
//     m := len(word1)
//     n := len(word2)
//     dp := make([][]int, m+1)
//     for i := 0; i < m+1; i++ {
//         dp[i] = make([]int, n+1)
//     }
//     for i := 1; i <= m; i++ {
//         dp[i][0] = i
//     }
//     for i := 1; i <=n; i++ {
//         dp[0][i] = i
//     }
//     for i := 1; i <= m; i++ {
//         for j := 1; j <= n; j++ {
//             if word1[i-1] == word2[j-1] {
//                 dp[i][j] = dp[i-1][j-1]
//             } else {
//                 x := min(dp[i-1][j-1], dp[i][j-1])
//                 dp[i][j] = min(x, dp[i-1][j]) + 1
//             }
//         }
//     }
//     return dp[m][n]
// }
// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }
