package solution

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		fee    int
		want   int
	}{
		{
			[]int{1, 3, 2, 8, 4, 9},
			2,
			8,
		},
		{
			[]int{1, 3, 7, 5, 10, 3},
			3,
			6,
		},
	}
	for _, tt := range tests {
		if got := maxProfit(tt.prices, tt.fee); got != tt.want {
			t.Errorf("maxProfit() = %v, want %v", got, tt.want)
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
