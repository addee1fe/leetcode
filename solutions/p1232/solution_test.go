package solution

import "testing"

func TestCheckStraightLine(t *testing.T) {
	tests := []struct {
		coordinates [][]int
		want        bool
	}{
		{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{6, 7}}, true},
		{[][]int{[]int{1, 1}, []int{2, 2}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{7, 7}}, false},
	}
	for _, tt := range tests {
		if got := checkStraightLine(tt.coordinates); got != tt.want {
			t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
		}
	}
}
