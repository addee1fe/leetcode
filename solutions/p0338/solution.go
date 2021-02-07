package solution

func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	p := 1
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			res[i] = 1
			p = i
			continue
		}
		res[i] = 1 + res[i-p]
	}
	return res
}

// Hacker's Delight, Figure 5-2.
func bitCount(x uint64) int {
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}
