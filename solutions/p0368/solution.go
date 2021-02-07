package solution

import (
	"sort"
)

// TODO: refactor this bullshit

func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	parent := make([]int, n)
	maxNum, maxIndex := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if nums[j]%nums[i] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parent[i] = j
				if dp[i] > maxNum {
					maxNum = dp[i]
					maxIndex = i
				}
			}
		}
	}
	res := make([]int, 0, n)
	currIndex := maxIndex
	for i := 0; i < maxNum; i++ {
		res = append(res, nums[currIndex])
		currIndex = parent[currIndex]
	}
	return res
}
