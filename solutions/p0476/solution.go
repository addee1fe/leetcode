package solution

func findComplement(n int) int {
	x := ^0
	for x&n > 0 {
		x <<= 1
	}
	return ^x ^ n
}
