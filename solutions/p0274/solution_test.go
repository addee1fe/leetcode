package solution

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		want      int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
	}
	for _, tt := range tests {

		if got := hIndex(tt.citations); got != tt.want {
			t.Errorf("hIndex(%v) = %v, want %v", tt.citations, got, tt.want)
		}
	}
}
