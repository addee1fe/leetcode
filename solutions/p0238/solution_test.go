package solution

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}
	for _, tt := range tests {
		if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
		}
	}
}
