package solution

import (
	"sort"
)

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0169
// BenchmarkMajorityElement-8              154294911                7.66 ns/op
// BenchmarkMajorityElementMap-8           11069761               105 ns/op
// BenchmarkMajorityElementSort-8           8522024               139 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0169     4.669s

// Boyer-Moore algorithm
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	c, candidate := 0, 0
	for _, v := range nums {
		if c == 0 {
			candidate = v
		}
		if v == candidate {
			c++
		} else {
			c--
		}
	}
	return candidate
}

func majorityElementMap(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	votes := make(map[int]int)
	for _, v := range nums {
		votes[v]++
		if votes[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

func majorityElementSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[len(nums)/2]
}
