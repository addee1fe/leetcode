package solution

import "testing"

func TestSearchMatrix(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		{
			[][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
	}
	for _, tt := range tests {
		if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
			t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
		}
	}
}
