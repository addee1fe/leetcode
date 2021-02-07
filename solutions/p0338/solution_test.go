package solution

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{
			0,
			[]int{0},
		},
		{
			1,
			[]int{0, 1},
		},
		{
			2,
			[]int{0, 1, 1},
		},
		{
			3,
			[]int{0, 1, 1, 2},
		},
		{
			4,
			[]int{0, 1, 1, 2, 1},
		},
		{
			5,
			[]int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		if got := countBits(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("countBits() = %v, want %v", got, tt.want)
		}
	}
}
