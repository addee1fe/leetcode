package solution

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	var s []int
	j := 0
	for _, v := range pushed {
		s = append(s, v)
		for len(s) != 0 && j < n && s[len(s)-1] == popped[j] {
			s = s[:len(s)-1]
			j++
		}
	}
	return j == n
}
