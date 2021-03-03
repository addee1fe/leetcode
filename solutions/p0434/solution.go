package solution

func countSegments(s string) int {
	var sc int
	for i := range s {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			sc++
		}
	}
	return sc
}
