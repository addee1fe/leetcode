package solution

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{
			[]int{},
			true,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{1, 2},
			true,
		},
		{
			[]int{3, 5, 1},
			true,
		},
		{
			[]int{1, 2, 4},
			false,
		},
		{
			[]int{0, 1, 1, 2, 3, 5, 8},
			false,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			true,
		},
	}
	for _, tt := range tests {
		if got := canMakeArithmeticProgression(tt.arr); got != tt.want {
			t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
		}
	}
}
