package solution

import (
	"sort"
)

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	sort.Ints(citations)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if citations[mid] < n-mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return n - lo
}
