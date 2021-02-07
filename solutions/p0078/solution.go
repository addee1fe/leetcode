package solution

func subsets(nums []int) [][]int {
	n := len(nums)
	p := 1 << n
	res := make([][]int, 0, p)
	for i := 0; i < p; i++ {
		var set []int
		for j, v := range nums {
			if i>>j&1 != 0 {
				set = append(set, v)
			}
		}
		res = append(res, set)
	}
	return res
}
