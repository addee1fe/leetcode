package solution

func scoreOfParentheses(S string) int {
	var res, bal int
	for i := range S {
		if S[i] == '(' {
			bal++
		} else {
			bal--
			if S[i-1] == '(' {
				res += 1 << bal
			}
		}
	}
	return res
}
