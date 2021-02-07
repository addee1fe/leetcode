package solution

import "sort"

// add second pointer to the end of the slice
// like j
// so we start from both ends and do things like
// target = nums[i] + nums[j]
// if val, ok := index[v]; val + target == 0 && ok{
//	res = append(res, nums[index[v], index[i], index[j]]
//}
func threeSum(nums []int) [][]int {
	const target int = 0
	var results [][]int
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := v + nums[l] + nums[r]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				results = append(results, []int{v, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[l] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}

	}
	return results
}
