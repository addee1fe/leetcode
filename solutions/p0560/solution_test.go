package solution

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
	}
	for _, tt := range tests {
		if got := subarraySum(tt.nums, tt.k); got != tt.want {
			t.Errorf("subarraySum() = %v, want %v", got, tt.want)
		}

	}
}
