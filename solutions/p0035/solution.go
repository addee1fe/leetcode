package solution

import "sort"

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0035
// BenchmarkSearchInsertFor-8      239043828                5.04 ns/op
// BenchmarkSearchInsert-8         181028715                6.59 ns/op
// BenchmarkSearchInsertSort-8     97626471                11.3 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0035     4.990s

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}

	}
	return lo
}

func searchInsertSort(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	return i
}

func searchInsertFor(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, v := range nums {
		if v == target || v > target {
			return i
		}
	}
	return len(nums)
}
