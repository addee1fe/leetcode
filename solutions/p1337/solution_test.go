package solution

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			3,
			[]int{2, 0, 3},
		},
		{
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			2,
			[]int{0, 2},
		},
	}
	for _, tt := range tests {
		if got := kWeakestRows(tt.mat, tt.k); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
		}
	}
}
