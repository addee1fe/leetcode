package solution

import "testing"

func TestFindLHS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 3, 2, 2, 5, 2, 3, 7},
			5,
		},
		{
			[]int{1, 2, 3, 4},
			2,
		},
		{
			[]int{1, 1, 1, 1},
			0,
		},
	}
	for _, tt := range tests {
		if got := findLHS(tt.nums); got != tt.want {
			t.Errorf("findLHS() = %v, want %v", got, tt.want)
		}
	}
}
