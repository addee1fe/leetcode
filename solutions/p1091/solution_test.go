package solution

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			2,
		},
		{
			[][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			4,
		},
		{
			[][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			-1,
		},
		{
			[][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 1},
			},
			-1,
		},
	}
	for _, tt := range tests {
		if got := shortestPathBinaryMatrix(tt.grid); got != tt.want {
			t.Errorf("shortestPathBinaryMatrix() = %v, want %v", got, tt.want)
		}
	}
}
