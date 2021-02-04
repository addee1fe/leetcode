package solution

func findLHS(nums []int) int {
	m := make(map[int]int, len(nums))
	var res int
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if m[k-1] == 0 {
			continue
		}
		if v+m[k-1] > res {
			res = v + m[k-1]
		}
	}
	return res
}
