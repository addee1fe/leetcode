package solution

// similar to p0771
func canConstruct(ransomNote string, magazine string) bool {
	// cut some corners
	if len(magazine) == 0 && len(ransomNote) == 0 {
		return true
	}
	if len(ransomNote) > len(magazine) {
		return false
	}
	sl := make([]int, 26)
	for _, v := range magazine {
		sl[v-'a']++
	}
	for _, v := range ransomNote {
		sl[v-'a']--
		if sl[v-'a'] < 0 {
			return false
		}
	}
	return true
}
