package solution

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	p := -10000
	for i := range s {
		if s[i] == c {
			p = i
		}
		res[i] = i - p
	}
	p = 10000
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			p = i
		}
		res[i] = min(res[i], p-i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
