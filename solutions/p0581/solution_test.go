package solution

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{2, 6, 4, 8, 10, 9, 15},
			5,
		},
		{
			[]int{1, 2, 3, 4},
			0,
		},
	}
	for _, tt := range tests {
		if got := findUnsortedSubarray(tt.nums); got != tt.want {
			t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
		}
	}
}
func TestMin(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 1},
		{2, 3, 2},
		{3, 4, 3},
		{4, 5, 4},
		{7, 6, 6},
		{10, 8, 8},
	}
	for _, tt := range tests {
		if got := min(tt.a, tt.b); got != tt.want {
			t.Errorf("min() = %v, want %v", got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
		{4, 5, 5},
		{7, 6, 7},
		{10, 8, 10},
	}
	for _, tt := range tests {
		if got := max(tt.a, tt.b); got != tt.want {
			t.Errorf("max() = %v, want %v", got, tt.want)
		}
	}
}
