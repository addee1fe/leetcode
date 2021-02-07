package solution

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		want      int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{100}, 1},
		{[]int{0, 1, 3, 5, 6}, 3},
	}
	for _, tt := range tests {
		if got := hIndex(tt.citations); got != tt.want {
			t.Errorf("hIndex(%v) = %v, want %v", tt.citations, got, tt.want)
		}
	}
}
