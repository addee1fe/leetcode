package solution

import "strings"

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	res := []byte(s)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				res[i] = '*'
			}
		}
	}
	for _, v := range stack {
		res[v] = '*'
	}
	return strings.ReplaceAll(string(res), "*", "")
}
