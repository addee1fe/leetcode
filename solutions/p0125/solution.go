package solution

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = clean(s)
	return s == reverse(s)
}

func clean(s string) string {
	s = strings.ToLower(s)
	var sb strings.Builder
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			sb.WriteRune(v)
		}
	}
	return sb.String()
}

func reverse(s string) string {
	sl := []rune(s)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}
