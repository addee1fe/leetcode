package solution

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	if l1 == 0 {
		return true
	}
	l2 := len(s2)
	if l2 == 0 || l2 < l1 {
		return false
	}
	a1, a2 := [26]int{}, [26]int{}
	for i, v := range s1 {
		a1[v-'a']++
		a2[s2[i]-'a']++
	}
	if match(a1, a2) {
		return true
	}

	for i := l1; i < l2; i++ {
		a2[s2[i]-'a']++
		a2[s2[i-l1]-'a']--
		if s2[i] != s2[i-l1] && match(a1, a2) {
			return true
		}
	}
	return false
}

func match(a1, a2 [26]int) bool {
	if a1 == a2 {
		return true
	}
	return false
}
