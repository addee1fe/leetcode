package solution

func maxArea(height []int) int {
	m, l, r := 0, 0, len(height)-1
	for l < r {
		m = max(m, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
