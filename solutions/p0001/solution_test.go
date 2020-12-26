package solution

import (
	"reflect"
	"testing"
)

var testData = []struct {
	nums   []int
	target int
	result []int
}{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		result: []int{0, 1},
	},
	{
		nums:   []int{1, 2, 3, 4},
		target: 7,
		result: []int{2, 3},
	},
	{
		nums:   []int{2, 7, 3, 4},
		target: 5,
		result: []int{0, 2},
	},
}

func TestTwoSum(t *testing.T) {
	for _, td := range testData {
		got := twoSum(td.nums, td.target)
		if !reflect.DeepEqual(got, td.result) {
			t.Errorf("twoSum(%v, %v) = %v, want %v", td.nums, td.target, got, td.result)
		}
	}
}
