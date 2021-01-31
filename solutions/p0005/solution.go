package solution

// For better solution look for Manacher's algorithm
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var start, end int
	for i := range s {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		l := l1
		if l2 > l {
			l = l2
		}
		if l > end-start {
			start = i - (l-1)/2
			end = i + l/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
