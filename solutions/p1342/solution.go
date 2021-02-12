package solution

func numberOfSteps(num int) int {
	count := 0
	for ; num > 0; count++ {
		if num%2 == 0 {
			num >>= 1
		} else {
			num--
		}
	}
	return count
}
