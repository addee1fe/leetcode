package solution

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			[]int{1, 2, 2, 4},
			[]int{2, 3},
		},
		{
			[]int{1, 1},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		if got := findErrorNums(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
		}
	}
}
