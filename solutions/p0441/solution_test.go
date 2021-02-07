package solution

import "testing"

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 2},
		{5, 2},
		{8, 3},
	}
	for _, tt := range tests {
		if got := arrangeCoins(tt.n); got != tt.want {
			t.Errorf("arrangeCoins(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
