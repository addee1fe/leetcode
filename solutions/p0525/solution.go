package solution

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	sum, max := 0, 0
	for i, v := range nums {
		if v == 1 {
			sum++
		} else {
			sum--
		}
		if sum == 0 {
			max = i + 1
		} else {
			in, ok := m[sum]
			if ok {
				if i-in > max {
					max = i - in
				}
			} else {
				m[sum] = i
			}
		}
	}
	return max
}
