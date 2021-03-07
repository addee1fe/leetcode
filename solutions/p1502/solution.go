package solution

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	sort.Ints(arr)
	diff := arr[0] - arr[1]
	for i := 2; i < len(arr); i++ {
		if arr[i-1]-arr[i] != diff {
			return false
		}
	}
	return true
}
