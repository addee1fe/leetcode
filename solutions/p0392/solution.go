package solution

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	var c int
	for i := range t {
		if t[i] == s[c] {
			c++
		}
		if c == len(s) {
			return true
		}
	}
	return false
}
