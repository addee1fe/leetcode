package solution

import (
	"testing"
)

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{3, -1, 2, -1}, 4},
		{[]int{3, -2, 2, -3}, 3},
		{[]int{-2, -3, -1}, -1},
	}
	for _, tt := range tests {
		if got := maxSubarraySumCircular(tt.A); got != tt.want {
			t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
		}

	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{-1, -2, -1},
		{-2, -1, -1},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{-1, -2, -2},
		{-2, -1, -2},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}
