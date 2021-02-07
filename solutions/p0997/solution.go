package solution

func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {
		return N
	}
	c := make([]int, N+1)
	for _, v := range trust {
		c[v[0]]--
		c[v[1]]++
	}
	for i := 0; i <= N; i++ {
		if c[i] == N-1 {
			return i
		}
	}
	return -1
}
