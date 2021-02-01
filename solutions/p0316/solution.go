package solution

func removeDuplicateLetters(s string) string {
	if len(s) <= 1 {
		return s
	}
	var res []rune
	var h [26]int
	var seen [26]bool
	for i, v := range s {
		h[v-'a'] = i
	}

	for i, v := range s {
		for len(res) > 0 {
			l := res[len(res)-1]
			if v <= l && h[l-'a'] >= i && !seen[v-'a'] {
				seen[l-'a'] = false
				res = res[:len(res)-1]
				continue
			}
			break
		}
		if !seen[v-'a'] {
			res = append(res, v)
			seen[v-'a'] = true
		}
	}
	return string(res)
}
