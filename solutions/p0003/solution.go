package solution

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	var ans int
	index := [26]int{}
	for j, i := 0, 0; j < n; j++ {
		i = max(index[s[j]-'a'], i)
		ans = max(ans, j-i+1)
		index[s[j]-'a'] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
