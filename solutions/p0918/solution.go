package solution

func maxSubarraySumCircular(A []int) int {
	total := 0
	maxSum := -30000
	cMax := 0
	minSum := 30000
	cMin := 0
	for _, v := range A {
		cMax = max(cMax+v, v)
		maxSum = max(maxSum, cMax)
		cMin = min(cMin+v, v)
		minSum = min(minSum, cMin)
		total += v
	}
	if maxSum > 0 {
		return max(maxSum, total-minSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
