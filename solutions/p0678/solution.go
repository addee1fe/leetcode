package solution

func checkValidString(s string) bool {
	max, min := 0, 0
	for _, v := range s {
		switch {
		case v == '(':
			max++
			min++
		case v == ')':
			max--
			min--
		case v == '*':
			max++
			min--
		}
		if max < 0 {
			return false
		}
		if min < 0 {
			min = 0
		}
	}
	return min == 0
}
