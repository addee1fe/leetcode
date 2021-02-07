package solution

import "testing"

func Test_islandPerimeter(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0}},
			16},
	}
	for _, tt := range tests {
		if got := islandPerimeter(tt.grid); got != tt.want {
			t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
		}
	}
}
