package solution

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	h := [26]int{}

	for i := range s {
		h[s[i]-'a']++
		h[t[i]-'a']--
	}

	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
