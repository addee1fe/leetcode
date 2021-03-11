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
