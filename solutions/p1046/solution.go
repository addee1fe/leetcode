package solution

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for i := len(stones) - 1; i > 0; i-- {
		stones[i-1] = stones[i] - stones[i-1]
		sort.Ints(stones)
	}
	return stones[0]
}
