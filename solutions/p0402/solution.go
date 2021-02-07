package solution

// TODO: review solution
func removeKdigits(num string, k int) string {
	if k <= 0 || num == "" {
		return num
	}
	if len(num) <= k {
		return "0"
	}
	length := len(num)
	stack := make([]byte, 0)
	stack = append(stack, '0'-1)
	for i := 0; i < length; {
		if k == 0 {
			stack = append(stack, num[i:]...)
			break
		}
		if num[i] >= stack[len(stack)-1] {
			stack = append(stack, num[i])
			i++
		} else {
			stack = stack[:len(stack)-1]
			k--
		}
	}
	stack = stack[1 : len(stack)-k]
	nozeroIndex := 0
	for ; nozeroIndex < len(stack)-1; nozeroIndex++ {
		if stack[nozeroIndex] != '0' {
			break
		}
	}
	stack = stack[nozeroIndex:]
	return string(stack)
}
