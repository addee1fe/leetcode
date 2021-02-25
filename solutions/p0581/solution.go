package solution

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// min and max for avoid mess with functions
	mn, mx := 1<<63-1, -1<<63
	flag := false
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			flag = true
		}
		if flag {
			mn = min(mn, nums[i])
		}
	}
	flag = false
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			flag = true
		}
		if flag {
			mx = max(mx, nums[i])
		}
	}
	var l, r int
	for l = 0; l < n; l++ {
		if mn < nums[l] {
			break
		}
	}
	for r = n - 1; r >= 0; r-- {
		if mx > nums[r] {
			break
		}
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
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
