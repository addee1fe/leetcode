package solution

import (
	"reflect"
	"testing"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			[]int{8, 1, 2, 2, 3},
			[]int{4, 0, 1, 1, 3},
		},
		{
			[]int{6, 5, 4, 8},
			[]int{2, 1, 0, 3},
		},
		{
			[]int{7, 7, 7, 7},
			[]int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		if got := smallerNumbersThanCurrent(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
		}
	}
}
