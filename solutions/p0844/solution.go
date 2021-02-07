package solution

func backspaceCompare(S string, T string) bool {
	return build(S) == build(T)
}

func build(s string) string {
	ans := make([]rune, 0)
	for _, v := range s {
		if v != '#' {
			ans = append(ans, v)
		} else {
			if len(ans)-1 >= 0 {
				ans = ans[:len(ans)-1]
			}
		}
	}
	return string(ans)
}
