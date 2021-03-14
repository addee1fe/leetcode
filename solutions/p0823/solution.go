package solution

import (
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	mod := int(1e9 + 7)
	sort.Ints(arr)
	n := len(arr)
	index := make(map[int]int, n)

	for i, v := range arr {
		index[v]++
		for j := 0; j < i; j++ {
			k := arr[j]
			if k*k > v {
				break
			}
			if v%k == 0 {
				a := v / k
				v1, ok1 := index[k]
				v2, ok2 := index[a]
				if ok1 && ok2 {
					if k == a {
						index[v] += v1 * v2 % mod
					} else {
						index[v] += 2 * v1 * v2 % mod
					}
				}
			}
		}
	}
	var res int
	for _, v := range index {
		res = (res + v) % mod
	}
	return res
}
