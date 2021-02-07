package solution

func bitwiseComplement(N int) int {
	x := 1
	for N > x {
		x = x*2 + 1
	}
	return x ^ N
}
