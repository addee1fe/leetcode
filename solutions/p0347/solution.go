package solution

import "sort"

// Also this can be done with PriorityQueue
// For better explanation go to documentation
// https://golang.org/pkg/container/heap/

// Item ...
type Item struct {
	v int // value
	f int // frequency
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, 0)
	for _, n := range nums {
		freq[n]++
	}
	q := make([]Item, len(freq))
	i := 0
	for k, v := range freq {
		q[i] = Item{v: k, f: v}
		i++
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].f > q[j].f
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = q[i].v
	}
	return res
}
