package solution

import "sort"

// Please note
// This is extremely dummy solution and in BEST case the time complexity is O(n) which not meet the requirements

func findMedianSortedArrays(a []int, b []int) float64 {
	c := append(a, b...)
	sort.Ints(c)
	if len(c)%2 == 0 {
		mid := (len(c) / 2) - 1
		return float64(c[mid]+c[mid+1]) / 2.0
	}
	return float64(c[len(c)/2])
}
