package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i, v := range nums {
		if _, ok := m[v]; ok && i-m[v] <= k {
			return true
		}
		m[v] = i
	}
	return false
}
