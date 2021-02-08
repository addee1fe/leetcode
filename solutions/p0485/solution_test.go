package solution

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{1, 1, 0, 1, 1, 1},
			3,
		},
	}
	for _, tt := range tests {
		if got := findMaxConsecutiveOnes(tt.nums); got != tt.want {
			t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
		}
	}
}
