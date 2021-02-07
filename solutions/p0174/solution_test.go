package solution

import "testing"

func TestCalculateMinimumHP(t *testing.T) {
	tests := []struct {
		dungeon [][]int
		want    int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
	}
	for _, tt := range tests {
		if got := calculateMinimumHP(tt.dungeon); got != tt.want {
			t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
		}
	}
}
