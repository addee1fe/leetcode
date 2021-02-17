package solution

// NOTE: this solution is not mine
func letterCasePermutation(S string) []string {
	ret := []string{S}
	return helper(S, 0, ret)
}

func helper(S string, start int, ret []string) []string {
	for i := start; i < len(S); i++ {
		if '0' <= S[i] && S[i] <= '9' {
			continue
		}
		slice := []byte(S)
		if 'A' <= slice[i] && slice[i] <= 'Z' {
			slice[i] = slice[i] - 'A' + 'a'
		} else {
			slice[i] = slice[i] - 'a' + 'A'
		}
		ret = append(ret, string(slice))
		ret = helper(string(slice), i+1, ret)
	}
	return ret
}
