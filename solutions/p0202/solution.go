package solution

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n > 1 {
		m[n] = true
		var x int
		for x, n = n, 0; x > 0; x /= 10 {
			d := x % 10
			n += d * d
		}
		if m[n] {
			return false
		}
	}
	return true
}
