package solution

func firstUniqChar(s string) int {
	var h [26]int
	for _, v := range s {
		h[v-'a']++
	}
	for i, v := range s {
		if h[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
