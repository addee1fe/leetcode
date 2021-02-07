package solution

// Ugly problems require ugly solutions :-)
func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	ugly := make([]int, n)
	ugly[0] = 1
	f2, f3, f5 := 0, 0, 0
	for i := 1; i < n; i++ {
		ugly[i] = min(ugly[f2]*2, ugly[f3]*3, ugly[f5]*5)
		if ugly[i] == ugly[f2]*2 {
			f2++
		}
		if ugly[i] == ugly[f3]*3 {
			f3++
		}
		if ugly[i] == ugly[f5]*5 {
			f5++
		}
	}
	return ugly[n-1]
}

func min(a, b, c int) int {
	m := a
	if m > b {
		m = b
	}
	if m > c {
		m = c
	}
	return m
}
