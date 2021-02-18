package solution

//
func numberOfArithmeticSlices(A []int) int {
	var count, sum int
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			count++
		} else {
			sum += (count + 1) * (count) / 2
			count = 0
		}
	}
	return sum + count*(count+1)/2
}
