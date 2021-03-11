package solution

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			0,
		},
		{
			[]int{1},
			1,
			1,
		},
		{
			[]int{1},
			2,
			2,
		},
	}

	for _, tt := range tests {
		if got := coinChange(tt.coins, tt.amount); got != tt.want {
			t.Errorf("coinChange() = %v, want %v", got, tt.want)
		}
	}
}
