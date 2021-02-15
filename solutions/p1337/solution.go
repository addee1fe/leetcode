package solution

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	score := make([]int, m)
	for i := 0; i < m; i++ {
		var j int
		for ; j < n; j++ {
			if mat[i][j] == 0 {
				break
			}
		}
		score[i] = j*m + i
	}
	sort.Ints(score)
	for i := range score {
		score[i] %= m
	}
	return score[:k]
}
