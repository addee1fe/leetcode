package solution

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{[]int{}, []int{1}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}
	for _, tt := range tests {
		if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("plusOne() = %v, want %v", got, tt.want)
		}
	}
}
