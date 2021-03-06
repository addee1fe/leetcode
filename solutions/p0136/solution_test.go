package solution

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}
	for _, tt := range tests {
		if got := singleNumber(tt.nums); got != tt.want {
			t.Errorf("singleNumber() = %v, want %v", got, tt.want)
		}
	}
}
