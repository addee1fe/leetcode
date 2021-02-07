package solution

import "testing"

func Test_numSquares(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{12, 3},
		{13, 2},
	}
	for _, tt := range tests {
		if got := numSquares(tt.n); got != tt.want {
			t.Errorf("numSquares(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
