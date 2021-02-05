package solution

import "testing"

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			8,
		},
		{
			[][]int{
				{3, 2},
				{1, 0},
			},
			0,
		},
		{
			[][]int{
				{1, -1},
				{-1, -1},
			},
			3,
		},
		{
			[][]int{
				{-1},
			},
			1,
		},
	}
	for _, tt := range tests {
		if got := countNegatives(tt.grid); got != tt.want {
			t.Errorf("countNegatives() = %v, want %v", got, tt.want)
		}
	}
}

func TestCountRow(t *testing.T) {
	tests := []struct {
		row  []int
		want int
	}{
		{
			[]int{1, 0},
			0,
		},
		{
			[]int{-1},
			1,
		},
		{
			[]int{1, -1},
			1,
		},
		{
			[]int{-1, -1},
			2,
		},
		{
			[]int{4, 3, 2, -1},
			1,
		},
		{
			[]int{1, 1, -1, -2},
			2,
		},
		{
			[]int{-1, -1, -2, -3},
			4,
		},
	}
	for _, tt := range tests {
		if got := countRow(tt.row); got != tt.want {
			t.Errorf("countRow(%v) = %v, want %v", tt.row, got, tt.want)
		}
	}
}
