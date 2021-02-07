package solution

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return []int{}
	}
	uniq := [26]int{}
	curr := [26]int{}
	res := []int{}
	for _, v := range p {
		uniq[v-'a']++
	}
	for i, v := range s {
		curr[v-'a']++
		if i >= len(p) {
			curr[s[i-len(p)]-'a']--
		}
		if curr == uniq {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
