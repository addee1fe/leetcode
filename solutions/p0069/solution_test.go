package solution

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
	}
	for _, tt := range tests {
		if got := mySqrt(tt.x); got != tt.want {
			t.Errorf("mySqrt() = %v, want %v", got, tt.want)
		}
	}
}
