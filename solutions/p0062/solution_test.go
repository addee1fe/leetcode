package solution

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{3, 2, 3},
		{7, 3, 28},
	}
	for _, tt := range tests {
		if got := uniquePaths(tt.m, tt.n); got != tt.want {
			t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
		}
	}
}
