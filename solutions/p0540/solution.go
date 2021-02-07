package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0540
// BenchmarkSingleNonDuplicate-8           166798855                7.26 ns/op
// BenchmarkSingleNonDuplicateXor-8        147418432                8.18 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0540     4.094s

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) >> 1
		if nums[mid] == nums[mid^1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

// unfortunatelly this solution does not meet requirements(O(log n) - time and O(1) space)
func singleNonDuplicateXor(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	u := nums[0]
	for i := 1; i < len(nums); i++ {
		u ^= nums[i]
	}
	return u
}
