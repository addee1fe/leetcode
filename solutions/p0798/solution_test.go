package solution

import "testing"

func TestBestRotation(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{
			[]int{2, 3, 1, 4, 0},
			3,
		},
		{
			[]int{1, 3, 0, 2, 4},
			0,
		},
	}
	for _, tt := range tests {
		if got := bestRotation(tt.A); got != tt.want {
			t.Errorf("bestRotation() = %v, want %v", got, tt.want)
		}
	}
}
