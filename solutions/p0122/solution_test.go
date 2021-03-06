package solution

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want int
	}{
		{[]int{7,1,5,3,6,4},7},
		{[]int{1,2,3,4,5},4},
		{[]int{7,6,4,3,1}, 0},
	}
	for _, tt := range tests {
		if got := maxProfit(tt.prices); got != tt.want {
			t.Errorf("maxProfit() = %v, want %v", got, tt.want)
		}
	}
}
