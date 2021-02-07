package solution

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		args []int
		want bool
	}{
		{
			[]int{2, 3, 1, 1, 4}, true,
		},
		{
			[]int{3, 2, 1, 0, 4}, false,
		},
	}
	for _, tt := range tests {
		if got := canJump(tt.args); got != tt.want {
			t.Errorf("canJump() = %v, want %v", got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{2, 3, 3},
		{4, 5, 5},
		{5, 6, 6},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("canJump() = %v, want %v", got, tt.want)
		}
	}
}
