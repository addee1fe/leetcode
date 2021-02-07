package solution

import (
	"testing"
)

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 8},
		{8, 9},
		{9, 10},
		{10, 12},
	}
	for _, tt := range tests {
		if got := nthUglyNumber(tt.n); got != tt.want {
			t.Errorf("nthUglyNumber(%v) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{1, 2, 3, 1},
		{2, 1, 3, 1},
		{3, 2, 1, 1},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b, tt.c); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}

	}
}
