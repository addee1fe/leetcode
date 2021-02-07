package solution

import "testing"

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, tt := range tests {
		if got := maxSubArray(tt.nums); got != tt.want {
			t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
		}
	}
}
