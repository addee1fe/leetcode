package solution

import (
	"testing"
)

func TestCountSquares(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		matrix [][]int
		want   int
	}{{
		[][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}, 15},

		{[][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}, 7},
	}
	for _, tt := range tests {
		if got := countSquares(tt.matrix); got != tt.want {
			t.Errorf("countSquares() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{-1, -2, -2},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}
