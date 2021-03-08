package solution

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return 2
		}
	}
	return 1
}
