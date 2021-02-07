package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0201
// BenchmarkRangeBitwise-8                 215127164                5.55 ns/op
// BenchmarkRangeBitwiseAndRecursive-8     42692241                26.6 ns/op
// BenchmarkStolenRangeBitwise-8           50597631                22.1 ns/op

func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}
	i := 0
	for m != n {
		m >>= 1
		n >>= 1
		i++
	}
	return n << i
}

func rangeBitwiseAndRecursive(m, n int) int {
	if m == n {
		return m
	}
	if n > m {
		return rangeBitwiseAndRecursive(m/2, n/2) << 1
	}
	return m
}

// NOTE: Found this when looked for better solutions
// slightly faster than recursive approach
func stolenRangeBitwiseAnd(m int, n int) int {
	if m == n {
		return m
	}
	veri := 0x80000000
	rgtMov := 31
	res := 0

	for m&veri == n&veri {
		if m&veri == veri {
			res |= (1 << rgtMov)
		}
		rgtMov--
		m <<= 1
		n <<= 1
	}
	return res
}
