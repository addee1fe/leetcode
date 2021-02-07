package solution

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{3, 0, 22, 5, 0, 0, 12, 4, 0}, []int{3, 22, 5, 12, 4, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		moveZeroes(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("moveZeroes() = %v, want %v", tt.nums, tt.want)
		}
	}
}
