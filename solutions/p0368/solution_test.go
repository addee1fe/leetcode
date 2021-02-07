package solution

import (
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
	}
	for _, tt := range tests {
		if got := largestDivisibleSubset(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
		}
	}
}
