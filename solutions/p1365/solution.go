package solution

// also you can simply solve this with nested loops
// TODO: add solution with nested loops and benchmerks

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	buckets := make([]int, 101)
	for _, v := range nums {
		buckets[v]++
	}
	for i := 1; i < len(buckets); i++ {
		buckets[i] += buckets[i-1]
	}
	for i, v := range nums {
		if v == 0 {
			res[i] = 0
		} else {
			res[i] = buckets[v-1]
		}
	}
	return res
}
