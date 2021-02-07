package solution

import "testing"

func Test_change(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, tt := range tests {
		if got := change(tt.amount, tt.coins); got != tt.want {
			t.Errorf("change() = %v, want %v", got, tt.want)
		}
	}
}
