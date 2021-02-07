package solution

import "testing"

func TestLastStoneWeight(t *testing.T) {

	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
	}
	for _, tt := range tests {
		if got := lastStoneWeight(tt.stones); got != tt.want {
			t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
		}

	}
}
