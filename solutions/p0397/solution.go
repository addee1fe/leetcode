package solution

func integerReplacement(n int) int {
	var count int
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%4 == 1 || n == 3 {
			n--
		} else {
			n++
		}
		count++
	}
	return count
}
