package solution

import (
	"math/rand"
)

// Unfortunatelly I dont' rly understand this solution and even the task

// Solution ...
type Solution struct {
	prefixSums []int
}

// Constructor ...
func Constructor(w []int) Solution {
	p := make([]int, len(w))
	var s int
	for i, v := range w {
		s += v
		p[i] = s
	}
	return Solution{p}
}

// PickIndex ...
func (s *Solution) PickIndex() int {
	r := rand.Intn(s.prefixSums[len(s.prefixSums)-1])
	lo, hi := 0, len(s.prefixSums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if r < s.prefixSums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
